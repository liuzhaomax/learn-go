package channel

import (
	"fmt"
	"testing"
)

func TestChannel(t *testing.T) {
	ch := make(chan string)
	go count(5, "sheep", ch)
	for message := range ch {
		fmt.Println(message)
	}
}
