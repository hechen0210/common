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
func GetAbsPath() string {
	dir, err := os.Executable()
	if err != nil {
		return ""
	}
	path := filepath.Dir(dir)
	return strings.Replace(path, "\\", "/", -1)
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
