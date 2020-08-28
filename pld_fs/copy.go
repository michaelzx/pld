package pld_fs

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

/**
 * 拷贝文件夹,同时拷贝文件夹中的文件
 * @param fromPath   		需要拷贝的文件夹路径: D:/test
 * @param toPath		拷贝到的位置: D:/backup/
 */
func CopyDir(fromPath string, toPath string) error {
	// 检测目录正确性
	if fromInfo, err := os.Stat(fromPath); err != nil {
		fmt.Println(err.Error())
		return err
	} else {
		if !fromInfo.IsDir() {
			e := errors.New("fromPath 不是一个正确的目录！")
			return e
		}
	}
	if toInfo, err := os.Stat(toPath); err != nil {
		fmt.Println(err.Error())
		return err
	} else {
		if !toInfo.IsDir() {
			e := errors.New("toInfo不是一个正确的目录！")
			return e
		}
	}

	err := filepath.Walk(fromPath, func(filePath string, f os.FileInfo, err error) error {
		fmt.Println(filePath)
		if f == nil {
			return err
		}
		newPath := strings.Replace(filePath, fromPath, filepath.Join(toPath, filepath.Base(fromPath)), 1)
		if !f.IsDir() {
			fmt.Println("-----------------")
			fmt.Println(filePath)
			fmt.Println("↓")
			fmt.Println(newPath)
			fmt.Println("-----------------")
			CopyFile(filePath, newPath)
		} else {
			CreateIfNotExist(newPath)
		}
		return nil
	})
	if err != nil {
		fmt.Printf(err.Error())
	}
	return err
}

// 生成目录并拷贝文件
func CopyFile(fromPath, toPath string) (w int64, err error) {
	fromFile, err := os.Open(fromPath)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer fromFile.Close()
	// 为文件创建目录
	toDirPath := filepath.Dir(toPath)
	if !Exists(toDirPath) {
		// 创建目录
		CreateIfNotExist(toDirPath)
	}
	dstFile, err := os.Create(toPath)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer dstFile.Close()

	return io.Copy(dstFile, fromFile)
}
