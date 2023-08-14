package test

import "testing"

func TestQuickSort(t *testing.T) {

}

// 快速排序, 这个排序也是递归, 只不过这个递归和归并排序的递归不太一样的地方是, 快速排序主要操作是在调用递归函数之前,而归并排序是在调用递归函数之后的merge操作中
func QuickSortVictor(data []int) []int {
	// 快速排序, 从数组中选区一个基准点,将数组中的数分为两部分
	//一部分小于基准, 一部分大于基准
	//middle_index := len(data) / 2
	//遍历data, 将比这个值小的放到一个数组里, 将比这个值大的放到另一个数组中
	//mini_slice := make([]int, len(data))
	//
	return nil
}

// 快速排序算法
func quickSortGPT(data []int) []int {
	// 如果数组长度小于等于1，直接返回，无需排序
	if len(data) <= 1 {
		return data
	}

	// 选择最后一个元素作为基准
	pivot := data[len(data)-1]
	left := make([]int, 0)
	right := make([]int, 0)

	// 将小于等于基准的元素放在左边，大于基准的元素放在右边
	for _, num := range data[:len(data)-1] {
		if num <= pivot {
			left = append(left, num)
		} else {
			right = append(right, num)
		}
	}

	// 递归地对左右两部分进行快速排序
	sortedLeft := quickSortGPT(left)
	sortedRight := quickSortGPT(right)

	// 合并左、基准和右三个部分，得到排序结果
	return append(append(sortedLeft, pivot), sortedRight...)
}
