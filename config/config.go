/*
@Time : 2019/11/11 1:14 上午
@Author : hechen
@File : config
@Software: GoLand
*/
package config

import (
	"fmt"
	"github.com/hechen0210/common/helper"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"reflect"
	"strconv"
	"strings"
)

type ConfigFile struct {
	Path string
	File string
}

type ConfigData struct {
	configFile string
	data       map[string]Item
}

var configs = ConfigData{
	configFile: "",
	data:       map[string]Item{},
}

/**
加载配置文件
*/
func (config *ConfigFile) Load() *ConfigData {
	var paths []string
	if config.Path == "" {
		config.Path = helper.GetAbsPath()
	}
	paths = append(paths, strings.TrimSuffix(config.Path, "/"))
	if config.File == "" {
		config.File = "config.yaml"
	}
	paths = append(paths, config.File)
	fullPath := strings.Join(paths, "/")
	fileContent, err := ioutil.ReadFile(fullPath)
	if err != nil {
		fmt.Println("读取配置文件失败....(" + err.Error() + ")")
		os.Exit(1)
	}
	configs.configFile = fullPath
	var _content map[interface{}]interface{}
	err = yaml.Unmarshal(fileContent, &_content)
	if err != nil {
		fmt.Println("解析配置文件失败....(" + err.Error() + ")")
		os.Exit(1)
	}
	parse(_content, "")
	fmt.Println("配置文件加载完成")
	return &configs
}

/**
解析配置文件
*/
func parse(config map[interface{}]interface{}, prefix string) {
	for key, item := range config {
		dataType := reflect.TypeOf(item).String()
		index := setIndex(key, prefix)
		switch data := item.(type) {
		case map[interface{}]interface{}:
			parse(data, index)
		default:
			configs.data[index] = Item{
				DataType: dataType,
				Data:     item,
			}
		}
	}
}

func setIndex(key interface{}, prefix string) string {
	var index string
	if reflect.TypeOf(key).String() == "int" {
		index = strconv.Itoa(key.(int))
	} else {
		index = key.(string)
	}
	if prefix != "" {
		index = prefix + "." + index
	}
	return index
}

/**
获取节点内容
*/
func (c *ConfigData) GetSection(key ...string) *Section {
	realKey := strings.Join(key, ".")
	data := map[string]Item{}
	for index, item := range c.data {
		if strings.HasPrefix(index, realKey) {
			data[strings.TrimPrefix(index, realKey+".")] = item
		}
	}
	return &Section{
		name: realKey,
		data: data,
	}
}

func (c *ConfigData) Get(key ...string) *Item {
	realKey := strings.Join(key, ".")
	if item, exist := c.data[realKey]; exist {
		return &item
	} else {
		for index, item := range c.data {
			if strings.HasSuffix(index, "."+realKey) {
				return &item
			}
		}
	}
	return &Item{
		DataType: realKey,
		Data:     nil,
	}
}
