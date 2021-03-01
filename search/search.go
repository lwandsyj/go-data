//Package search 搜索
package search

//SimpleSearch 简单搜索--无序搜索
func SimpleSearch(arr []interface{}, data interface{}) (index int) {
	if len(arr) == 0 {
		index = -1
		return
	}
	for key, value := range arr {
		if value == data {
			index = key
			break
		}
	}
	return
}

//BinarySearch 二分查找--序列必须是有序的
func BinarySearch(sortArr []int, data int) int {
	left := 0
	count := len(sortArr)
	if count == 0 {
		return -1
	}
	right := count - 1
	//循环条件必须保证左边始终小于右边，不然终止循环
	for left < right {
		// 查找中间值，go 整数相除放回整数，比如5/2=2
		mid := (left + right) / 2
		//查看中间元素是否是要查找的元素
		if sortArr[mid] == data {
			return mid
		} else if sortArr[mid] > data {
			// 这里必须要求序列是有序，不然这里无法成立
			// 中间数大于要查找的数据，那么表示如果存在则必须在前半部分
			right = mid - 1
		} else {
			// 否则则后半部分
			left = mid + 1
		}
	}
	// 查找不到返回-1
	return -1
}
