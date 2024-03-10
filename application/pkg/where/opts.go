// Package where
// @Author: zhangdi
// @File: opts
// @Version: 1.0.0
// @Date: 2023/8/3 14:43
package where

// OptLike  like查询，value传值示例 %12%
const OptLike = "LIKE"

// OptIn In查询，value 传值示例[]string{"Alice", "Bob", "Charlie"
const OptIn = "IN"

// OptBetween  传值示例[]int{1, 3}，表示查找范围从1到3
const OptBetween = "BETWEEN"

// OptRepeat 调用示例 Opt("num",OptRepeat,1) 表示num字段重复数量大于1的所有数据

const Greater = ">"
const greaterEqual = ">="
const Less = "<"
const LessEqual = "<="
const Equal = "="

//大于 小于 大于或等于 小于或等于
