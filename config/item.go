/*
@Time : 2019/11/16 4:22 下午
@Author : hechen
@File : item
@Software: GoLand
*/
package config

import (
	"github.com/shopspring/decimal"
	"strconv"
)

type Item struct {
	DataType string
	Data     interface{}
}

/**
返回字符串配置
*/
func (data *Item) ToString() string {
	if data != nil {
		switch data.DataType {
		case "string":
			return data.Data.(string)
		case "bool":
			return strconv.FormatBool(data.Data.(bool))
		case "int":
			return strconv.Itoa(data.Data.(int))
		case "float64":
			return decimal.NewFromFloat(data.Data.(float64)).String()
		}
	}
	return ""
}

func (data *Item) ToInt() int {
	if data != nil {
		switch data.DataType {
		case "string":
			result, err := strconv.Atoi(data.Data.(string))
			if err == nil {
				return result
			}
			return 0
		case "bool":
			if data.Data.(bool) {
				return 1
			}
			return 0
		case "int":
			return data.Data.(int)
		case "float64":
			return int(decimal.NewFromFloat(data.Data.(float64)).IntPart())
		}
	}
	return 0
}

/**
返回bool
*/
func (data *Item) ToBool() bool {
	if data != nil {
		switch data.DataType {
		case "string":
			result, err := strconv.ParseBool(data.Data.(string))
			if err == nil {
				return result
			}
			return false
		case "bool":
			return data.Data.(bool)
		case "int":
			return data.Data.(int) == 1
		case "float64":
			return false
		}
	}
	return false
}

/**
返回float64
*/
func (data *Item) ToFloat64() float64 {
	if data != nil {
		switch data.DataType {
		case "string":
			result, err := strconv.ParseFloat(data.Data.(string), 64)
			if err == nil {
				return result
			}
			return 0
		case "bool":
			return 0
		case "int":
			return float64(data.Data.(int))
		case "float64":
			return data.Data.(float64)
		}
	}
	return 0
}

/**
返回字符串切片
*/
func (data *Item) ToSString() []string {
	result := []string{}
	if data.Data != nil {
		for _, item := range data.Data.([]interface{}) {
			switch value := item.(type) {
			case string:
				result = append(result, value)
			case int:
				result = append(result, strconv.Itoa(value))
			case bool:
				result = append(result, strconv.FormatBool(value))
			case float64:
				result = append(result, decimal.NewFromFloat(value).String())
			}
		}
	}
	return result
}

/**
返回整形切片
*/
func (data *Item) ToSInt() []int {
	result := []int{}
	if data.Data != nil {
		for _, item := range data.Data.([]interface{}) {
			switch value := item.(type) {
			case string:
				num, err := strconv.Atoi(value)
				if err != nil {
					num = 0
				}
				result = append(result, num)
			case int:
				result = append(result, value)
			case bool:
				num := 0
				if value {
					num = 1
				}
				result = append(result, num)
			case float64:
				result = append(result, int(decimal.NewFromFloat(value).IntPart()))
			}
		}
	}
	return result
}

/**
返回float64切片
*/
func (data *Item) ToSFloat() []float64 {
	result := []float64{}
	if data.Data != nil {
		for _, item := range data.Data.([]interface{}) {
			switch value := item.(type) {
			case string:
				num, err := strconv.ParseFloat(value, 64)
				if err != nil {
					num = 0
				}
				result = append(result, num)
			case bool:
				result = append(result, 0)
			case int:
				result = append(result, float64(value))
			case float64:
				result = append(result, value)
			}
		}
	}
	return result
}

/**
返回bool切片
*/
func (data *Item) ToSBool() []bool {
	result := []bool{}
	if data.Data != nil {
		for _, item := range data.Data.([]interface{}) {
			switch value := item.(type) {
			case string:
				_bool, err := strconv.ParseBool(value)
				if err != nil {
					_bool = false
				}
				result = append(result, _bool)
			case bool:
				result = append(result, value)
			case int:
				result = append(result, value == 1)
			case float64:
				result = append(result, false)
			}
		}
	}
	return result
}
