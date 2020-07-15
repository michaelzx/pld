package pld_fs

import (
	"github.com/michaelzx/pld/pld_random"
	"strconv"
	"time"
)

func RandomFilename() string {
	return strconv.FormatInt(time.Now().UnixNano()/1000, 10) + "_" + pld_random.RandomNumStr(6)
}
