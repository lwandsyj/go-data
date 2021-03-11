package num

import (
	"fmt"
	"path"
	"runtime"
)

//SumNum 计算给定数字各个位数的总和
func SumNum(num int) int {
	//
	//dir, _ := os.Getwd()
	_, file, line, ok := runtime.Caller(0)
	base := path.Base(file)
	fmt.Println(path.Dir(file))
	fmt.Println(base)
	dir, filename := path.Split(file)
	fmt.Println(dir, filename)
	ext := path.Ext(file)
	fmt.Println(ext)
	fmt.Println(file, line, ok)
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
