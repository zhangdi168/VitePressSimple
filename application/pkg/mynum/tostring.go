// Package mynum
// @Author: zhangdi
// @File: conv
// @Version: 1.0.0
// @Date: 2023/8/28 16:29
package mynum

import "strconv"

type Numbers interface {
	uint | int | uint16 | uint8 | uint64 | uint32 | int8 | int16 | int32 | int64
}

// ToString  通用的泛型方法
//mynum.ToString[uint](16)
func ToString[T Numbers](val T) string {
	return strconv.Itoa(int(val))
}

func UintToString(str uint) string {
	return strconv.Itoa(int(str))
}

func Uint16ToString(str uint16) string {
	return strconv.Itoa(int(str))
}

func Uint64ToString(str uint64) string {
	return strconv.Itoa(int(str))
}

func Uint32ToString(str uint32) string {
	return strconv.Itoa(int(str))
}

func IntToString(str int) string {
	return strconv.Itoa(int(str))
}

func Int16ToString(str int16) string {
	return strconv.Itoa(int(str))
}

func Int32ToString(str int32) string {
	return strconv.Itoa(int(str))
}

func Int64ToString(str int64) string {
	return strconv.Itoa(int(str))
}
