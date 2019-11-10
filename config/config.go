/*
@Time : 2019/11/11 1:14 上午
@Author : hechen
@File : config
@Software: GoLand
*/
package config

import (
	"errors"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Config struct {
	Path string
	File string
}

type Content interface {
}

func (config *Config) Load() (content *Content, err error) {
	if config.Path == "" {

	}
	if config.File == "" {
		config.File = "config.yaml"
	}
	fileContent, err := ioutil.ReadFile(config.Path + config.File)
	if err != nil {
		return content, errors.New("读取配置文件失败....(" + err.Error() + ")")
	}
	err = yaml.Unmarshal(fileContent, &content)
	if err != nil {
		return content, errors.New("解析配置文件失败....(" + err.Error() + ")")
	}
	return content, err
}
