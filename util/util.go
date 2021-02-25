package util

import "fmt"

//Reverse 反转
func Reverse(s string) string {
	if s == "" {
		return s
	}
	result := []rune(s)
	l := len(result)
	begin := 0
	end := l - 1
	for begin < end {
		result[begin], result[end] = result[end], result[begin]
		begin++
		end--
	}
	return string(result)
}

//TestReverse ...
func TestReverse() {
	s := "abcde"
	d := Reverse(s)
	fmt.Println(s)
	fmt.Println(d)
}
