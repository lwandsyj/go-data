// 判断字符串是否回文
package palindromic

import "go-data/util"

//ValidPalindromic 判断是否为回文，回文即字符串反转以后和原来的字符串一样
func ValidPalindromic(s string) bool {
	if util.Reverse(s) == s {
		return true
	}
	return false
}
