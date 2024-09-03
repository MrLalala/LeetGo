package main

import "slices"

func countWays(nums []int) (ans int) {
	slices.Sort(nums)
	if nums[0] > 0 { // 一个学生都不选
		ans = 1
	}
	for i := 1; i < len(nums); i++ {
		if nums[i-1] < i && i < nums[i] {
			ans++
		}
	}
	// 最后选择所有的
	return ans + 1 // 一定可以都选
}
