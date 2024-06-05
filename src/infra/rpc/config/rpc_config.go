package config

import (
	"context"
	"fmt"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/eyebluecn/sc-bff/src/common/config"
	"github.com/eyebluecn/sc-misc-idl/kitex_gen/sc_misc_api/miscservice"
	"github.com/eyebluecn/sc-subscription-idl/kitex_gen/sc_subscription_api/subscriptionservice"
	"time"
)

var (
	MiscRpcClient         miscservice.Client
	SubscriptionRpcClient subscriptionservice.Client
)

func InitRpcClient() {

	miscRpcClient, err := miscservice.NewClient("MiscService",
		client.WithHostPorts(fmt.Sprintf("0.0.0.0:%v", config.MiscServerPort)), //地址和端口
		client.WithConnectTimeout(100*time.Microsecond),                        //连接超时时间 100ms
		client.WithRPCTimeout(2*time.Second),                                   //一次RPC调用最大时间 2s

	)
	if err != nil {
		klog.CtxInfof(context.Background(), "miscservice client init error: %v", err)
		panic(err)
	}
	subscriptionRpcClient, err := subscriptionservice.NewClient("SubscriptionService",
		client.WithHostPorts(fmt.Sprintf("0.0.0.0:%v", config.SubscriptionServerPort)),
		client.WithConnectTimeout(100*time.Microsecond), //连接超时时间 100ms
		client.WithRPCTimeout(2*time.Second),            //一次RPC调用最大时间 2s
	)
	if err != nil {
		klog.CtxInfof(context.Background(), "subscriptionservice client init error: %v", err)
		panic(err)
	}

	MiscRpcClient = miscRpcClient
	SubscriptionRpcClient = subscriptionRpcClient

}
