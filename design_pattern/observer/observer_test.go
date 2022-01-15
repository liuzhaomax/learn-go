/****************************************************************************
 * @copyright   LIU Zhao
 * @authors     LIU Zhao (liuzhaomax@163.com)
 * @date        2022/1/16 4:25
 * @version     v1.0
 * @filename    observer_test.go
 * @description
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
