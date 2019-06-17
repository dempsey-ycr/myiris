package tool

import (
	"fmt"
	"os"
	"strings"
)

// 判断目录是否存在
func DirExists(dir string) bool {
	_, err := os.Stat(dir)
	return err == nil || os.IsExist(err)
}

// 判断所给路径是否为文件夹
func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

// 判断所给路径是否为文件
func IsFile(path string) bool {
	return !IsDir(path)
}

// 目录不存在则创建并返回，存在返回原路径
func CreatePath(path string) string {
	var dir = path
	if IsFile(path) {
		dir = path[:strings.LastIndex(path, "/")]
	}

	if DirExists(dir) {
		return path
	}

	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		fmt.Println("make dir all error", err)
		return ""
	}
	return path
}
