package main

import (
	"fmt"
	"go-data/stack"
	"strconv"
	"strings"
)

func learnStatck() {
	s := stack.New()
	s.Push(2)
	s.Push(3)
	s.Push(4)
	fmt.Println(s)
	top := s.Pop()
	fmt.Println(top)

	a := make([]int, 0)
	b := make([]int, 2, 3)
	fmt.Println(a) //[]
	fmt.Println(b) //[0 0]

	d := []int{1, 2, 3, 4, 5, 6}
	c := d[:2]
	c = append(c, d[4:]...)
	fmt.Println("d=", d)
	fmt.Println("c=", c)
}

func isPalindrome(x int) bool {
	s := strconv.Itoa(x)
	arr := []rune(s)
	l := len(arr)
	for i := 0; i < l/2; i++ {
		if arr[i] != arr[l-1-i] {
			return false
		}
	}
	return true
}

func sort(arr []string) []string {
	l := len(arr)
	for i := 0; i < l; i++ {
		for j := 0; j < l-i-1; j++ {
			if arr[j]+arr[j+1] > arr[j+1]+arr[j] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

func convert(arr []int) []string {
	rtn := make([]string, 0)
	for _, value := range arr {
		rtn = append(rtn, strconv.Itoa(value))
	}
	return rtn
}
func main() {
	a := []int{3, 30, 34, 5, 9}
	b := convert(a)
	fmt.Println(len(b))
	b = sort(b)
	fmt.Println(strings.Join(b, ""))
}
