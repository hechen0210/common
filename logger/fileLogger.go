/**
@Time : 2019/11/19 18:57
@Author : hechen
@File : fileLogger
@Software: GoLand
*/
package logger

import (
	"fmt"
	"github.com/hechen0210/common/helper"
	"log"
	"os"
	"path"
	"strings"
)

func (l *FileLogger) SetBaseContent(content string) *FileLogger {
	l.baseContent = content
	return l
}

func (l *FileLogger) Print() *FileLogger {
	fmt.Println(l.LogContent)
	return l
}

func (l *FileLogger) Write(content string) *FileLogger {
	l.LogContent = l.baseContent + " " + content
	write(l.getLogfile(), l.LogContent)
	return l
}

func (l *FileLogger) getLogfile() string {
	var filePath string
	var fileName string
	rotate := l.rotate.getRotate()
	if l.rotate.Type == Dir {
		filePath = getDir(l.path, rotate)
		fileName = l.name
	} else {
		filePath = l.path
		fileName = getFileName(l.name, rotate)
	}
	return createFile(filePath + "/" + fileName)
}

/**
内容写入文件
*/
func write(fileName, content string) {
	file, _ := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, 0666)
	defer file.Close()
	logFile := log.New(file, "", log.LstdFlags|log.Lshortfile)
	logFile.Println(content)
}

func getDir(dir string, rotate string) string {
	if rotate != "" {
		dir = dir + "/" + rotate
	}
	return createDir(dir)
}
func getFileName(fileName string, rotate string) (fullName string) {
	var _fullName []string
	ext := path.Ext(fileName)
	name := strings.TrimSuffix(fileName, ext)
	_fullName = append(_fullName, name)
	if rotate != "" {
		_fullName = append(_fullName, rotate)
	}
	fullName = strings.Join(_fullName, "_")
	return fullName + ext
}

/**
创建目录
*/
func createDir(fullPath string) string {
	if !helper.PathExist(fullPath) {
		_ = os.MkdirAll(fullPath, 0755)
	}
	return fullPath
}

/**
创建文件
*/
func createFile(filePath string) string {
	if !helper.FileExist(filePath) {
		file, _ := os.Create(filePath)
		_ = file.Close()
	}
	return filePath
}
