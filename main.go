package main

import (
	"fmt"
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

func transpose(matrix [][]int) [][]int {
	l := len(matrix)
	var max int
	for i := 0; i < l; i++ {
		count := len(matrix[i])
		if max < count {
			max = count
		}
	}
	return nil
}

func main() {
	var result [][]int
	fmt.Println(result)
}
