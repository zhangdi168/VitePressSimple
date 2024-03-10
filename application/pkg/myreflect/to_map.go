// Package myreflect
// @Author: zhangdi
// @File: to_map
// @Version: 1.0.0
// @Date: 2023/11/14 18:45
package myreflect

import (
	"fmt"
	"reflect"
	"strings"
)

// StructToMap 将结构体(传入[非指针]类型)转成map key取传入的tag名，如json
func StructToMap(s any, tagName string) map[string]string {
	result := make(map[string]string)

	// 通过反射获取结构体的类型和值
	t := reflect.TypeOf(s)
	v := reflect.ValueOf(s)

	// 遍历结构体字段
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i)

		// 获取字段的json标签和值
		jsonTag := field.Tag.Get(tagName)
		tagValArr := strings.Split(jsonTag, ",")
		tagVal := jsonTag
		if len(tagValArr) > 0 {
			tagVal = tagValArr[0]
		}
		if tagVal == "" {
			tagVal = field.Name
		}
		fieldValue := fmt.Sprintf("%v", value.Interface())

		result[tagVal] = fieldValue
	}

	return result
}
