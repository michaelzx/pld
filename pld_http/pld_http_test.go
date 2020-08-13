package pld_http

import (
	"fmt"
	"testing"
)

func TestDownloadFile(t *testing.T) {
	err := DownloadFile("https://www.baidu.com/img/PCfb_5bf082d29588c07f842ccde3f97243ea.png",
		"./test.png", "./test.png")
	if err != nil {
		fmt.Println(err)
	}
}
