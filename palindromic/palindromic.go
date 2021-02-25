// 判断字符串是否回文
package palindromic

import (
	"go-data/util"
	"strconv"
)

//ValidPalindromic 判断是否为回文，回文即字符串反转以后和原来的字符串一样
func ValidPalindromic(s string) bool {
	if util.Reverse(s) == s {
		return true
	}
	return false
}

func isPalindrome(x int) bool {
	s := strconv.Itoa(x)
	arr := []rune(s)
	l := len(arr)
	begin := 0
	end := l - 1
	for begin < end {
		arr[begin], arr[end] = arr[end], arr[begin]
		begin++
		end--
	}
	return s == string(arr)
}
