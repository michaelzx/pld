package pld_ueditor

// Scrawl 未实现
// Catcher 未实现
func (ue *UEditor) UploadImage(ctx interface{}) (rsp *UploadRsp, err error) {
	fileHeader, err := GetFileHeader(ctx, ue.config.ImageFieldName)
	if err != nil {
		return nil, err
	}
	return ue.uploader.Upload(fileHeader, ue.config.ImageMaxSize, ue.config.ImageAllowFiles, ue.config.ImagePathFormat)
}

func (ue *UEditor) UploadVideo(ctx interface{}) (rsp *UploadRsp, err error) {
	fileHeader, err := GetFileHeader(ctx, ue.config.VideoFieldName)
	if err != nil {
		return nil, err
	}
	return ue.uploader.Upload(fileHeader, ue.config.VideoMaxSize, ue.config.VideoAllowFiles, ue.config.VideoPathFormat)
}
func (ue *UEditor) UploadFile(ctx interface{}) (rsp *UploadRsp, err error) {
	fileHeader, err := GetFileHeader(ctx, ue.config.FileFieldName)
	if err != nil {
		return nil, err
	}
	return ue.uploader.Upload(fileHeader, ue.config.FileMaxSize, ue.config.FileAllowFiles, ue.config.FilePathFormat)
}
