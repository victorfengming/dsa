package test

import (
	"fmt"
	"testing"
)

func TestMergingSort(t *testing.T) {
	// 构造数据
	data := []int{
		7, 3, 9, 6,
	}
	data_sorted := mergingSortVictor(data)
	fmt.Print(data_sorted)
}

func mergingSortVictor(data []int) []int {
	// 判断是否是一个元素, 如果是一个元素直接返回
	if len(data) == 1 {
		return data
	}
	// 分成左右两部分,分别排序
	left := data[:len(data)/2]
	right := data[len(data)/2:]
	leftList := mergingSortVictor(left)
	rightList := mergingSortVictor(right)
	// 合并两个数组
	//双指针遍历
	res := make([]int, 0)
	left_index := 0
	right_index := 0
	for {
		// 跳出条件
		if left_index == len(leftList) && right_index == len(rightList) {
			break
		}
		if leftList[left_index] < rightList[right_index] {
			res = append(res, leftList[left_index])
			left_index += 1
		} else {
			res = append(res, rightList[right_index])
			right_index += 1
		}
		// 如果哪个移动不了了, 就把另一个全搞过去
		if left_index == len(leftList) {
			// 把右拼接到结果中
			res = append(res, rightList[right_index:]...)
			break
		}
		// 如果哪个移动不了了, 就把另一个全搞过去
		if right_index == len(rightList) {
			// 把右拼接到结果中
			res = append(res, leftList[left_index:]...)
			break
		}

	}
	return res
}

/**
是的，你的归并排序实现看起来是正确的，包含了分割和合并两个步骤。你将输入数组递归地分割成左右两部分，然后使用双指针遍历的方式在合并阶段将两个有序的子数组合并成一个有序的数组。这是一个经典的归并排序算法。

这里的代码逻辑和步骤都是正确的，你已经考虑到了各种情况，如左边数组或右边数组遍历完后的处理等。

总之，你的归并排序实现是有效的，可以用于排序整数切片。如果你测试并验证了代码，确保它在各种情况下都能正常工作，那么就可以放心使用了。

-----------------------------------------------------

在这个优化后的版本中，我将合并操作独立到了一个函数 mergeSortedLists 中，这有助于提高代码的可读性和模块性。

我还为 result 分片预分配了足够的容量，避免了在每次追加元素时都进行内存重新分配的开销。

另外，我将递归的中点计算提前到一处地方，避免在每次递归调用时都计算一次。

这些优化可能不会显著改变算法的性能，但可以使代码更清晰、可读性更高，并且在一些情况下稍微减少一些开销。
*/

func mergingSortGPT(data []int) []int {
	if len(data) <= 1 {
		return data
	}

	mid := len(data) / 2
	left := mergingSortVictor(data[:mid])
	right := mergingSortVictor(data[mid:])

	return mergeSortedListsGPT(left, right)
}

func mergeSortedListsGPT(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))

	leftIndex, rightIndex := 0, 0
	for leftIndex < len(left) && rightIndex < len(right) {
		if left[leftIndex] <= right[rightIndex] {
			result = append(result, left[leftIndex])
			leftIndex++
		} else {
			result = append(result, right[rightIndex])
			rightIndex++
		}
	}

	result = append(result, left[leftIndex:]...)
	result = append(result, right[rightIndex:]...)

	return result
}
