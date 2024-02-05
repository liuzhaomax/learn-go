package main

import (
	"fmt"
	sentinel "github.com/alibaba/sentinel-golang/api"
	"github.com/alibaba/sentinel-golang/core/base"
	"github.com/alibaba/sentinel-golang/core/config"
	"github.com/alibaba/sentinel-golang/core/flow"
	"github.com/alibaba/sentinel-golang/logging"
	"math/rand"
	"time"
)

// https://sentinelguard.io/zh-cn/docs/golang/flow-control.html

const resName = "cart-order"

func main() {
	configuration := config.NewDefaultConfig()
	configuration.Sentinel.Log.Logger = logging.NewConsoleLogger()
	err := sentinel.InitWithConfig(configuration)
	if err != nil {
		panic(err)
	}

	// var rules []*flow.Rule
	// rule := &flow.Rule{
	// 	Resource:               resName,
	// 	TokenCalculateStrategy: flow.Direct,
	// 	ControlBehavior:        flow.Reject,
	// 	// 1000ms内允许的最大流量是10，QPS=10
	// 	Threshold:        10,
	// 	StatIntervalInMs: 1000,
	// }
	// rules = append(rules, rule)
	// _, err = flow.LoadRules(rules)
	// if err != nil {
	// 	panic(err)
	// }

	ch := make(chan struct{})

	for i := 0; i < 2; i++ {
		go func() {
			for {
				entry, blockError := sentinel.Entry(resName, sentinel.WithTrafficType(base.Inbound)) // 对resName资源进行埋点，类型为限入
				if blockError != nil {
					fmt.Println("流量过大，开启限流")
					time.Sleep(time.Duration(rand.Uint64()%10) * time.Millisecond)
				} else {
					fmt.Println("正常通过")
					time.Sleep(time.Duration(rand.Uint64()%10) * time.Millisecond)
					entry.Exit()
				}
				// time.Sleep(time.Millisecond * 200) // 200ms来2个，规则是100ms过1个：全过
				time.Sleep(time.Millisecond * 100) // 100ms来2个，规则是100ms过1个：过一半
			}
		}()
	}

	go func() {
		time.Sleep(time.Second * 1) // 1秒后开启限流
		_, err = flow.LoadRules([]*flow.Rule{
			{
				Resource:               resName, // 通过此属性锚定资源
				TokenCalculateStrategy: flow.Direct,
				ControlBehavior:        flow.Reject,
				Threshold:              10,
				StatIntervalInMs:       1000,
			},
		})
		if err != nil {
			panic(err)
		}
	}()

	<-ch
}
