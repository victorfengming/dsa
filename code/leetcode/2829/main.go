package main

func minimumSum(n int, k int) int {
	j := 1
	// 初始化一个变量用于存储总和
	sum := 0
	jmap := make(map[int]interface{}, 0)
	i := 0
	for {
		if i >= n {
			break
		}
		// 先判断j是不是比k大了如果比k大了就不会相等,直接存就行了
		if j >= k {
			sum += j
			j += 1
			i += 1
			continue
		}
		if _, ok := jmap[j]; ok {

		} else {
			sum += j
			jmap[k-j] = nil
			i += 1
		}
		j += 1
	}
	return sum
}

func main() {
	minimumSum(5, 4)
}
