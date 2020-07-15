package pld_ueditor

import (
	"errors"
	"io"
	"math/rand"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func (ue *UEditor) UploadImage(r *http.Request) (*UploadRsp, error) {
	defer func() {
		r.Body.Close()
	}()
	// 获得文件
	file, fileHeader, err := ue.getUploadFile(r, ue.config.ImageFieldName)
	if err != nil {
		return nil, err
	}
	// 校验文件类型
	ext := filepath.Ext(fileHeader.Filename)
	if err = ue.checkFileType(ext, ue.config.ImageAllowFiles); err != nil {
		return nil, err
	}
	// 校验文件大小
	if err = ue.checkFileSize(fileHeader.Size, ue.config.ImageMaxSize); err != nil {
		return nil, err
	}

	webPath := ue.getFileWebPath(fileHeader.Filename, ue.config.ImagePathFormat)
	serverPath := filepath.Join(ue.webRoot, webPath)
	if err = ue.saveFile(file, serverPath); err != nil {
		return nil, err
	}
	return &UploadRsp{
		State:    stateSuccess,
		URL:      webPath,
		Title:    filepath.Base(webPath),
		Original: fileHeader.Filename,
		Type:     ext,
		Size:     fileHeader.Size,
	}, nil
}

func (ue *UEditor) getUploadFile(r *http.Request, fieldName string) (multipart.File, *multipart.FileHeader, error) {
	file, fileHeader, err := r.FormFile(fieldName)
	if err != nil {
		return nil, nil, err
	}
	if file == nil || fileHeader == nil {
		// 上传文件为空
		return nil, nil, errors.New(UPLOAD_FILE_IS_EMPTY)
	}
	return file, fileHeader, nil
}

// 校验文件大小
func (ue *UEditor) checkFileSize(fileSize int64, maxSize int) error {
	if fileSize > int64(maxSize) {
		return errors.New(ERROR_SIZE_EXCEED)
	}
	return nil
}
func (ue *UEditor) getFileWebPath(oriName, pathFormat string) string {
	timeNow := time.Now()
	timeNowFormat := time.Now().Format("2006_01_02_15_04_05")
	timeArr := strings.Split(timeNowFormat, "_")

	format := pathFormat

	format = strings.Replace(format, "{yyyy}", timeArr[0], 1)
	format = strings.Replace(format, "{mm}", timeArr[1], 1)
	format = strings.Replace(format, "{dd}", timeArr[2], 1)
	format = strings.Replace(format, "{hh}", timeArr[3], 1)
	format = strings.Replace(format, "{ii}", timeArr[4], 1)
	format = strings.Replace(format, "{ss}", timeArr[5], 1)

	timestamp := strconv.FormatInt(timeNow.UnixNano(), 10)
	format = strings.Replace(format, "{time}", string(timestamp), 1)

	pattern := "{rand:(\\d)+}"
	if ok, _ := regexp.MatchString(pattern, format); ok {
		// 生成随机字符串
		exp, _ := regexp.Compile(pattern)
		randLenStr := exp.FindSubmatch([]byte(format))

		randLen, _ := strconv.Atoi(string(randLenStr[1]))
		randStr := strconv.Itoa(rand.Int())
		randStrLen := len(randStr)
		if randStrLen > randLen {
			randStr = randStr[randStrLen-randLen:]
		}
		// 将随机传替换到format中
		format = exp.ReplaceAllString(format, randStr)
	}

	ext := filepath.Ext(oriName)

	return format + ext
}

// 校验文件类型
func (ue *UEditor) checkFileType(fileType string, allowTypes []string) error {
	valid := false
	for _, fileTypeValid := range allowTypes {
		if strings.ToLower(fileType) == fileTypeValid {
			valid = true
			break
		}
	}
	if !valid {
		return errors.New(ERROR_TYPE_NOT_ALLOWED)
	}
	return nil
}
func (ue *UEditor) saveFile(srcFile io.Reader, destFilePath string) error {

	fileDir := filepath.Dir(destFilePath)
	exists, err := checkPathExists(fileDir)
	if err != nil {
		return errors.New(ERROR_FILE_STATE)
	}

	if !exists {
		// 文件夹不存在，创建
		if err = os.MkdirAll(fileDir, 0766); err != nil {
			return errors.New(ERROR_CREATE_DIR)
		}
	}

	dstFile, err := os.OpenFile(destFilePath, os.O_WRONLY|os.O_RDONLY|os.O_CREATE, 0666)
	if err != nil {
		return errors.New(ERROR_DIR_NOT_WRITEABLE)
	}
	defer func() {
		dstFile.Close()
	}()
	_, err = io.Copy(dstFile, srcFile)
	if err != nil {
		return errors.New(ERROR_WRITE_CONTENT)
	}
	return nil
}
func checkPathExists(path string) (bool, error) {
	// 获取path的信息
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}

	if os.IsNotExist(err) {
		return false, nil
	}

	return false, err
}
