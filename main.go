package main

import (
	"fmt"
	"go-data/sort"
	"go-data/stack"
	"strconv"
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

func main() {
	arr := []int{3, 4, 5, 2, 1}
	arr = sort.SelectSort(arr)
	fmt.Println(arr)
}
