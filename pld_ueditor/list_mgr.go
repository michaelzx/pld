package pld_ueditor

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type ListMgr struct {
	webRoot   string
	uploadDir string
}

func NewListMgr(webRoot string, uploadDir string) *ListMgr {
	return &ListMgr{webRoot: webRoot, uploadDir: uploadDir}
}

func (lm *ListMgr) getFileList(start int, size int, allowTypes []string) (fileList *FileList) {
	fmt.Println("allowTypes", allowTypes)
	files := make([]FileItem, 0)

	path := filepath.Join(lm.webRoot, lm.uploadDir)

	if err := lm.walkFiles(path, allowTypes, &files); err != nil {
		return &FileList{
			State: err.Error(),
			List:  nil,
			Start: start,
			Total: 0,
		}
	}
	fmt.Println("files", files)
	if size == 0 {
		size = 20
	}

	end := start + size

	i := end
	listLen := len(files)
	if i > listLen {
		i = listLen
	}

	items := make([]FileItem, 0, 0)
	for i := i - 1; i < listLen && i >= 0 && i >= start; i-- {
		items = append(items, files[i])
	}
	if len(items) == 0 {
		return &FileList{
			State: stateNoMatchFile,
			List:  items,
			Start: start,
			Total: 0,
		}
	}
	return &FileList{
		State: stateSuccess,
		List:  items,
		Start: start,
		Total: len(items),
	}
}

// 递归获取文件列表
func (lm *ListMgr) walkFiles(path string, allowFiles []string, files *[]FileItem) (err error) {
	fileInfo, err := os.Stat(path)
	if err == nil {
		if !fileInfo.IsDir() {
			return
		}
		filepath.Walk(path, func(fileName string, info os.FileInfo, err error) error {
			if fileName != "." && fileName != ".." && fileName != "walk:.DS_Store" && fileName != path {
				next := fileName
				if !info.IsDir() {
					ext := filepath.Ext(fileName)
					for _, allowItem := range allowFiles {
						urlPath := next[len(lm.webRoot):]
						if strings.ToLower(ext) == allowItem {
							fmt.Println("fileName", fileName)
							*files = append(*files, FileItem{
								Url:   urlPath,
								Mtime: info.ModTime().Unix(),
							})
							fmt.Println("files", files)
							break
						}
					}
				}
			}

			return nil
		})
		return
	}
	return
}
