package stdlog

import (
	"log"
	"os"
	"testing"
)

type ColorGroup struct {
	ID     int
	Name   string
	Colors []string
}

func TestLogger_Debug(t *testing.T) {
	l := &Logger{
		logger: log.New(os.Stderr, debugPrefix, log.LstdFlags),
		level:  LevelDebug,
	}

	l.Debug("Debug")
	l.Info("Info")
	l.Warn("Warn")
	l.Error("Error")

	l.Panic("Panic")
	l.Fatal("Fatal")
}
