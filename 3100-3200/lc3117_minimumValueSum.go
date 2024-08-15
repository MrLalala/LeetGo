package main

import "math"

func minimumValueSum(nums []int, andValues []int) int {
	numLength, valueLength := len(nums), len(andValues)
	if numLength < valueLength {
		return -1
	}
	type args struct{ numIdx, valueIdx, andResult int }
	cache := make(map[args]int)

	const INF = math.MaxInt32 / 2

	var dfs func(int, int, int) int
	dfs = func(numIdx, valueIdx, andResult int) int {
		if numLength-numIdx < valueLength-valueIdx {
			return INF
		}
		// 如果任意一个数组遍历完了，那么剩下的元素必须全部选择
		if numIdx == numLength || valueIdx == valueLength {
			// 判断是否有剩余的元素
			if valueLength-valueIdx+numLength-numIdx == 0 {
				return 0
			}
			return INF
		}

		andResult &= nums[numIdx]
		a := args{numIdx, valueIdx, andResult}
		if n, ok := cache[a]; ok {
			return n
		}
		// 不选择当前的nums[numIdx]作为分割点
		ret := dfs(numIdx+1, valueIdx, andResult)
		if andResult == andValues[valueIdx] {
			// 只有当andResult == andValues[valueIdx]时，才能选择当前的nums[numIdx]作为分割点
			// 记录数组的"值"
			ret = min(ret, dfs(numIdx+1, valueIdx+1, -1)+nums[numIdx])
		}
		cache[a] = ret
		return ret
	}
	if ret := dfs(0, 0, -1); ret == INF {
		return -1
	} else {
		return ret
	}
}
