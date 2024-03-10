// Package myreflect
// @Author: zhangdi
// @File: get_struct_json
// @Version: 1.0.0
// @Date: 2023/10/20 14:44
package myreflect

import (
	"reflect"
	"strings"
)

// GetJSONTags 获取该结构体所有字段的json标签值放到字符串数组
// 传入指针地址 如 &Persion{}
func GetJSONTags(s interface{}) []string {
	var tags []string
	//v := reflect.ValueOf(s)
	t := reflect.TypeOf(s).Elem()
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		jsonTag := field.Tag.Get("json")
		val := ""
		if jsonTag == "" {
			//不存在则取字段名
			val = field.Name
		} else {
			jsonTagArr := strings.Split(jsonTag, ",")
			val = jsonTagArr[0]
		}

		tags = append(tags, val)
	}
	return tags
}
