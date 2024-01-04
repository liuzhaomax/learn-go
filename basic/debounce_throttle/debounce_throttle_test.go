package debounce_throttle

import (
	"fmt"
	"os"
	"testing"
	"time"
)

// 测试防抖是通过for循环定时3秒（保证在防抖函数触发前）连续触发五次，保证输出`testDebounce`只在最后一次开始计算
func TestDebounce(t *testing.T) {
	f := Debounce(func() {
		fmt.Println("testDebounce")
		os.Exit(0)
	}, 5)
	for i := 0; i < 5; i++ {
		f()
		time.Sleep(3 * time.Second)
	}
	time.Sleep(100 * time.Second)
}

// 测试节流是通过for循环定时3秒连续触发五次，保证输出`testThrottled`只在指定时间内不受触发次数影响
func TestThrottle(t *testing.T) {
	f := Throttle(func() {
		fmt.Println("testThrottled")
		os.Exit(0)
	}, 5)
	for i := 0; i < 5; i++ {
		f()
		time.Sleep(3 * time.Second)
	}
	time.Sleep(100 * time.Second)
}
