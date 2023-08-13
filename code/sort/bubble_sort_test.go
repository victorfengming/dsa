package test

import (
	"fmt"
	"testing"
)

func TestBubble(t *testing.T) {
	// 构造数据
	data := []int{
		7, 3, 9, 6,
	}
	data_sorted := MaoPaoSort(data)
	fmt.Print(data_sorted)
}

// 冒泡排序
// Written by victor

func MaoPaoSort(data []int) []int {
	// 遍历数组,每次把最大的移动到数组的最后面
	for j := len(data) - 1; j > 0; j-- {
		for i := 0; i < j; i++ {
			if data[i] <= data[i+1] {
				// 如果左边的小, 就不动
			} else {
				// 如果左边的大,就交换数值
				data[i] = data[i] ^ data[i+1]
				data[i+1] = data[i] ^ data[i+1]
				data[i] = data[i] ^ data[i+1]
			}
		}
	}
	return data
}
