/*
@Time : 2019/11/11 1:14 上午
@Author : hechen
@File : config
@Software: GoLand
*/
package config

import (
	"common/helper"
	"errors"
	"fmt"
	"gopkg.in/ffmt.v1"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"reflect"
	"strings"
)

type Config struct {
	Path string
	File string
}

type Item struct {
	DataType string
	Data     interface{}
}
type Content map[interface{}]Item

var configContent = make(map[interface{}]Item)

/**
加载配置文件
*/
func (config *Config) Load() (data *map[interface{}]Item, err error) {
	if config.Path == "" {
		config.Path = helper.GetAbsPath(os.Args[0])
	}
	if config.File == "" {
		config.File = "config.yaml"
	}
	var fullPath string
	if strings.HasSuffix(config.Path, "/") {
		fullPath = config.Path + config.File
	} else {
		fullPath = config.Path + "/" + config.File
	}
	fileContent, err := ioutil.ReadFile(fullPath)
	if err != nil {
		return data, errors.New("读取配置文件失败....(" + err.Error() + ")")
	}
	var _content map[interface{}]interface{}
	err = yaml.Unmarshal(fileContent, &_content)
	if err != nil {
		return data, errors.New("解析配置文件失败....(" + err.Error() + ")")
	}
	parse(_content, "")
	ffmt.Print(configContent)
	return &configContent, err
}

func parse(config map[interface{}]interface{}, prefix string) {
	for key, item := range config {
		dataType := reflect.TypeOf(item).String()
		var index string
		if prefix != "" {
			index = prefix + "." + key.(string)
		} else {
			index = key.(string)
		}
		fmt.Println(item)
		configContent[index] = Item{
			DataType: dataType,
			Data:     item,
		}
		switch itemType := item.(type) {
		case map[interface{}]interface{}:
			parse(itemType, index)
		}
	}
}

/**
获取配置内容
*/
func (configContent *Content) Get(key string) {
	//	//fmt.Println(configContent)
	//	//fields := strings.Split(key, ".")
	//	//fmt.Println(fields)
	//content := *configContent
	//fmt.Println(content[fields[0]])
}
