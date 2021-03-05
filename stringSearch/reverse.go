package stringsearch

//Reverse 反转字符串
func Reverse(s string) string {
	// 字符串不可以修改，要想通过索引修改字符，需要转成切片
	arr := []rune(s)
	l := len(arr) // 获取字符串长度
	if l == 0 {
		return s
	}
	// 因为要前后对称交换，这里使用一半长度即可
	for i := 0; i < l/2; i++ {
		// 使用go 多重变量定义方式快速实现字符的反转
		arr[i], arr[l-i-1] = arr[l-i-1], arr[i]
	}
	// 把切片转换成字符串
	return string(arr)
}

//Reverse1 双指针反转字符串
func Reverse1(s string) string {
	arr := []rune(s)
	l := len(arr)
	if l == 0 {
		return s
	}
	// 双指针左右两个变量
	for i, j := 0, l-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	return string(arr)
}
