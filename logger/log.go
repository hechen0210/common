/**
@Time : 2019/11/19 10:28
@Author : hechen
@File : log
@Software: GoLand
*/
package logger

import (
	"common/helper"
	"os"
	"strconv"
	"strings"
	"time"
)

type Log struct {
	Path     string
	FileName string
	Rotate
}

type FileLogger struct {
	path        string
	name        string
	baseContent string
	rotate      Rotate
}

type MongoLogger struct {
	DbName string
	Rotate
}

type MysqlLogger struct {
	DbName string
	Rotate
}

/**
切割周期
*/
type Rotate struct {
	Type   string
	Period string
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

/**
新日志文件
*/
func (l Log) NewFileLogger() (logger *FileLogger, err error) {
	var paths []string
	var _logger FileLogger
	if l.Path == "" {
		l.Path = helper.GetAbsPath(os.Args[0])
	}
	paths = append(paths, strings.TrimSuffix(l.Path, "/"))
	_logger.path = strings.Join(paths, "/")
	_logger.name = l.FileName
	_logger.rotate = l.Rotate
	//写入初始化
	_logger.Write("======日志初始化完成======")
	return &_logger, nil
}

/**
按周期切割
*/
func (r *Rotate) getRotate() string {
	if r.Period == Weekly {
		year, week := time.Now().ISOWeek()
		return strconv.Itoa(year) + strconv.Itoa(week)
	} else if r.Period == Monthly {
		return time.Now().Format("200601")
	} else if r.Period == Yearly {
		return time.Now().Format("2006")
	}
	return time.Now().Format("20060102")
}
