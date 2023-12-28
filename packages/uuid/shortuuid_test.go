package uuid

import (
	"fmt"
	"github.com/kjk/betterguid"
	"github.com/lithammer/shortuuid"
	"testing"
)

func TestShortUUID(t *testing.T) {
	id := shortuuid.New()
	fmt.Println(id)

	s := betterguid.New()
	fmt.Println(s)

}
