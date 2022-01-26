/****************************************************************************
 * @copyright   LIU Zhao
 * @authors     LIU Zhao (liuzhaomax@163.com)
 * @date        2022/1/16 4:25
 * @version     v1.0
 * @filename    observer_test.go
 * @description 在主体里注册观察者，主体调用Notify，参数是某个事件，在Notify中将事件传入观察者并执行观察者的函数
 ***************************************************************************/
package observer

import (
	"sync"
	"testing"
	"time"
)

func TestFib(t *testing.T) {
	//for x := range Fib(10) {
	//	fmt.Println(x)
	//}
	sub := &EventSubject{
		Observers: sync.Map{},
	}
	obs1 := &EventObserver{Id: 1, Time: time.Now()}
	obs2 := &EventObserver{Id: 2, Time: time.Now()}

	sub.AddListener(obs1)
	sub.AddListener(obs2)

	for x := range Fib(10) {
		sub.Notify(Event{Data: x})
	}
}
