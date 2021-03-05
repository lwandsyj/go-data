package stringsearch

import "fmt"

/*
 窗口滑动算法
*/

// 求取数组中最大值
func maxSlice(arr []int) int {
	max := arr[0] // 默认最大值为第一个
	l := len(arr) // 数组长度
	for i := 1; i < l; i++ {
		// 和数组中元素依次比较，
		// 如果max 小于当前元素，则默认设置当前元素为最大值
		if max < arr[i] {
			max = arr[i]
		}
	}
	return max
}

//1. 固定长度
//maxSlidingWindow ...
func MaxSlidingWindow(nums []int, k int) []int {
	fmt.Println("k=", k)
	rtn := []int{} // 返回值
	l := len(nums) // 元素个数
	if l == 0 {
		return rtn
	}
	index := 0
	for i := k; i <= l; i++ {
		//窗口逐步向右移动,即窗口中左边删除一个，右边添加一个
		window := nums[index:i]
		fmt.Println("window1=", window)
		max := maxSlice(window)
		rtn = append(rtn, max)
		index++
	}
	return rtn
}
