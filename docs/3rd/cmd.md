# CMD

## 安装 cobra

```bash
# install cobra lib
go get -u github.com/spf13/cobra@latest

# install cobra-cli command
go install github.com/spf13/cobra-cli@latest

# ensure GOPATH/bin is configed in the PATH
export PATH=$PATH:$(go env GOPATH)/bin
```

## 初始化项目

```bash
mkdir go-starter
cd go-starter
go mod init github.com/yanjiulab/go-starter
cobra-cli init --author "Yanjiulab" --license MIT --viper

cobra-cli add serve
cobra-cli add cli

cobra-cli add create -p 'configCmd'

```