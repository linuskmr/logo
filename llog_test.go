package llog

import (
	"testing"
)

func TestMyLogger_Info(t *testing.T) {
	Info("Hallo", "Linus")
}