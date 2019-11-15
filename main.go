/*
@Time : 2019/11/10 8:39 下午
@Author : hechen
@File : main
@Software: GoLand
*/
package main

import (
	"common/config"
	"fmt"
)

func main() {
	c := config.Config{
		Path: "./config",
		File: "",
	}
	config, err := c.Load()
	fmt.Println(err,config)
	//config.Get("mysql.read")
}
