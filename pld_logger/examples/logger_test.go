package examples

import (
	"github.com/michaelzx/pld/pld_logger"
	"github.com/michaelzx/pld/pld_logger/stdlog"
	"os"
	"testing"
)

type Foo struct {
	ID       int64
	SiteID   int64
	Title    string
	FormType int32
}

var foo = &Foo{
	ID:       1,
	SiteID:   2,
	Title:    "xxxxxx",
	FormType: 999,
}

func TestStdLog(t *testing.T) {
	pld_logger.UseStdLog(os.Stderr, stdlog.LevelDebug, true)
	pld_logger.Debug("Debug", *foo)
}

func TestZapLog(t *testing.T) {
	pld_logger.UseZapLog(true, "")
	pld_logger.Debug("Debug", foo)
}
