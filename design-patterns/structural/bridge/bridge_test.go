package bridge

import "testing"

func TestCircle_Draw(t *testing.T) {
	red := Circle{}
	red.Constructor(100, 100, 10, &RedCircle{})
	red.Draw()

	yellow := Circle{}
	yellow.Constructor(200, 200, 20, &YellowCircle{})
	yellow.Draw()
}
