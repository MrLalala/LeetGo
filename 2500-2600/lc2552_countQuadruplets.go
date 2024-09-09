package main

import "slices"

func countQuadruplets(nums []int) (ans int64) {
	n := len(nums)
	// great[k][j]: 在位置k之后比j大的数有多少个
	great := make([][]int, n)
	// 聪明啊..
	// 完美绕过了累加
	// 首先, 比 n-1 大的数是0个
	great[n-1] = make([]int, n+1)
	// 从后向前迭代分界点k, 依次计算有多少个比nums[k+1]小的数
	for k := n - 2; k >= 2; k-- {
		// 通过clone, 继承k+1的结果, 这次迭代只需要关心nums[k+1]的情况
		great[k] = slices.Clone(great[k+1])
		// 从前向后迭代x
		// 迭代的上限为什么是nums[k+1]呢? 是因为通过k是分界点, 小于nums[k+1]的数就是[1, nums[k+1])
		// 所以这里就保证了 great[k] 中的数是小于 nums[k+1] 的个数
		// 由此对于每一个x(j), 都可以计算出在 k 右边大于 j 的数有多少个
		// 注意核心是 nums[k+1], 由此确定了有多少个数是小于 nums[k+1] 的
		for x := 1; x < nums[k+1]; x++ {
			great[k][x]++
		}
	}

	less := make([]int, n+1)
	// i, j, k, l
	// 在k右边大于nums[j]的
	// 在j左边小于nums[k]的
	// 确定 j和k 的位置, 通过乘法原理, less[j][nums[k]] * great[k][nums[j]] 就是满足上升四元组的个数
	for j := 1; j < n-2; j++ {
		for x := nums[j-1] + 1; x <= n; x++ {
			less[x]++
		}
		for k := j + 1; k < n-1; k++ {
			if nums[j] > nums[k] {
				ans += int64(less[nums[k]] * great[k][nums[j]])
			}
		}
	}
	return
}
