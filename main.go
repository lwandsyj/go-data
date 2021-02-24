package main

import (
	"fmt"
	"go-data/stack"
	"reflect"
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

func main() {
	var a *[]interface{} = &[]interface{}{}
	fmt.Println(a)
	*a = append(*a, 2)
	fmt.Println(a)

	b := make([]interface{}, 0)
	c := &b
	fmt.Println(reflect.TypeOf(c))
	*c = append(*c, 4)
	fmt.Println(c)
}
