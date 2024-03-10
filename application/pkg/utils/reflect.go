// Package utils
// @Author: zhangdi
// @File: reflect
// @Version: 1.0.0
// @Date: 2023/5/18 10:13
package utils

import (
	"errors"
	"reflect"
	"strings"
)

// IsZeroValue 判断一个变量是否是零值
func IsZeroValue(v interface{}) bool {
	value := reflect.ValueOf(v)
	switch value.Kind() {
	case reflect.Slice, reflect.Map, reflect.Chan, reflect.Func, reflect.Ptr:
		// 复合类型、函数类型和指针类型的零值是 nil
		return value.IsNil()
	case reflect.Struct:
		// 结构体类型需要判断所有字段是否为零值
		for i := 0; i < value.NumField(); i++ {
			if !IsZeroValue(value.Field(i).Interface()) {
				return false
			}
		}
		return true
	default:
		// 其他类型直接与类型的零值进行比较
		return reflect.DeepEqual(v, reflect.Zero(reflect.TypeOf(v)).Interface())
	}
}

// StructToMap 通过反射的方式获取一个结构体的所有字段，并将其映射到一个map
//注意的到的map会区分大小写，也就是字段首字母大写，得到的key首字母也是大写
func StructToMap(data interface{}) (map[string]interface{}, error) {
	v := reflect.ValueOf(data)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() != reflect.Struct {
		return nil, errors.New("not a struct")
	}
	res := make(map[string]interface{})
	typ := v.Type()
	for i := 0; i < v.NumField(); i++ {
		f := v.Field(i)
		name := typ.Field(i).Name
		if f.Kind() == reflect.Ptr {
			if f.IsNil() {
				res[name] = nil
				continue
			}
			f = f.Elem()
		}
		res[name] = f.Interface()
	}
	return res, nil
}

// 将map 的key转小写
func LowercaseKeysMap(originalMap map[string]interface{}) map[string]interface{} {
	lowerCaseMap := make(map[string]interface{}, len(originalMap))

	for k, v := range originalMap {
		newKey := strings.ToLower(k)
		lowerCaseMap[newKey] = v
	}

	return lowerCaseMap
}

// StructToMapFilterNil  通过反射的方式获取一个结构体的所有字段，并将其映射到一个map
// 如果字段是零值或者空值，则map中不加入该字段
func StructToMapFilterNil(data interface{}) (map[string]interface{}, error) {
	v := reflect.ValueOf(data)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() != reflect.Struct {
		return nil, errors.New("not a struct")
	}
	res := make(map[string]interface{})
	typ := v.Type()
	for i := 0; i < v.NumField(); i++ {
		f := v.Field(i)
		name := typ.Field(i).Name
		if f.Kind() == reflect.Ptr {
			if f.IsNil() {
				continue
			}
			f = f.Elem()
		}
		zeroValue := reflect.Zero(f.Type())
		if reflect.DeepEqual(f.Interface(), zeroValue.Interface()) {
			continue
		}
		switch f.Kind() {
		case reflect.String:
			if f.String() == "" {
				continue
			}
		case reflect.Slice, reflect.Map:
			if f.Len() == 0 {
				continue
			}
		}
		res[name] = f.Interface()
	}
	return res, nil
}
