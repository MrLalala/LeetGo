package main

func maxStrength(nums []int) int64 {
	mn, mx := nums[0], nums[0]
	for _, x := range nums[1:] {
		// 忽略x的正负值, 选取最小值/最大值
		// 不选, 选一个值, 选取最大/最小*x
		mn, mx = min(mn, x, mn*x, mx*x),
			max(mx, x, mn*x, mx*x)
	}
	return int64(mx)
}
