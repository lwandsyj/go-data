package num

//SumNum 计算给定数字各个位数的总和
func SumNum(num int) int {
	var sum int
	// 如果num >0 则一次取个位值
	for num > 0 {
		// 取个位值，并依次相加
		sum += num % 10
		// 获取新的数值
		num = num / 10
	}
	return sum
}

//ReverseNum 反转数字
func ReverseNum(num int) int {
	var sum int
	// 遍历各个位数的值
	for num > 0 {
		// 小位数转大数则一步一步*10
		sum = sum*10 + num%10
		num = num / 10
	}
	return sum
}
