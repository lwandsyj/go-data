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

// 反转字符串是否和原来的字符一样
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

// 二分查找算法，比较第一个和最后一个是否相等，依次类推
func isPalindromeBy(x int) bool {
	s := strconv.Itoa(x)
	arr := []rune(s)
	l := len(arr)
	for i := 0; i <= l/2; i++ {
		// 逐个判断，直到
		if arr[i] != arr[l-1-i] {
			return false
		}
	}
	return true
}
