package where

import "reflect"

// ArrContains  它可以接收任何类型的切片和元素，并返回元素是否属于给定切片
func arrContains(slice interface{}, item interface{}) bool {
	// 获取切片变量的类型
	switch reflect.TypeOf(slice).Kind() {
	// 判断切片类型是否是 Slice 或 Array 类型
	case reflect.Slice, reflect.Array:
		// 将 slice 转成反射对象
		s := reflect.ValueOf(slice)

		// 遍历反射对象中的元素
		for i := 0; i < s.Len(); i++ {
			// 比较当前元素和目标元素是否相等
			if reflect.DeepEqual(item, s.Index(i).Interface()) {
				return true
			}
		}
	}

	return false
}
