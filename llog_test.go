package llog

import (
	"testing"
)

func TestMyLogger_Info(t *testing.T) {
	Info("Hallo", "Linus")
}

func TestPrint(t *testing.T) {
	Print("Hallo")
}
