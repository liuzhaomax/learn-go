/****************************************************************************
 * @copyright   LIU Zhao
 * @authors     LIU Zhao (liuzhaomax@163.com)
 * @date        2022/1/16 4:02
 * @version     v1.0
 * @filename    observer.go
 * @description
 ***************************************************************************/
package observer

import (
	"fmt"
	"sync"
	"time"
)

type Event struct {
	Data int
}

type Observer interface {
	NotifyCallback(event Event)
}

type Subject interface {
	AddListener(observer Observer)
	RemoveListener(observer Observer)
	Notify(event Event)
}

type EventObserver struct {
	Id   int
	Time time.Time
}

type EventSubject struct {
	Observers sync.Map
}

func (eo *EventObserver) NotifyCallback(event Event) {
	fmt.Printf("Received: %d after %v\n", event.Data, time.Since(eo.Time))
}

func (es *EventSubject) AddListener(ob Observer) {
	es.Observers.Store(ob, struct{}{})
}

func (es *EventSubject) RemoveListener(ob Observer) {
	es.Observers.Delete(ob)
}

func (es *EventSubject) Notify(event Event) {
	es.Observers.Range(func(key, value interface{}) bool {
		if key == nil {
			return false
		}
		key.(Observer).NotifyCallback(event)
		return true
	})
}

func Fib(n int) chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for i, j := 0, 1; i < n; i, j = i+j, i {
			out <- i
		}
	}()
	return out
}
