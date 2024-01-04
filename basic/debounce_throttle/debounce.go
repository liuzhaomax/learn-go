package debounce_throttle

import "context"

func Debounce(f func(), wait int) func() {
	var cf context.CancelFunc
	return func() {
		if cf != nil {
			cf()
		}
		cf = SetTimeout(f, wait)
	}
}
