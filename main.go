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
	log, err := logger.Log{
		Path:     "./",
		Dir:      "logs",
		FileName: "log.log",
		Rotate: logger.Rotate{
			Period:   logger.Daily,
			RotateBy: logger.Dir,
		},
		Sync:    false,
	}.New()
	fmt.Println(log, err)
}
