package pld_fs

import "testing"

func TestCopyDir(t *testing.T) {
	CopyDir("/Users/michael/workspace/cirs/cirs-doc-keeper/test/A", "/Users/michael/workspace/cirs/cirs-doc-keeper/test/移动测试目录")
}
