package pld_ueditor

type UploadRsp struct {
	State    string `json:"state"`
	URL      string `json:"url"`
	Title    string `json:"title"`    // 文件名
	Original string `json:"original"` // 原始名称
	Type     string `json:"type"`
	Size     int64  `json:"size"`
}
