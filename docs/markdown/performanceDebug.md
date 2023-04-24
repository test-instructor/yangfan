# 性能测试资源占用问题排查及解决


* 运行压测前资源消耗

![img.png](https://testerhome.com/uploads/photo/2022/8eee468a-aee8-4a91-b43c-519a9056ae68.png)

* 运行压测时(10个并发用户)时资源消耗

![img_1.png](https://testerhome.com/uploads/photo/2022/4eb6045a-9768-46f2-bd04-3ce38be8a567.png)

* 压测数据丢失，从压测报告中发现有大量的测试数据丢失，导致数据不完整

![img_5.png](https://testerhome.com/uploads/photo/2022/61c5a0e0-57d5-4a99-a8da-dfef1c4d79df.png)

* 通过日志分析，启动压测时大量运行python插件，通过进程管理，goland下挂着N多的python程序，还好每次运行插件的目录是同一个，否则就可以看到运行时生成N多的debugtalk.py文件

![img_2.png](https://testerhome.com/uploads/photo/2022/c1f3b0dc-b8c2-4165-8a37-35aa7187ebef.png)

![img_3.png](https://testerhome.com/uploads/photo/2022/0fe73714-100f-4c08-a9a5-41a3784d38cb.png)


# 问题排查
* 根据日志得出，有大量开始运行python插件的信息，根据`start to prepare python plugin`查到`initPlugin`下调用了`BuildPlugin`
* 在`NewCaseRunner`的时候调用了`initPlugin`
* boomer中有2个地方调用了`NewCaseRunner`函数
![img_4.png](https://testerhome.com/uploads/photo/2022/e5464e61-2818-4135-8343-360df5b58552.png)

# 解决问题思路及优缺点


1. 在用例运行结束时停止`plugin`的操作
   * 优点：不用太多操作，不用考虑线程问题
   * 缺点：每个并发数都会运行一个plugin，循环运行，如果一千个用户，那么运行就会有一千个python进程，从上面看，每个python进程需要13M，大约需要13G内存，如果一万个用户需要`130G`内存，显然也不是很好
2. 修改`NewCaseRunner`，在调用`initPlugin`时增加限制，一个debugtalk只运行一次
   * 优点：一次并发只运行一次
   * 缺点：会造成plugin压力比较大，一万个用户同时操作可能会导致问题

综合考虑考虑下，采用`2`方式会更加稳妥，缺点可以采用分布式压测做分流问题，分布式压测后续可以通过k8s部署多个工作节点

# 解决方案
1. 每次运行时有个唯一值，以此来确定是否需要执行initPlugin，好在之前的设计的 debugtalk path 是唯一值，这次不用改动到其他代码
2. 通过单例设计模式实现，因为是多线程运行，所以需要加锁
3. 通过map实现可以同时运行多个实例，主要是兼容接口测试，否则不同项目访问到的可能是同一个plugin实例

```go
// Package hrp NewCaseRunner 方法下把initPlugin修改成yangfanInitPlugin
package hrp

func (r *HRPRunner) NewCaseRunner(testcase *TestCase) (*CaseRunner, error) {
   caseRunner := &CaseRunner{
      testCase:  testcase,
      hrpRunner: r,
      parser:    newParser(),
   }

   // init parser plugin
   //plugin, err := initPlugin(testcase.Config.Path, r.venv, r.pluginLogOn)
   //压测运行时会同时运行多个plugin，用单例方式控制每次压测任务只能运行一个plugin
   plugin, err := yangfanInitPlugin(testcase.Config.Path, r.venv, r.pluginLogOn)

   if err != nil {
      return nil, errors.Wrap(err, "init plugin failed")
   }
   if plugin != nil {
      caseRunner.parser.plugin = plugin
      caseRunner.rootDir = filepath.Dir(plugin.Path())
   }

   // ... 省略其他代码
}
```

```go
// Package hrp 增加 yangfanInitPlugin 函数
package hrp

import (
   "github.com/httprunner/funplugin"
   "github.com/pkg/errors"
   "sync"
)

var yangfanPlugin = make(map[string]*funplugin.IPlugin)
var mutex sync.Mutex

func yangfanInitPlugin(path, venv string, logOn bool) (plugin funplugin.IPlugin, err error) {
   plugins := yangfanPlugin[path]
   if plugins == nil {
      mutex.Lock()
      defer mutex.Unlock()
      if plugins == nil {
         plugin, err = initPlugin(path, venv, logOn)
         if err != nil {
            return nil, errors.Wrap(err, "init plugin failed")
         }
         yangfanPlugin[path] = &plugin
         plugins = &plugin
      }
   }
   return *plugins, nil
}

```

# 解决后效果
* 数据不再丢失

![img_6.png](https://testerhome.com/uploads/photo/2022/4c0652b4-2b7d-4ea2-b5f0-598a06dd97bb.png)

* cpu、内存数据正常
![img_7.png](https://testerhome.com/uploads/photo/2022/d3e0b96c-919f-46a4-ab6a-8403dfb42ab2.png)

# 往期文档
* yangfan 自动化测试平台[开源项目](https://testerhome.com/opensource_projects/yangfan)
* 了解如何接入httprunner
* [测试平台接入 HttpRunner V4（一）基本功能接入](https://testerhome.com/topics/35126)
* [测试平台接入 HttpRunner V4（二）使用 config 实现用例之间的参数传递](https://testerhome.com/topics/35125)
* [测试平台接入 HttpRunner V4（三）详细功能介绍](https://testerhome.com/topics/35161)