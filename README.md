

## 安装环境
如果首次运行，需要先安装thriftgo和kitex，参考[kitex官网](https://cloudwego.cn/zh/docs/kitex/getting-started/prerequisite/)
```shell
# 安装thriftgo 会将thriftgo命令安装至GOPATH.
go install github.com/cloudwego/hertz/cmd/hz@latest
# 验证是否安装好了thriftgo  eg:thriftgo 0.3.12
thriftgo --version

# 安装hz 会将kitex命令安装至GOPATH.
go install github.com/cloudwego/hertz/cmd/hz@latest
# 验证是否安装好了hz  eg:hz version v0.9.0
hz --version
```


## IDL
```shell
# 生成idl中定义的模型
hz model -module github.com/eyebluecn/sc-bff --idl idl/smart_classroom_bff.thrift --model_dir ./idl_gen
```

## 安装依赖
```shell
go get github.com/eyebluecn/sc-misc-idl@master
go get github.com/eyebluecn/sc-subscription-idl@master
```


