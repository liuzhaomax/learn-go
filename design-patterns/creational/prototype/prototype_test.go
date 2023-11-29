package prototype

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestConcretePrototype_Clone(t *testing.T) {
	name := "abcde"
	p := ConcretePrototype{name: name}
	newProto := p.Clone()
	assert.Equal(t, name, newProto.Name())
}
