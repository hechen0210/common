/*
@Time : 2019/11/17 4:24 上午
@Author : hechen
@File : file
@Software: GoLand
*/
package helper

import "os"

func FileExist(fileName string) bool {
	var exist = true
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		exist = false
	}
	return exist
}
