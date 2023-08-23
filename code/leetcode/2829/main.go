package main

func minimumSum(n int, k int) int {
	//从1开始遍历
	// 使用索引遍历字符串
	j := 1
	jslice := make([]int, 0)
	for i := 0; i < n; i++ {
		// 先判断j是不是比k大了如果比k大了就不会相等,直接存就行了
		if j >= k {
			jslice = append(jslice, j)
			j += 1
			continue
		}

	mark:
		// 判断k是否和前面的能够组成和为k的
		// 遍历jslice
		// 使用循环遍历空切片
		for _, value := range jslice {
			if j+value == k {
				j += 1
				goto mark
			}
		}
		// 如果符合条件, 就存放到数组中
		jslice = append(jslice, j)
		// 累加
		j += 1
	}
	// 初始化一个变量用于存储总和
	sum := 0

	// 遍历数组并将值相加
	for _, value := range jslice {
		sum += value
	}
	return sum
}
