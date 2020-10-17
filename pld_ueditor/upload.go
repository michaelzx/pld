package pld_ueditor

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gofiber/fiber/v2"
	"github.com/michaelzx/pld/pld_fs"
	"github.com/valyala/fasthttp"
	"math/rand"
	"mime/multipart"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func (ue *UEditor) UploadImage4Fiber(ctx *fiber.Ctx) (*UploadRsp, error) {
	// 获得文件
	fileHeader, err := ctx.FormFile(ue.config.ImageFieldName)
	if err != nil {
		return nil, err
	}
	return ue.toSave(fileHeader)

}

func (ue *UEditor) UploadImage4Gin(ctx *gin.Context) (*UploadRsp, error) {
	fileHeader, err := ctx.FormFile(ue.config.ImageFieldName)
	if err != nil {
		return nil, err
	}
	return ue.toSave(fileHeader)
}

func (ue *UEditor) toSave(fileHeader *multipart.FileHeader) (rsp *UploadRsp, err error) {
	// 校验文件大小
	if err = ue.checkFileSize(fileHeader.Size, ue.config.ImageMaxSize); err != nil {
		err = fmt.Errorf("文件大小不符合规则:%w", err)
		return
	}
	// 校验文件类型
	ext := filepath.Ext(fileHeader.Filename)
	err = ue.checkFileType(ext, ue.config.ImageAllowFiles)
	if err != nil {
		err = fmt.Errorf("文件类型不符合规则:%w", err)
		return
	}
	webPath := ue.getFileWebPath(fileHeader.Filename, ue.config.ImagePathFormat)
	serverPath := filepath.Join(ue.webRoot, webPath)
	pld_fs.CreateIfNotExist(filepath.Dir(serverPath))
	err = fasthttp.SaveMultipartFile(fileHeader, serverPath)
	if err != nil {
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

// 校验文件大小
func (ue *UEditor) checkFileSize(fileSize int64, maxSize int) error {
	if fileSize == 0 {
		return errors.New(UPLOAD_FILE_IS_EMPTY)
	}
	if fileSize > int64(maxSize) {
		return errors.New(ERROR_SIZE_EXCEED)
	}
	return nil
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

// 文件路径
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
