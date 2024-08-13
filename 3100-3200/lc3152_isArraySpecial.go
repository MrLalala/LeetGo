package main

import "sort"

func isArraySpecial(nums []int, queries [][]int) []bool {
	// var queueGroup = make([]int, len(nums))
	queueGroup := nums
	var pre = nums[0]
	queueGroup[0] = 0
	for i, num := range nums[1:] {
		queueGroup[i+1] = queueGroup[i]
		if num&1 == pre&1 {
			queueGroup[i+1]++
		}
		pre = num
	}

	var ret = make([]bool, len(queries))
	for i, query := range queries {
		start, end := query[0], query[1]
		ret[i] = queueGroup[start] == queueGroup[end]
	}
	return ret
}

func isArraySpecial2(nums []int, queries [][]int) []bool {
	var queueGroup []int
	queueGroup = append(queueGroup, 0)
	for i, num := range nums[1:] {
		if num&1 == nums[i]&1 {
			// 奇偶性相同, 当前位置单独成组
			queueGroup = append(queueGroup, i+1)
		}
	}
	queueGroup = append(queueGroup, len(nums))

	// queueGroup = append(queueGroup, len(nums))
	var ret = make([]bool, len(queries))
	for i, query := range queries {
		start, end := query[0], query[1]
		if start == end {
			ret[i] = true
			continue
		}
		// 一定是不会超过上限的
		startGroup := sort.SearchInts(queueGroup, start)
		if startGroup >= len(queueGroup) {
			ret[i] = false
			continue
		}
		if queueGroup[startGroup] != start {
			ret[i] = end < queueGroup[startGroup]
		} else {
			ret[i] = end < queueGroup[startGroup+1]
		}
	}
	return ret
}
