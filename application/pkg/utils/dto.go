// Package utils
// @Author: zhangdi
// @File: dto
// @Version: 1.0.0
// @Date: 2023/5/18 10:15
package utils

import "reflect"

// DtoToEntity 将 DTO 转换成实体对象
func DtoToEntity(dto interface{}, entity interface{}) {
	// 获取Dto对象和Entity对象的反射值
	dtoValue := reflect.ValueOf(dto).Elem()       // dto必须传入一个指针类型
	entityValue := reflect.ValueOf(entity).Elem() // entity必须传入一个指针类型

	// 遍历Dto对象的每个字段并转换
	for i := 0; i < dtoValue.NumField(); i++ { // NumField返回结构体中字段的数量
		// 获取Dto对象的字段值
		dtoField := dtoValue.Field(i)

		// 如果Dto对象的字段值不为空，则进行转换
		if !dtoField.IsZero() {
			// 获取Dto对象的字段名
			fieldName := dtoValue.Type().Field(i).Name

			// 根据字段名获取Entity对象的同名字段
			entityField := entityValue.FieldByName(fieldName)

			// 若找到同名字段且类型相同，将Dto对象的值复制给Entity对象
			if entityField.IsValid() && entityField.Type() == dtoField.Type() {
				entityField.Set(dtoField)
			}
		}
	}
}
