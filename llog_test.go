package llog

import (
	"testing"
)

func TestMyLogger_Info(t *testing.T) {
	Info("Info", "Test")
}

func TestPrint(t *testing.T) {
	Print("Print")
}

func TestError(t *testing.T) {
	Error("Error")
}

func TestDebug(t *testing.T) {
	Debug("Debug")
}
