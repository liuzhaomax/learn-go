/****************************************************************************
 * @copyright   LIU Zhao
 * @authors     LIU Zhao (liuzhaomax@163.com)
 * @date        2022/1/16 4:52
 * @version     v1.0
 * @filename    decorator.go
 * @description
 ***************************************************************************/
package decorator

import (
	"log"
	"math"
	"time"
)

type PiFunc func(int) float64

func WrapLogger(fun PiFunc, logger *log.Logger) PiFunc {
	return func(n int) float64 {
		fn := func(n int) (result float64) {
			defer func(t time.Time) {
				logger.Printf("took=%v, n=%v, result=%v", time.Since(t), n, result)
			}(time.Now())
			return fun(n)
		}
		return fn(n)
	}
}

func Pi(n int) float64 {
	ch := make(chan float64)
	for i := 0; i < n; i++ {
		go func(ch chan float64, i float64) {
			ch <- 4 * math.Pow(-1, i) / (2*i + 1)
		}(ch, float64(i))
	}
	result := 0.0
	for i := 0; i < n; i++ {
		result += <-ch
	}
	return result
}
