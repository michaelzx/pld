package pld_ueditor

const (
	stateSuccess     = "SUCCESS" // 上传成功标记，在UEditor中内不可改变，否则flash判断会出错
	stateError       = "ERROR"
	stateNoMatchFile = "no match file"
)

const (
	BIGGER_THAN_UPLOAD_MAX_FILESIZE = "文件大小超出 upload_max_filesize 限制"
	BIGGER_THAN_MAX_FILE_SIZE       = "文件大小超出 MAX_FILE_SIZE 限制"
	FILE_NOT_COMPLETE               = "文件未被完整上传"
	NO_FILE_UPLOAD                  = "没有文件被上传"
	UPLOAD_FILE_IS_EMPTY            = "上传文件为空"
	ERROR_TMP_FILE                  = "临时文件错误"
	ERROR_TMP_FILE_NOT_FOUND        = "找不到临时文件"
	ERROR_SIZE_EXCEED               = "文件大小超出网站限制"
	ERROR_TYPE_NOT_ALLOWED          = "文件类型不允许"
	ERROR_CREATE_DIR                = "目录创建失败"
	ERROR_DIR_NOT_WRITEABLE         = "目录没有写权限"
	ERROR_FILE_MOVE                 = "文件保存时出错"
	ERROR_FILE_NOT_FOUND            = "找不到上传文件"
	ERROR_WRITE_CONTENT             = "写入文件内容错误"
	ERROR_UNKNOWN                   = "未知错误"
	ERROR_DEAD_LINK                 = "链接不可用"
	ERROR_HTTP_LINK                 = "链接不是http链接"
	ERROR_HTTP_CONTENTTYPE          = "链接contentType不正确"
	INVALID_URL                     = "非法 URL"
	INVALID_IP                      = "非法 IP"
	ERROR_BASE64_DATA               = "base64图片解码错误"
	ERROR_FILE_STATE                = "文件系统错误"
	ERRPR_READ_REMOTE_DATA          = "读取远程链接出错"
)
