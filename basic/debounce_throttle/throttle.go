package debounce_throttle

import "context"

func Throttle(f func(), wait int) func() {
	var cf context.CancelFunc
	return func() {
		if cf == nil {
			cf = SetInterval(f, wait)
		}
	}
}
