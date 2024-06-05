package main

import (
	"fmt"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/eyebluecn/sc-bff/src/common/config"
	config2 "github.com/eyebluecn/sc-bff/src/infra/rpc/config"
	"github.com/eyebluecn/sc-bff/src/router"
)

func main() {
	h := server.Default(server.WithHostPorts(fmt.Sprintf("0.0.0.0:%v", config.BffServerPort)))

	//初始化RPC客户端
	config2.InitRpcClient()

	//注册路由
	router.NewRouter().Register(h)

	h.Spin()
}
