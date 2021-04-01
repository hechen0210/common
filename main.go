/*
 * @Author: your name
 * @Date: 2021-03-17 15:32:45
 * @LastEditTime: 2021-03-17 15:50:23
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /common/main.go
 */
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
)

func main() {
	Config,err := config.Load(&config.Config{
		FileName: "./config/config.yaml",
		Use:      "env",
		Env: config.Env{
			Prefix:       "app_config",
			IgnorePrefix: true,
		},
	})
	fmt.Println(err)
	fmt.Println(Config)
	//c := config.ConfigFile{
	//	Path: "./config",
	//	File: "",
	//}
	//data := c.Load()
	//config := data.GetSection("urls")
	//a := config.Get("inside").ToSString()
	//for key, item := range a {
	//	fmt.Println(key, item)
	//}
	//fileLogger, err := logger.Log{
	//	Path:     "./logs",
	//	FileName: "log.log",
	//	Rotate: logger.Rotate{
	//		Type:   logger.Dir,
	//		Period: logger.Daily,
	//	},
	//}.NewFileLogger()
	//fileLogger.SetBaseContent("进入结算").Write("订单号").Print()
	//fileLogger.Write("aiyaya")
	//fmt.Println(err)
	// mc := mysql.Config{
	// 	Host:          "127.0.0.1",
	// 	User:          "root",
	// 	Password:      "",
	// 	Port:          "3306",
	// 	DbName:        "7dayWechat",
	// 	Prefix:        "7day_",
	// 	SingularTable: false,
	// }.New()
	// var admin struct {
	// 	Id        int    `json:"id"`
	// 	Account   string `json:"account"`
	// 	Password  string `json:"-"`
	// 	Name      string `json:"name"`
	// 	CreatedAt string `json:"created_at"`
	// 	UpdatedAt string `json:"updated_at"`
	// }
	// c := mc.Client
	// err := c.Debug().Find(&admin)
	// fmt.Println(err)
	//if err != nil {
	//	fileLogger.SetBaseContent("").Write(err.Error())
	//}
	//var cc []string
	//r := mc.Client.Raw("select * from ysb_users12").Scan(&cc)
	//fmt.Println(r.Error)
}
