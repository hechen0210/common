/*
@Time : 2019/11/16 3:00 上午
@Author : hechen
@File : section
@Software: GoLand
*/
package config

import (
	"github.com/shopspring/decimal"
	"strconv"
	"strings"
)

type Section struct {
	name string
	data map[string]Item
}

/**
获取配置内容
*/
func (c *Section) Get(key ...string) *Item {
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

func (c *Section) ToMString() map[string]string {
	result := make(map[string]string)
	if len(c.data) != 0 {
		for index, item := range c.data {
			switch item.DataType {
			case "string":
				result[index] = item.Data.(string)
			case "bool":
				result[index] = strconv.FormatBool(item.Data.(bool))
			case "int":
				result[index] = strconv.Itoa(item.Data.(int))
			case "float64":
				result[index] = decimal.NewFromFloat(item.Data.(float64)).String()
			}
		}
	}
	return result
}

func (c *Section) ToMInt() map[string]int {
	result := make(map[string]int)
	if len(c.data) != 0 {
		for index, item := range c.data {
			switch item.DataType {
			case "string":
				num, err := strconv.Atoi(item.Data.(string))
				if err != nil {
					num = 0
				}
				result[index] = num
			case "bool":
				if item.Data.(bool) {
					result[index] = 1
				} else {
					result[index] = 0
				}
			case "int":
				result[index] = item.Data.(int)
			case "float64":
				result[index] = int(decimal.NewFromFloat(item.Data.(float64)).IntPart())
			}
		}
	}
	return result
}

func (c *Section) ToMFloat() map[string]float64 {
	result := make(map[string]float64)
	if len(c.data) != 0 {
		for index, item := range c.data {
			switch item.DataType {
			case "string":
				num, err := strconv.ParseFloat(item.Data.(string), 64)
				if err != nil {
					num = 0
				}
				result[index] = num
			case "bool":
				if item.Data.(bool) {
					result[index] = 1
				} else {
					result[index] = 0
				}
			case "int":
				result[index] = float64(item.Data.(int))
			case "float64":
				result[index] = item.Data.(float64)
			}
		}
	}
	return result
}

func (c *Section) ToMBool() map[string]bool {
	result := make(map[string]bool)
	if len(c.data) != 0 {
		for index, item := range c.data {
			switch item.DataType {
			case "string":
				_bool, err := strconv.ParseBool(item.Data.(string))
				if err != nil {
					_bool = false
				}
				result[index] = _bool
			case "bool":
				result[index] = item.Data.(bool)
			case "int":
				result[index] = item.Data.(int) == 1
			case "float64":
				result[index] = false
			}
		}
	}
	return result
}
