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