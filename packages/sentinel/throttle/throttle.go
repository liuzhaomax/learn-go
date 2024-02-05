package main

import (
	"fmt"
	sentinel "github.com/alibaba/sentinel-golang/api"
	"github.com/alibaba/sentinel-golang/core/base"
	"github.com/alibaba/sentinel-golang/core/flow"
	"time"
)

// https://sentinelguard.io/zh-cn/docs/golang/flow-control.html
// 这种方式主要用于处理间隔性突发的流量，例如消息队列。想象一下这样的场景，在某一秒有大量的请求到来，
// 而接下来的几秒则处于空闲状态，我们希望系统能够在接下来的空闲期间逐渐处理这些请求，
// 而不是在第一秒直接拒绝多余的请求。

// 当请求到来的时候，如果当前请求距离上个通过的请求通过的时间间隔不小于预设值，则让当前请求通过；
// 否则，计算当前请求的预期通过时间，如果该请求的预期通过时间小于规则预设的 timeout 时间，
// 则该请求会等待直到预设时间到来通过（排队等待处理）；
// 若预期的通过时间超出最大排队时长，则直接拒接这个请求。

const resName = "cart-order"

func main() {
	err := sentinel.InitDefault()
	if err != nil {
		panic(err)
	}

	var rules []*flow.Rule
	rule := &flow.Rule{
		Resource:               resName,
		TokenCalculateStrategy: flow.Direct,
		ControlBehavior:        flow.Throttling, // 匀速限流
		Threshold:              10,
		StatIntervalInMs:       1000, // 1000毫秒允许10个，平均一个100ms
		// MaxQueueingTimeMs:      500,  // 一个队列最多等500毫秒，超过500就拒绝
		MaxQueueingTimeMs: 50, // 最大排队时长  一个队列最多等50毫秒，超过50就拒绝
	}
	rules = append(rules, rule)
	_, err = flow.LoadRules(rules)
	if err != nil {
		panic(err)
	}

	ch := make(chan struct{})

	for i := 0; i < 2; i++ {
		count := 0
		routine := ""
		go func(i int) {
			if i == 0 {
				routine = "A"
			} else {
				routine = "B"
			}
			for {
				entry, blockError := sentinel.Entry(resName, sentinel.WithTrafficType(base.Inbound)) // 对resName资源进行埋点，类型为限入
				if blockError != nil {
					fmt.Printf("流量过大，开启限流: %s - %d\n", routine, count)
					// time.Sleep(time.Duration(rand.Uint64()%10) * time.Millisecond)
				} else {
					fmt.Printf("正常通过: %s - %d\n", routine, count)
					// time.Sleep(time.Duration(rand.Uint64()%10) * time.Millisecond)
					entry.Exit()
				}
				count++
				// time.Sleep(time.Millisecond * 100) // 100ms过2个，规则是100ms1个：过一半
				time.Sleep(time.Millisecond * 50) // 50ms过2个，规则是100ms1个
				if count%2 == 0 {
					fmt.Println("--------------")
				}
			}
		}(i)
	}

	<-ch
}
