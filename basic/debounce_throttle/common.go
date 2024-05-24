package debounce_throttle

import (
	"context"
	"time"
)

func SetTimeout(f func(), timeout int) context.CancelFunc {
	ctx, cancelFunc := context.WithCancel(context.Background())
	go func() {
		select {
		case <-ctx.Done():
			return
		case <-time.After(time.Duration(timeout) * time.Second):
			f()
		}
	}()
	return cancelFunc
}

func SetInterval(f func(), timeout int) context.CancelFunc {
	ctx, cancelFunc := context.WithCancel(context.Background())
	go func() {
		for {
			time.Sleep(time.Duration(timeout) * time.Second)
			select {
			case <-ctx.Done():
				return
			default:
				f()
			}
		}
	}()
	return cancelFunc
}
