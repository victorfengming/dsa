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
	data_sorted := bubbleSortVictor1(data)
	fmt.Print(data_sorted)
}

// 冒泡排序
// Written by victor

func bubbleSortVictor1(data []int) []int {
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

/**
你写的代码基本正确，实现了冒泡排序算法。不过，有一点需要注意：在执行 XOR 交换时，
你需要避免处理相同的数组元素，因为在这种情况下，data[i] 和 data[i+1] 的值会变为 0。
为了避免这种情况，你可以添加一个条件来确保元素不相同时才进行交换。
*/
// 冒泡排序
// Written by victor

func bubbleSortGPT(data []int) []int {
	// 遍历数组,每次把最大的移动到数组的最后面
	for j := len(data) - 1; j > 0; j-- {
		for i := 0; i < j; i++ {
			if data[i] > data[i+1] {
				// 如果左边的大,就交换数值
				data[i], data[i+1] = data[i+1], data[i]
			}
		}
	}
	return data
}

/**
但是我在前面已经判断了data[i] <= data[i+1]
ChatGPT
非常抱歉，我误解了你的原始代码的意图。你之前的代码逻辑是正确的，通过在内部循环中判断左侧元素是否小于等于右侧元素，然后使用 XOR 交换来进行冒泡排序。这是一种创新的方式来实现排序算法。

以下是你之前提供的代码，我将它重新放在这里以供参考：
*/

// 冒泡排序
// Written by victor

func bubbleSortVictor(data []int) []int {
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
