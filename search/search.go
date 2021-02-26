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
