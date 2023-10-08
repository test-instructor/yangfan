## 前端环境
1. 前往https://nodejs.org/zh-cn/下载当前版本node 
2. 命令行运行 node -v 若控制台输出版本号则前端环境搭建成功 
3. node 版本需大于 16.4 
4. 开发工具推荐vscode https://code.visualstudio.com/
5. 安装依赖
   ```shell
   cd web
   npm install
   ```

## 后端环境
1. 下载golang安装 版本号需>=1.19
   * 国际: https://golang.org/dl/
   * 国内: https://golang.google.cn/dl/
2. 命令行运行 go 若控制台输出各类提示命令 则安装成功 输入 go version 确认版本大于 1.18 
3. 开发工具推荐 Goland
4. 安装依赖
   ```shell
   cd server
   go mod tidy
   ```
### 服务初始化 
1. 服务列表
   * server 后端服务
   * run 用例运行服务
   * timer 定时任务服务
   * master 性能测试master服务
   * work 性能测试worker服务
   * web 前端服务
2. 安装依赖
    ```shell
    cd server # 其他服务同理，每个服务都需要安装
    go mod tidy
    ```
### 本地调试

#### 软件包路径

* server ：github.com/test-instructor/yangfan/server
* run ：github.com/test-instructor/yangfan/run
* timer ：github.com/test-instructor/yangfan/timer
* master ：github.com/test-instructor/yangfan/master
* work：github.com/test-instructor/yangfan/work

#### 添加配置

1. 进入编辑配置页面
2. 添加“Go 构建”的配置
3. 运行种类选择软件包
4. 软件包路径选择对应的服务
5. 名称默认为软件包路径，为了更加简洁可以直接改成对应服务的名称

![image-20230928162540358](https://qiniu.yangfan.gd.cn//markdown/image-20230928162540358.png)

![image-20230928162759825](https://qiniu.yangfan.gd.cn//markdown/image-20230928162759825.png)

![image-20230928162928134](https://qiniu.yangfan.gd.cn//markdown/image-20230928162928134.png)