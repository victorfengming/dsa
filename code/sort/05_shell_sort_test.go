package test

import (
	"fmt"
	"testing"
)

func shellSort(arr []int) {
	n := len(arr)
	gap := n / 2 // 初始间隔为数组长度的一半

	// 循环直到间隔为1
	for gap > 0 {
		// 对每个间隔进行插入排序
		for i := gap; i < n; i++ {
			temp := arr[i]
			j := i

			// 在子序列中应用插入排序算法，以使每个子序列基本有序
			for j >= gap && arr[j-gap] > temp {
				// 如果前面的元素较大，将较大元素后移
				arr[j] = arr[j-gap]
				j -= gap
			}

			// 将当前元素插入到合适的位置
			arr[j] = temp
		}
		gap /= 2 // 缩小间隔
	}
}

func TestShellSort(t *testing.T) {

	arr := []int{12, 34, 54, 2, 3}
	fmt.Println("排序前:", arr)

	shellSort(arr)

	fmt.Println("排序后:", arr)
}
