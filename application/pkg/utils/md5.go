// Package utils
// @Author: zhangdi
// @File: md5
// @Version: 1.0.0
// @Date: 2023/5/6 10:54
package utils

import (
	"crypto/md5"
	"fmt"
)

func MD5(str string) string {
	// 创建一个新的 MD5 哈希器对象。
	hasher := md5.New()

	// 将字符串转换成字节数组，并写入哈希器中。
	hasher.Write([]byte(str))

	// 计算哈希值，并使用字符串格式化输出生成16进制字符串。
	return fmt.Sprintf("%x", hasher.Sum(nil))
	//需要注意的是，由于我们只对传入字符串进行了哈希处理，而没有进行加盐或其他处理
	//（例如，使用多次迭代或慢哈希函数），因此可能会存在密码被破解等安全问题。
	//建议在实际开发中结合具体业务场景选择合适的加密方案来保障系统的安全性。
}
