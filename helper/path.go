/*
@Time : 2019/11/11 1:32 上午
@Author : hechen
@File : path
@Software: GoLand
*/
package helper

import (
	"os"
	"path/filepath"
	"strings"
)

/**
获取路径绝对地址
*/
func GetAbsPath(path string) string {
	dir, err := filepath.Abs(filepath.Dir(path))
	if err != nil {
		return ""
	}
	return strings.Replace(dir, "\\", "/", -1)
}

/**
判断文件夹是否存在
*/
func PathExist(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}
