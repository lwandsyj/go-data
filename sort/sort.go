package sort

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
