在源码解析过程中，我们已经讲解了 Httprunner（hrp）的核心功能。如果您发现有部分内容不够清楚，或者希望更深入了解特定功能的源代码，欢迎随时联系作者。

对于 Go 语言，很多测试从业者可能对安装第三方库、编译和交叉编译不太熟悉，这可能增加了学习成本。在本文中，我们将详细介绍如何进行编译，而在后续文章中，我们将进一步探讨一些简单的修改，例如跳过逻辑、修改请求和响应的头部状态、使用模拟数据、添加断言等。如果您有二次开发的需求，也可以随时联系作者。

## 修改版本号
虽然不强制要求修改版本号，但为了更好地区分不同版本，建议您修改版本号，并可以在版本号前或后添加前缀，例如："yangfan.v4.4.0" 或 "v4.4.0.yangfan"。版本号的存放路径应为："/httprunner/hrp/internal/version/VERSION"，它应该是一个字符串格式的文本文件。

## 更新依赖
```shell
# 首先，将 fork 的项目克隆到本地
git clone git@github.com:taylor9158/httprunner.git
```
### 命令行更新
```shell
cd httprunner
# 此步骤可以跳过，一般情况下都会出现timeout的错误
go mod tidy
```
#### 错误处理
请注意，更新依赖库时可能会因网络问题而出现超时错误。您可以通过设置 Go 代码模块来解决此问题：
```shell
go env -w GOPROXY=https://goproxy.cn,direct
go mod tidy
```

### 使用IDE(Goland)进行更新
1. 打开httprunner目录
2. 在终端中执行命令
```shell
# 此步骤可以跳过，一般情况下都会出现timeout的错误
go mod tidy
```
#### 错误处理
PS: 当使用 IDE 更新依赖库时，也可能会因网络问题而出现超时错误。您可以通过设置代码模块来解决这个问题
1. 打开「设置 - Go - go 模块」
2. 在「go 模块」中启用「启用 Go 模块集成」，并在输入框中填入「GOPROXY=https://goproxy.cn,direct」
    ![gosetting](./img/gosetting.png)
3. 重新执行「go mod tidy」（这一步骤可选）
4. 如果设置后执行仍然有问题，尝试重启 Goland，然后再次执行「go mod tidy」

## 编译
在完成依赖安装后，您可以直接执行脚本进行编译，脚本的路径是「scripts/build.sh」，可以通过执行 make 或 bash 来进行编译。

提示：如果在添加了标签后无法成功编译，目前还不清楚问题出在哪里。这是当前系统下的构建文件。如果您想构建适用于其他系统的可执行文件，就需要进行交叉编译。
```shell
# Usage:
# $ make build
# $ make build tags=opencv
# or
# $ bash scripts/build.sh
# $ bash scripts/build.sh opencv
```
这是构建当前系统下用的文件，如果你想构建其他系统的可执行文件，就要用到交叉编译
### 交叉编译
1. 交叉编译使用以下变量：
    * GOOS：用于指定目标操作系统，例如 "windows"、"linux"、"darwin"（macOS）、"freebsd" 等。
    * GOARCH：用于指定目标架构，例如 "amd64"、"386"、"arm" 等。
    * CGO_ENABLED：在交叉编译时，建议将其设置为 0，以确保生成的二进制文件不依赖于系统的 C 库，从而确保生成的二进制文件可以在不同平台上正确运行。
2. 以编译linux为例，执行构建命令
    ```shell
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 bash scripts/build.sh
    ```
## 编译windows可执行文件出现错误
1. 用命令***CGO_ENABLED=0 GOOS=windows GOARCH=amd64 bash scripts/build.sh***无法正常输出后缀为exe的文件
2. `build.sh` 文件中使用编译的命令为***go build -ldflags '-s -w' -o "output/hrp" hrp/cmd/cli/main.go***
3. 由于输出的统一目录为`output/hrp`,所以生成的文件无法执行，下面是正确的执行：
```shell
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -ldflags '-s -w' -o "output/hrp.exe" hrp/cmd/cli/main.go
```