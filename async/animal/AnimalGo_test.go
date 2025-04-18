package animal

import (
	"context"
	"testing"
)

func TestAnimalGo(t *testing.T) {
	ctx, _ := context.WithCancel(context.Background())
	AnimalGo(ctx, 10)
}
