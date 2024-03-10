// Package repository
// @Author: zhangdi
// @File: update_config
// @Version: 1.0.0
// @Date: 2023/11/26 20:08
package repository

import (
	"reflect"
	"strings"
	"unicode"
	"wailstemplate/application/pkg/utils"
)

// UpdateColumnsAll 强制更新所有字段，不管是否为零值
const UpdateColumnsAll = "#--all--#"

// UpdateColumnsNotZero 更更新非零值的字段
const UpdateColumnsNotZero = "#--not_zero--#"

// UpdateColumnsNotZeroAndStr 更新非零值和所有的【字符串】类型字段
const UpdateColumnsNotZeroAndStr = "#--not_zero_and_str--#"

// UpdateColumnsNotZeroAndNumber 更新非零值和所有的【整型】类型字段
const UpdateColumnsNotZeroAndNumber = "#--not_zero_and_number--#"

// UpdateTypList 允许的type列表
var UpdateTypList = []string{UpdateColumnsAll, UpdateColumnsNotZero,
	UpdateColumnsNotZeroAndStr, UpdateColumnsNotZeroAndNumber}

// UpdateConfig 更新配置
// UpdateFields优先级大于UpdateColumnsType
type UpdateConfig struct {
	//更新的字段类型 可选值有上面四个常量
	UpdateColumnsType string
	//指定更新的字段，当传入这个切片时 只会更新该列表里的字段值，不传入则值默认取UpdateColumnsNotZero
	UpdateFields []string
}

func GetUpdateColumns(ptr interface{}, updateColumnType string) []string {

	//传入的类型不符合类型列表则返回一个空的切片
	if !utils.ArrContains(UpdateTypList, updateColumnType) {
		return make([]string, 0)
	}

	// 定义一个切片用于存储列名
	var columns []string
	// 获取指针的值并获取其类型
	rv := reflect.ValueOf(ptr).Elem()
	rt := rv.Type()
	// 遍历结构体的所有字段
	for i := 0; i < rt.NumField(); i++ {
		// 获取字段信息
		field := rt.Field(i)
		// 获取字段的 gorm 标签和 json 标签
		gormTag := field.Tag.Get("gorm")
		// 获取字段的值
		value := rv.Field(i).Interface()
		isZero := reflect.DeepEqual(value, reflect.Zero(field.Type).Interface())

		//字段名
		columnName := getColumnName(gormTag)
		if columnName == "" { //没有定义列名，则取字段名的下划线命名法
			columnName = snakeCase(field.Name)
		}
		switch updateColumnType {
		case UpdateColumnsAll: //所有的字段
			columns = append(columns, columnName)
		case UpdateColumnsNotZero: //非零值字段
			if !isZero {
				// 非零值 将字段名插入到切片中
				columns = append(columns, columnName)
			}
		case UpdateColumnsNotZeroAndStr: //非零值字段+所有的字符串类型字典
			if !isZero || field.Type.Kind() == reflect.String {
				// 非零值 将字段名插入到切片中
				columns = append(columns, columnName)
			}
		case UpdateColumnsNotZeroAndNumber: //非零值字段+所有的数字类型字段
			if !isZero || typeIsNumber(field.Type) {
				// 非零值 将字段名插入到切片中
				columns = append(columns, columnName)
			}
		}

	}
	// 返回需要更新的字段列表
	return columns
}

// 这个函数用于从 gorm 标签中提取列名，tag 为 gorm 标签的字符串。函数的逻辑如下：
//
// 首先通过 strings.Index 查找 tag 中 "column:" 的索引位置，如果找到了则返回位置，否则返回 -1。
// 如果找到了 "column:" 的索引位置，则使用 tag[idx+7:] 从 "column:" 后面截取字符串，得到 tag 的剩余内容。
// 然后再次使用 strings.Index 查找 tag 中 ";" 的索引位置，如果找到了则返回位置，否则返回 -1。
// 如果找到了 ";" 的索引位置，则使用 tag[:idx] 从开头截取到 ";" 的内容，即为最终的列名。
// 如果没有找到 ";" 的索引位置，说明没有多余的内容需要截取，直接返回 tag 即可。
// 如果没有找到 "column:" 的索引位置，说明 gorm 标签中没有 "column:" 属性，返回空字符串。
func getColumnName(tag string) string {
	// 查找 "column:" 的索引位置
	if idx := strings.Index(tag, "column:"); idx != -1 {
		// 截取 "column:" 后面的内容
		tag = tag[idx+7:]
		// 查找 ";" 的索引位置，用于截取多余的内容
		if idx := strings.Index(tag, ";"); idx != -1 {
			tag = tag[:idx]
		}
		return tag
	}
	return ""
}

// NameToLine 下划线命名
func snakeCase(camel string) string {
	var snake []rune
	for i, c := range camel {
		if unicode.IsUpper(c) {
			if i > 0 && (i+1 < len(camel) && unicode.IsLower(rune(camel[i+1])) || unicode.IsLower(rune(camel[i-1]))) {
				snake = append(snake, '_')
			}
			snake = append(snake, unicode.ToLower(c))
		} else {
			snake = append(snake, c)
		}
	}
	return string(snake)
}

// toTypeScriptType 判断类型是否为数字要bool类型也属于整型范畴
func typeIsNumber(golangType reflect.Type) bool {
	switch golangType.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Float32, reflect.Float64, reflect.Bool:
		return true
	default:
		return false
	}

	return false
}
