/*
@Time : 2019/11/10 8:39 下午
@Author : hechen
@File : main
@Software: GoLand
*/
package main

import (
	"common/config"
	"common/logger"
	"common/mysql"
	"fmt"
)

func main() {
	c := config.ConfigFile{
		Path: "./config",
		File: "",
	}
	data := c.Load()
	config := data.GetSection("urls")
	a := config.Get("inside").ToSString()
	for key, item := range a {
		fmt.Println(key, item)
	}
	fileLogger, err := logger.Log{
		Path:     "./logs",
		FileName: "log.log",
		Rotate: logger.Rotate{
			Type:   logger.Dir,
			Period: logger.Daily,
		},
	}.NewFileLogger()
	fileLogger.SetBaseContent("进入结算").Write("订单号")
	fmt.Println(fileLogger, err)
	mc, err := mysql.Config{
		Host:          "127.0.0.1",
		User:          "root",
		Password:      "",
		Port:          "3306",
		DbName:        "new_ysb",
		Prefix:        "ysb_",
		SingularTable: false,
	}.New()
	if err != nil {
		fileLogger.SetBaseContent("").Write(err.Error())
	}
	fmt.Println(mc, err)
}
