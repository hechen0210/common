/**
* Created by GoLand
* User: hechen
* Date: 2020/10/6
* Time: 6:29 下午
 */
package helper

import "reflect"

func Struct2Map(obj interface{}) map[string]interface{} {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	var data = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		if v.Field(i).Interface() != 0 && v.Field(i).Interface() != "" {
			data[t.Field(i).Name] = v.Field(i).Interface()
		}
	}
	return data
}
