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

/*
Config 配置参数
Use file 文件，env 环境变量，chiefByFile 合并但以配置文件为主 chiefByEnv 合并但以环境变量为主
*/
type Config struct {
	FileName string
	Use      string // file,env,chiefByFile,chiefByEnv
	Env      Env
}

type Env struct {
	Prefix       string
	IgnorePrefix bool
}

type ConfigFile struct {
	Path string
	File string
}

type ConfigData struct {
	configFile string
	data       map[string]Item
}

var UseType = []string{"file", "env", "chiefByFile", "chiefByEnv"}

/*
Load 加载配置文件
@param readEnv 读取环境变量，如果为true，优先读取环境变量
@param envPrefix 环境变量前缀
@param ignorePrefix 忽略环境变量前缀
*/
func Load(config *Config) *ConfigData {
	if !helper.Contains(UseType, config.Use) {
		fmt.Println("use 类型错误,只能使用file,env,chiefByFile,chiefByEnv")
		os.Exit(1)
	}
	envData := config.loadByEnv()
	if config.Use == "env" {
		return envData
	}
	fileData := config.loadByFile()
	if config.Use == "file" {
		return fileData
	}
	configData := ConfigData{configFile: fileData.configFile}
	data := fileData.data
	for key, item := range envData.data {
		if _, exist := fileData.data[key]; exist {
			if config.Use == "chiefByEnv" {
				data[key] = item
			}
		} else {
			data[key] = item
		}
	}
	configData.data = data
	return &configData
}

/**
解析配置文件
*/
func (c *Config) ParseFile() *ConfigFile {
	if c.FileName == "" {
		fmt.Println("配置文件不能为空")
		os.Exit(1)
	}
	var parseFile = strings.Split(c.FileName, "/")
	parseFileLen := len(parseFile)
	if parseFileLen == 1 {
		return &ConfigFile{
			Path: "",
			File: parseFile[parseFileLen-1],
		}
	}
	return &ConfigFile{
		Path: strings.Join(parseFile[0:parseFileLen-1], "/"),
		File: parseFile[parseFileLen-1],
	}
}

/*
loadByEnv 从环境变量读取配置
*/
func (c *Config) loadByEnv() *ConfigData {
	configData := make(map[string]Item)
	env := os.Environ()
	for _, item := range env {
		_env := strings.Split(item, "=")
		key := strings.Replace(_env[0], "_", ".", -1)
		value := _env[1]
		valueType := reflect.TypeOf(value).String()
		if c.Env.Prefix != "" {
			if strings.HasPrefix(key, c.Env.Prefix) {
				if c.Env.IgnorePrefix {
					key = strings.TrimPrefix(key, c.Env.Prefix)
				}
				configData[key] = Item{
					DataType: valueType,
					Data:     value,
				}
			}
		} else {
			configData[key] = Item{
				DataType: valueType,
				Data:     value,
			}
		}
	}
	return &ConfigData{
		configFile: "",
		data:       configData,
	}
}

/*
loadByFile 从文件加载配置
*/
func (config *Config) loadByFile() *ConfigData {
	configFile := config.ParseFile()
	fullPath := configFile.getFilePath()
	fileContent, err := ioutil.ReadFile(fullPath)
	if err != nil {
		fmt.Println("读取配置文件失败....(" + err.Error() + ")")
		os.Exit(1)
	}
	var configs = ConfigData{
		configFile: fullPath,
		data:       map[string]Item{},
	}
	var content map[interface{}]interface{}
	err = yaml.Unmarshal(fileContent, &content)
	if err != nil {
		fmt.Println("解析配置文件失败....(" + err.Error() + ")")
		os.Exit(1)
	}
	configs.parse(content, "")
	fmt.Println("配置文件加载完成")
	return &configs
}

/*
getFilePath 获取配置文件完整路径
*/
func (c *ConfigFile) getFilePath() string {
	var paths []string
	if c.Path == "" {
		c.Path = helper.GetAbsPath()
	}
	paths = append(paths, strings.TrimSuffix(c.Path, "/"))
	if c.File == "" {
		c.File = "config.yaml"
	}
	paths = append(paths, c.File)
	fullPath := strings.Join(paths, "/")
	return fullPath
}

/**
解析配置文件
*/
func (c *ConfigData) parse(content map[interface{}]interface{}, prefix string) {
	for key, item := range content {
		dataType := reflect.TypeOf(item).String()
		index := setIndex(key, prefix)
		switch data := item.(type) {
		case map[interface{}]interface{}:
			c.parse(data, index)
		default:
			c.data[index] = Item{
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
