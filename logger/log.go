/*
@Time : 2019/11/17 2:37 上午
@Author : hechen
@File : log
@Software: GoLand
*/
package logger

import (
	"common/helper"
	"os"
	"path"
	"strconv"
	"strings"
	"time"
)

type Log struct {
	Path        string
	Dir         string
	FileName    string
	Rotate      Rotate
	LogType     []string
	Sync        bool
	CurrentPath string
	CurrentFile string
}

/**
切割周期
*/
type Rotate struct {
	Period   string
	RotateBy string
}

const (
	Daily   = "Daily"   //按日切割
	Weekly  = "Weekly"  //按周切割
	Monthly = "Monthly" //按月切割
	Yearly  = "Yearly"  //按年切割
	Dir     = "Dir"     //切割文件夹
	File    = "File"    //切割文件
	Info    = "Info"    //信息
	Errors  = "Error"   //错误
	Warning = "Warning" //警告
)

func (l Log) New() (logger *Log, err error) {
	if err = l.setPath(); err != nil {
		return nil, err
	}
	if err = l.setFile(); err != nil {
		return nil, err
	}
	return &l, nil
}

func (l *Log) AddLogType(logType string) (logger *Log, err error) {
	l.LogType = append(l.LogType, logType)
	return l, l.setFile()
}

/**
检查路径,不存在则创建
*/
func (l *Log) setPath() error {
	var paths []string
	if l.Path == "" {
		l.Path = helper.GetAbsPath(os.Args[0])
	}
	paths = append(paths, strings.TrimSuffix(l.Path, "/"))
	if l.Dir == "" {
		l.Dir = "logs"
	}
	paths = append(paths, strings.Trim(l.Dir, "/"))
	if l.Rotate.RotateBy == Dir {
		paths = append(paths, l.getRotate())
	}
	fullPath := strings.Join(paths, "/")
	if !helper.PathExist(fullPath) {
		return os.MkdirAll(fullPath, 0755)
	}
	l.CurrentPath = fullPath
	return nil
}

/**
设置文件
*/
func (l *Log) setFile() (err error) {
	var fileName []string
	if l.FileName == "" {
		l.FileName = "log.log"
	}
	fileExt := path.Ext(l.FileName)
	fileName = append(fileName, strings.TrimSuffix(l.FileName, fileExt))
	if l.LogType == nil {
		l.LogType = append(l.LogType, Info, Errors, Warning)
	}
	if l.Rotate.RotateBy == File {
		fileName = append(fileName, l.getRotate())
	}
	fileNameLen := len(fileName)
	var fullName string
	for _, item := range l.LogType {
		if fileNameLen == 1 {
			fullName = strings.Join([]string{fileName[0], item}, "_")
		} else {
			fullName = strings.Join([]string{fileName[0], item, fileName[1]}, "_")
		}
		fullName = fullName + fileExt
		if err = createFile(l.CurrentPath + "/" + fullName); err != nil {
			return err
		}
	}
	return nil
}

/**
按周期切割
*/
func (l *Log) getRotate() string {
	if l.Rotate.Period == Weekly {
		_, week := time.Now().ISOWeek()
		return strconv.Itoa(week)
	} else if l.Rotate.Period == Monthly {
		return time.Now().Format("200601")
	} else if l.Rotate.Period == Yearly {
		return time.Now().Format("2006")
	}
	return time.Now().Format("20060102")
}

/**
创建日志文件
*/
func createFile(filePath string) error {
	if !helper.FileExist(filePath) {
		file, err := os.Create(filePath)
		err = file.Close()
		return err
	}
	return nil
}

/**
切割文件
*/
func (l *Log) rotateLog() error {
	if l.Rotate.RotateBy == Dir {
		if err := l.setPath(); err != nil {
			return err
		}
	}
	if err := l.setFile(); err != nil {
		return err
	}
	return nil
}
