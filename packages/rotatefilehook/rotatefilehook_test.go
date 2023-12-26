package rotatefilehook

import (
	"testing"
)

func TestNewLogger(t *testing.T) {
	log := NewLogger()

	log.Info("Hello Info")
	log.Warn("Hello Warn")

	log.WithField("int", 123).WithField("string", "haha").Error("Hello Error")

	log.Say("Hello Say")
	log.Sayf("Hello Say %d", 123)
	log.SayWithField("Say with field", "int", 1)
	log.SayWithFields("Say with field", map[string]interface{}{
		"animal": "walrus",
		"val":    123,
	})
}
