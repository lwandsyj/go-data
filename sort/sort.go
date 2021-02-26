package sort

import "fmt"

//BubbleSort 冒泡排序
func BubbleSort(arr []int) []int {
	l := len(arr)
	for i := 0; i < l-1; i++ {
		for j := 0; j < l-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

//SelectSort 选择排序
func SelectSort(arr []int) []int {
	l := len(arr)
	// 当到最后一个时，差不多已排好序，不需要再比因此为l-1
	for i := 0; i < l-1; i++ {
		// 默认为第一个元素为最小值
		minIndex := i
		// 遍历循环下面的所有元素
		for j := i + 1; j < l; j++ {
			// 找到所有元素中最小的一个值
			if arr[minIndex] > arr[j] {
				// 记录最小值的索引
				minIndex = j
			}
		}
		if i != minIndex {
			// 交换两个元素的值
			arr[i], arr[minIndex] = arr[minIndex], arr[i]
		}

	}
	return arr
}

//InsertSort 插入排序
func InsertSort(arr []int) []int {
	l := len(arr)
	// 开始默认把第一个元素当做已排序好的元素,从第一个开始和已排序好的元素依次比较
	for i := 1; i < l; i++ {
		// 提取元素
		j := i
		// 和排序号的元素依次比较
		for j > 0 {
			//如果元素小于当前元素，则向前移动
			if arr[j] < arr[j-1] {
				// 交换两个元素的值
				arr[j], arr[j-1] = arr[j-1], arr[j]
				// 判断一次往前走一个位置，依次判断
				j--
			} else {
				// 如果大于当前元素则退出当前循环，继续外层循环
				break
			}

		}
	}
	return arr
}

//QuickSort 快速排序
func QuickSort(arr []int) {
	recursionSort(arr, 0, len(arr)-1)
}

// l:left,r:right
func recursionSort(arr []int, l, r int) {
	if l < r {
		index := getIndex(arr, l, r)
		fmt.Println(arr)
		// 左边部分
		recursionSort(arr, l, index)
		// 右边部分
		recursionSort(arr, index+1, r)
	}
}

//getIndex 获取基准值位置
func getIndex(arr []int, l, r int) int {
	base := arr[l] // 基准值
	for l < r {
		// 先从右边开始找到小于基准值的数据
		for l < r && arr[r] > base {
			r--
		}
		// 小于基准值的移到左边
		if l < r {
			arr[l] = arr[r]
			l++
		}
		// 然后从左边开始找到大于基准值的数据
		for l < r && arr[l] < base {
			l++
		}
		// 大于基准值的移到右边
		if l < r {
			arr[r] = arr[l]
			r--
		}
		arr[l] = base
	}
	return l
}
