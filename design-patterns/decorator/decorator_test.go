package decorator

import (
	"log"
	"os"
	"testing"
)

func TestWrapLogger(t *testing.T) {
	foo := WrapLogger(Pi, log.New(os.Stdout, "test", 1))
	foo(10000)
}
