package pld_ueditor

type FileItem struct {
	Url   string `json:"url"`   // 文件url
	Mtime int64  `json:"mtime"` // 最后编辑时间
}

type FileList struct {
	State string     `json:"state"` // 这些第一个字母要大写，否则不出结果
	List  []FileItem `json:"list"`
	Start int        `json:"start"`
	Total int        `json:"total"`
	// Name        string
	// Age         int
	// Slices      []string //slice
	// Mapstring   map[string]string
	// StructArray []List            //结构体的切片型
	// MapStruct   map[string][]List //map:key类型是string或struct，value类型是切片，切片的类型是string或struct
	//	Desks  List
}

func (ue *UEditor) ListImage(start int, size int) *FileList {
	return ue.listMgr.getFileList(start, size, ue.config.ImageAllowFiles)
}
func (ue *UEditor) ListFile(start int, size int) *FileList {
	return ue.listMgr.getFileList(start, size, ue.config.FileAllowFiles)
}
func (ue *UEditor) ListVideo(start int, size int) *FileList {
	return ue.listMgr.getFileList(start, size, ue.config.VideoAllowFiles)
}
