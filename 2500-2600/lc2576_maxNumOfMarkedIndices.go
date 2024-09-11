package main

import "slices"

func maxNumOfMarkedIndices(nums []int) int {
	slices.Sort(nums)
	i := 0
	for _, x := range nums[(len(nums)+1)/2:] {
		if nums[i]*2 <= x { // 找到一个匹配
			i++
		}
	}
	return i * 2
}
