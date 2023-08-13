package test

import (
	"fmt"
	"testing"
)

func TestSelectSort(t *testing.T) {
	// 构造数据
	data := []int{
		7, 3, 9, 6,
	}
	data_sorted := SelectSortVictor0(data)
	fmt.Print(data_sorted)

}

// 选择排序
func SelectSortVictor0(data []int) []int {
	//先搞一个有序的数组, 然后逐个将元素移动到有序的数组中
	sortedList := make([]int, 0)
	sortedList = append(sortedList, data[0])
	for i := 1; i < len(data); i++ {
		sortedList = append(sortedList, data[i])
		// 倒序遍历sortedList, 移动data[i],到合适的位置
		for j := len(sortedList) - 1; j > 0; j-- {
			if sortedList[j] >= sortedList[j-1] {
				// 大了不用动, 直接跳出循环, 不用比较了
				break
			} else {
				// 后面的小了, 向前移动, 交换两个变量的值
				sortedList[j] = sortedList[j] ^ sortedList[j-1]
				sortedList[j-1] = sortedList[j] ^ sortedList[j-1]
				sortedList[j] = sortedList[j] ^ sortedList[j-1]
			}
		}
	}
	return sortedList
}
