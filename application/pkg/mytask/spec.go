// Package mytask
// @Author: zhangdi
// @File: spcc_cnt
// @Version: 1.0.0
// @Date: 2023/11/16 13:56
package mytask

import "fmt"

// 参考链接：https://juejin.cn/post/7027090787305259022?from=search-suggest
// */5 * * * * ?   后往前读：周 月日时分秒
// 秒(0-59)  分(0-59)  小时(0-23)  日期(1-31)  月份(1-12)  星期几(0-6)
// ?只用于 日和 星期两个参数，表示不指定值，可以用于代替 *
//* 表示所有值，比如月份参数是 * 则表示每月

// FormatEverySpace 组装spec:每隔一段时间执行
func FormatEverySpace(hour int, minutes int, seconds int) string {
	spec := "@every "
	if hour >= 0 && hour <= 24 {
		//小时
		spec += fmt.Sprintf("%dh", hour)
	}
	if minutes >= 0 && minutes <= 60 {
		//分
		spec += fmt.Sprintf("%dm", minutes)
	}
	if seconds >= 0 && seconds <= 60 {
		//秒
		spec += fmt.Sprintf("%ds", seconds)
	}
	if spec == "@every " {
		//没有组装任何东西
		return ""
	}
	return spec
}

// FormatEveryDay 每天的某点执行
func FormatEveryDay(hour int, minutes int, seconds int) string {
	spec := fmt.Sprintf("%d %d %d * * ?", seconds, minutes, hour)
	return spec
}

//每隔5秒执行一次：*/5 * * * * ?
//
//每隔1分钟执行一次：0 */1 * * * ?
//
//每天23点执行一次：0 0 23 * * ?
//
//每天凌晨1点执行一次：0 0 1 * * ?
//
//每月1号凌晨1点执行一次：0 0 1 1 * ?
//
//每周一和周三晚上22:30: 00 30 22 * * 1,3
//
//在26分、29分、33分执行一次：0 26,29,33 * * * ?
//
//每天的0点、13点、18点、21点都执行一次：0 0 0,13,18,21 * * ?
//
//每年三月的星期四的下午14:10和14:40:  00 10,40 14 ? 3 4
//(*)表示 cron 表达式能匹配该字段的所有值。如在第5个字段使用星号(month)，表示每个月(/)表示增长间隔，如第1个字段(minutes) 值是 3-59/15，表示每小时的第3分钟开始执行一次，之后每隔 15 分钟执行一次（即 3、18、33、48 这些时间点执行），这里也可以表示为：3/15(,)用于枚举值，如第6个字段值是 MON,WED,FRI，表示 星期一、三、五 执行(-)表示一个范围，如第3个字段的值为 9-17 表示 9am 到 5pm 直接每个小时（包括9和17）(?)只用于 日(Day of month) 和 星期(Day of week)，表示不指定值，可以用于代替 *

// SpecFormatEverySpace 组装每隔一段时间执行的spec字符串
