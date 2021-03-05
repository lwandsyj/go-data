package main

import (
	"fmt"
)

func f1(x int) {
}

func f3() int {
	return 0
}
func f4(x int) int {
	return x
}
func f5() (x int, y string) {
	x = 0
	y = "hello"
	return
}
func f() {
	fmt.Println("f test")
}
func f2(x int, y string) {
	fmt.Println("x=", x, "y=", y)
}

// 函数作为参数
// fn 参数是有两个参数，没有返回值的函数类型
func test(fn func() (int, string)) {
	x, y := fn() // 有多个返回值的
	fmt.Println("x=", x, "y=", y)
}

//函数类型声明可以带有参数名称，也可以不带参数名称
func test1(fn func() (x int, y string)) {
	m, n := fn() // 有返回值的
	fmt.Println("m=", m, "n=", n)
}
func test2(x, y int) func(int) int {
	// 返回函数，类型声明和返回类型保持一致
	return func(z int) int {
		// 函数的返回值
		return x + y + z
	}
}

func index(s, p string) int {
	ls := len([]rune(s))
	lp := len([]rune(p))
	for i := 0; i+lp <= ls; i++ {
		tmp := string(s[i : i+lp])
		fmt.Println(tmp)
		if tmp == p {
			return i
		}
	}
	return -1
}

type people struct {
	name string
	age  int
}

// 使用people 是值类型传递，传递的是副本
// 变量修改的是副本的值，与原始数据无关
func (p people) setName(name string) {
	p.name = name
}

// 使用people 指针类型传递，传递的内存地址引用
// 变量修改的值会影响原始数据
func (p *people) setName1(name string) {
	p.name = name
}

func main() {
	p := people{
		name: "张三",
	}

	p.setName("李四")
	fmt.Println(p)

	c := &p
	c.setName1("李四")
	fmt.Println(c)
}
