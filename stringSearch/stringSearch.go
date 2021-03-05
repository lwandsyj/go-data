package stringsearch

import "fmt"

//BruteSearch 暴力破解版查找子字符串,返回索引，查找不到返回-1
func BruteSearch(s, p string) int {
	// 使用rune 是因为防止中文等其他类型导致的错误，rune 返回utf-8 字符
	lp := len([]rune(p)) // 返回p 的长度
	// 如果子串为空
	if lp == 0 {
		return 0
	}
	ls := len([]rune(s))
	for i := 0; i+lp <= ls; i++ {
		// 如果字符串中字符和p中第一个字符相等，则s[i]后面的字符逐个和p中字符一一相比较
		fmt.Println("if i=", i, "s[i]=", s[i], "p0=", p[0])
		if s[i] == p[0] {
			fmt.Println("i=", i, "s[i]=", s[i], "p0=", p[0])
			j := 1
			for ; j < lp; j++ {
				// 如果字符串s[i]后面的有和p中字符相等的情况，则表示不符合条件，终止内部循环
				if s[i+j] != p[j] {
					break // 终止内部循环，继续外部循环
				}
				// 如果相等则继续循环，一致到循环完毕，
			}
			//j==lp,循环完毕，表示s[i]后面的和p 所有字符都相等，表示找到字符串
			if j == lp {
				return i
			}
		}
	}
	return -1
}

//BruteSliceSearch 使用go 语言特有的切片功能优化
// 其他语言可以用截取字符串的形式获得
func BruteSliceSearch(s, p string) int {
	lp := len([]rune(p)) // 返回p 的长度
	// 如果子串为空
	if lp == 0 {
		return 0
	}
	ls := len([]rune(s)) // 字符串的长度

	// 这里用等号是因为切片不包好最后一位
	for i := 0; i+lp <= ls; i++ {
		tmp := s[i : i+lp] // 下标从0开始，切片,
		// 转成字符串判断是否和p 相等，相等返回i
		if string(tmp) == p {
			return i
		}
	}
	// 循环完毕，没有查找到表示不存在，返回-1
	return -1
}
