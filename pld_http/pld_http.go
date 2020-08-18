package pld_http

import (
	"errors"
	"github.com/michaelzx/pld/pld_fs"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

// 如果文件有存在，则先删除
func deleteIfExists(filePath string) error {
	if pld_fs.Exists(filePath) {
		return os.Remove(filePath)
	}
	return nil
}

func DownloadFile(url string, tempPath, savePath string) error {
	// ******************************************************
	// 下载的临时文件名称
	// ******************************************************
	tmpFilePath := tempPath + ".download" // 没下载成功之前，用download后缀作为临时文件
	tmpFileDirPath := filepath.Dir(tmpFilePath)
	pld_fs.CreateIfNotExist(tmpFileDirPath)
	// 如果临时文件已存在，则删除
	err := deleteIfExists(tmpFilePath)
	if err != nil {
		return err
	}
	// ******************************************************
	// 创建一个http client
	// ******************************************************
	client := new(http.Client)
	// client.Timeout = time.Second * 60 //设置超时时间
	resp, err := client.Get(url)
	if err != nil {
		return err
	}
	// ******************************************************
	// 读取服务器返回的文件大小
	// ******************************************************
	// fSize, err = strconv.ParseInt(resp.Header.Get("Content-Length"), 10, 32)
	// if err != nil {
	// 	return err
	// }
	// ******************************************************
	// 创建并写入临时文件
	// ******************************************************
	if resp.Body == nil {
		return errors.New("body is null")
	}
	defer resp.Body.Close()

	outFile, err := os.Create(tmpFilePath)
	if err != nil {
		return err
	}
	defer outFile.Close()

	// 将响应流和文件流对接起来
	_, err = io.Copy(outFile, resp.Body)
	if err != nil {
		return err
	}
	err = outFile.Close()
	if err != nil {
		return err
	}
	// ******************************************************
	// 用临时文件替换正式文件
	// ******************************************************
	err = os.Rename(tmpFilePath, savePath)
	if err != nil {
		return err
	}
	return nil
}
