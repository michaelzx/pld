package pld_logger

import (
	"cirs-sds-master-server/internal/configs"
	"testing"
)

func TestInit(t *testing.T) {
	Init(configs.LoggerConfig{Mode: "dev"})
	Debug("xxxx")
	Info("xxxx")
	Warn("xxxx")
	Error("xxxx")
	DPanic("xxxx")
	Panic("xxxx")
	Fatal("xxxx")
}
