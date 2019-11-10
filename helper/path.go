/*
@Time : 2019/11/11 1:32 上午
@Author : hechen
@File : path
@Software: GoLand
*/
package helper

import (
	"path/filepath"
	"strings"
)

func GetAbsPath(path string) string {
	dir, err := filepath.Abs(filepath.Dir(path))
	if err != nil {
		return ""
	}
	
	return strings.Replace(dir, "\\", "/", -1)
}
