package main

import (
	"fmt"
	sentinel "github.com/alibaba/sentinel-golang/api"
	"github.com/alibaba/sentinel-golang/core/base"
	"github.com/alibaba/sentinel-golang/core/flow"
	"math/rand"
	"time"
)

// https://sentinelguard.io/zh-cn/docs/golang/flow-control.html

const resName = "cart-order"

func main() {
	err := sentinel.InitDefault()
	if err != nil {
		panic(err)
	}

	var all int
	var through int
	var block int
	ch := make(chan struct{})

	var rules []*flow.Rule
	rule := &flow.Rule{
		Resource:               resName,
		TokenCalculateStrategy: flow.WarmUp,
		ControlBehavior:        flow.Reject,
		Threshold:              1000,
		WarmUpPeriodSec:        30, // 30秒通过1000个
	}
	rules = append(rules, rule)
	_, err = flow.LoadRules(rules)
	if err != nil {
		panic(err)
	}

	for i := 0; i < 10; i++ {
		go func() {
			for {
				all++
				entry, blockError := sentinel.Entry(resName, sentinel.WithTrafficType(base.Inbound)) // 对resName资源进行埋点，类型为限入
				if blockError != nil {
					// fmt.Println("流量过大，开启限流")
					time.Sleep(time.Duration(rand.Uint64()%10) * time.Millisecond)
					block++
				} else {
					// fmt.Println("正常通过")
					time.Sleep(time.Duration(rand.Uint64()%10) * time.Millisecond)
					entry.Exit()
					through++
				}
			}
		}()
	}

	go func() {
		var oldAll int
		var oldThrougth int
		var oldBlock int
		for {
			a := all - oldAll
			oldAll = all

			t := through - oldThrougth
			oldThrougth = through

			b := block - oldBlock
			oldBlock = block

			time.Sleep(time.Second * 1)
			fmt.Printf("all: %d, through: %d, block: %d\n", a, t, b)
			// 慢慢的block会变为0
		}
	}()

	<-ch
}
