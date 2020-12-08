package pld_logger

import (
	"github.com/michaelzx/pld/pld_config"
	"testing"
)

func TestInit(t *testing.T) {
	Init(pld_config.LoggerConfig{Mode: "dev"})
	Debug("xxxx")
	Info("xxxx")
	Warn("xxxx")
	Error("xxxx")
	DPanic("xxxx")
	Panic("xxxx")
	Fatal("xxxx")
}
