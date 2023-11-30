package channel

import (
	"fmt"
	"time"
)

func count(n int, animal string, ch chan string) {
	for i := 0; i < n; i++ {
		fmt.Println(i+1, animal)
		time.Sleep(time.Millisecond * 500)
	}
	close(ch)
}
