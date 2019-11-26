/*
@Time : 2019/11/10 8:39 下午
@Author : hechen
@File : main
@Software: GoLand
*/
package main

import (
	"fmt"
	"github.com/hechen0210/common/config"
	"github.com/hechen0210/common/logger"
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
	fileLogger.SetBaseContent("进入结算").Write("订单号").Print()
	fileLogger.Write("aiyaya")
	fmt.Println(err)
	//mc := mysql.Config{
	//	Host:          "192.168.0.235",
	//	User:          "root",
	//	Password:      "ysb123456",
	//	Port:          "3306",
	//	DbName:        "new_ysb",
	//	Prefix:        "ysb_",
	//	SingularTable: false,
	//}.New()
	//if err != nil {
	//	fileLogger.SetBaseContent("").Write(err.Error())
	//}
	//var cc []string
	//r := mc.Db.Raw("select * from ysb_users12").Scan(&cc)
	//fmt.Println(r.Error)
}
