package pld_ueditor

import (
	"os"
	"path/filepath"
	"strings"
)

type FileItem struct {
	Url   string `json:"url"`   // 文件url
	Mtime int64  `json:"mtime"` // 最后编辑时间
}

type ListImage struct {
	State string      `json:"state"` // 这些第一个字母要大写，否则不出结果
	List  []*FileItem `json:"list"`
	Start int         `json:"start"`
	Total int         `json:"total"`
	// Name        string
	// Age         int
	// Slices      []string //slice
	// Mapstring   map[string]string
	// StructArray []List            //结构体的切片型
	// MapStruct   map[string][]List //map:key类型是string或struct，value类型是切片，切片的类型是string或struct
	//	Desks  List
}

func (ue *UEditor) ListImage(start int, size int) *ListImage {
	fileList, err := ue.getFileList(start, size, ue.config.ImageAllowFiles)
	if err != nil {
		return &ListImage{
			State: err.Error(),
			List:  fileList,
			Start: start,
			Total: 0,
		}
	}
	if len(fileList) > 0 {
		return &ListImage{
			State: stateSuccess,
			List:  fileList,
			Start: start,
			Total: len(fileList),
		}
	} else {
		return &ListImage{
			State: stateNoMatchFile,
			List:  fileList,
			Start: start,
			Total: 0,
		}
	}
}
func (ue *UEditor) getFileList(start int, size int, allowTypes []string) (fileList []*FileItem, err error) {

	files := make([]*FileItem, 0)

	path := filepath.Join(ue.webRoot, ue.uploadDir)

	if err := ue.walkFiles(path, allowTypes, &files); err != nil {
		return nil, nil
	}

	if size == 0 {
		size = 20
	}

	end := start + size

	i := end
	listLen := len(files)
	if i > listLen {
		i = listLen
	}

	fileList = make([]*FileItem, 0)
	for i := i - 1; i < listLen && i >= 0 && i >= start; i-- {
		fileList = append(fileList, files[i])
	}

	return
}

// 递归获取文件列表
func (ue *UEditor) walkFiles(path string, allowFiles []string, files *[]*FileItem) (err error) {
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
						urlPath := next[len(ue.webRoot):]
						if strings.ToLower(ext) == allowItem {
							*files = append(*files, &FileItem{
								Url:   urlPath,
								Mtime: info.ModTime().Unix(),
							})
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
