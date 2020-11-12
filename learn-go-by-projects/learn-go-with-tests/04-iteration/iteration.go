package iteration

import "strings"

const repeatCnt = 5

/**
测试文档，这个注释不属于函数，不会在doc文档中出现
*/

// 这是Repeat函数的注释
func Repeat(s string) string {
	str := ""
	for i := 0; i < repeatCnt; i++ {
		str += s
	}
	return str
}

/**
这是RepeatWithStringsFunction函数的注释
*/
func RepeatWithStringsFunction(s string) string {
	return strings.Repeat(s, repeatCnt)
}
