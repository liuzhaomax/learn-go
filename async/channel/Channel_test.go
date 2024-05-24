package channel

import (
	"testing"
)

func TestChannel(t *testing.T) {
	ch := make(chan string)
	go count(5, "sheep", ch)
	<-ch
}
