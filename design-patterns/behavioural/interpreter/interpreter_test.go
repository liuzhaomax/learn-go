package interpreter

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestVariable_Interpret(t *testing.T) {
	expression := "z x w +"
	sentence := NewEvaluator(expression)
	variable := make(map[string]Expression)
	variable["z"] = &Integer{43}
	variable["w"] = &Integer{6}
	variable["x"] = &Integer{10}
	result := sentence.Interpret(variable)
	assert.Equal(t, 51, result)
}
