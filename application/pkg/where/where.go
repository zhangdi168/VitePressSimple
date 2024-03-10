// Package where
// @Author: zhangdi
// @File: common
// @Version: 1.0.0
// @Date: 2023/5/18 11:58
package where

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"wailstemplate/application/pkg/utils"
)

// Field 字段名-字段值结构体
type Field struct {
	Field string      `gorm:"not null" json:"field"`
	Value interface{} `gorm:"not null" json:"value"`
	//Opt 可选条件：= > < >=  <=  LIKE  BETWEEN IN
	Opt    string `gorm:"not null" json:"opt,omitempty"`
	ErrMsg string `json:"err_msg,omitempty"`
}

type Extra struct {
	OrderByColumn string `json:"order_by_column,omitempty"`
	OrderByDesc   bool   `json:"order_by_desc,omitempty"`
	PageSize      int    `json:"pageSize,omitempty"`
	PageNum       int    `json:"pageNum,omitempty"`
}

// StrToUint64 将字符串转uint64
func (s *Field) StrToUint64() *Field {
	NumberInt, err := strconv.ParseUint(s.Value.(string), 10, 64)
	if err != nil {
		s.ErrMsg = "类型转成uint64失败"
		return s
	}
	s.Value = NumberInt
	return s
}

// OptEqual 创建一个FieldWhere对象，操作符默认为 =
func OptEqual(fieldName string, fieldValue any) *Field {
	return &Field{Field: fieldName, Value: fieldValue, Opt: "=", ErrMsg: ""}
}

// Opt 创建一个带操作符opt的FieldWhere对象
func Opt(fieldName string, opt string, fieldValue interface{}) *Field {
	optList := []string{"=", ">", "<", ">=", "<=", OptLike, OptIn, OptBetween}
	if !utils.ArrContains(optList, opt) {
		return &Field{ErrMsg: "操作符不符合，可选操作符: = > < >=  <=  LIKE  BETWEEN IN"}
	}

	if opt == "IN" {
		if reflect.TypeOf(fieldValue).Kind() != reflect.Slice && reflect.TypeOf(fieldValue).Kind() != reflect.Array {
			return &Field{ErrMsg: "IN操作需要传入一个列表（数组或者切片）"}
		}
	} else if opt == "BETWEEN" {
		if reflect.TypeOf(fieldValue).Kind() != reflect.Slice && reflect.TypeOf(fieldValue).Kind() != reflect.Array {
			return &Field{ErrMsg: "BETWEEN操作需要传入一个长度为2的列表"}
		} else {
			sliceValue := reflect.ValueOf(fieldValue)
			if sliceValue.Len() != 2 {
				return &Field{ErrMsg: "BETWEEN操作需要传入一个长度为2的列表"}
			}
		}
	}
	return &Field{Field: fieldName, Value: fieldValue, Opt: opt, ErrMsg: ""}
}

// ConvertGormRepeatStr 找出所有的opt = REPEAT 的条件，并组装出Having和Group对应字符串,返回空则表示无重复选项
// demo:查询同时满足两个字段重号的记录:调用示例 Opt("num",OptRepeat,1) 表示num字段重复数量大于1的所有数据
// db.Group("email, phone_number")
// .Having("count(email) > 1 AND count(phone_number) > 1")
// .Find(&users)
//func ConvertGormRepeatStr(fieldWheres []*Field) (groupStr string, Having string) {
//	for _, fw := range fieldWheres {
//		switch fw.Opt {
//		case OptRepeat:
//			groupStr += fmt.Sprintf("%s,", fw.Field)
//			Having += fmt.Sprintf("count(%s) > %v AND ", fw.Field, fw.Value)
//		}
//	}
//	return strings.TrimRight(groupStr, ","), strings.TrimRight(Having, "AND ")
//}

// ConvertToGormWhere 将切片FieldWhere转换为符合gorm库的Where条件语句
func ConvertToGormWhere(fieldWheres []*Field) (string, []interface{}, error) {
	var whereStr string
	var whereValues []interface{}

	for _, fw := range fieldWheres {
		if fw.ErrMsg != "" {
			return "", nil, errors.New("异常错误：" + fw.Field + fw.Opt + fw.ErrMsg)
		}
		switch fw.Opt {
		case "=":
			whereStr += fmt.Sprintf("%s = ? AND ", fw.Field)
			whereValues = append(whereValues, fw.Value)
		case ">":
			whereStr += fmt.Sprintf("%s > ? AND ", fw.Field)
			whereValues = append(whereValues, fw.Value)
		case "<":
			whereStr += fmt.Sprintf("%s < ? AND ", fw.Field)
			whereValues = append(whereValues, fw.Value)
		case ">=":
			whereStr += fmt.Sprintf("%s >= ? AND ", fw.Field)
			whereValues = append(whereValues, fw.Value)
		case "<=":
			whereStr += fmt.Sprintf("%s <= ? AND ", fw.Field)
			whereValues = append(whereValues, fw.Value)
		case "LIKE":
			whereStr += fmt.Sprintf("%s LIKE ? AND ", fw.Field)
			whereValues = append(whereValues, fw.Value)
		case "IN":

			//inPlaceholder := strings.TrimRight(strings.Repeat("?,", len(inValues)), ",")
			whereStr += fmt.Sprintf("%s IN ? AND ", fw.Field)
			whereValues = append(whereValues, fw.Value)
		case "BETWEEN":
			//使用反射获取 fw.Value 的值和类型信息：
			val := reflect.ValueOf(fw.Value)
			//创建新的 []interface{} 切片：
			inValues := make([]interface{}, val.Len())
			//使用反射遍历原始数组，并将每个元素转换为 interface{} 类型：
			for i := 0; i < val.Len(); i++ {
				inValues[i] = val.Index(i).Convert(reflect.TypeOf((*interface{})(nil)).Elem()).Interface()
			}
			if len(inValues) != 2 {
				return "", nil, errors.New("BETWEEN操作需要传入一个长度为2的列表")
			}

			whereStr += fmt.Sprintf("%s BETWEEN ? AND ? AND ", fw.Field)
			whereValues = append(whereValues, inValues[0], inValues[1])
		}

	}
	whereStr = strings.TrimRight(whereStr, "AND ")
	return whereStr, whereValues, nil
}

// Format 组装多个条件，where使用Opt函数组装一个条件
func Format(wheres ...*Field) []*Field {
	res := append(
		make([]*Field, 0),
		wheres...,
	)
	return res
}

// FormatOne 组装1个条件
func FormatOne(FieldName, opt string, val any) []*Field {
	field := &Field{
		Field: FieldName,
		Value: val,
		Opt:   opt,
	}
	res := append(
		make([]*Field, 0),
		field,
	)
	return res
}
