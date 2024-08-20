package main

import (
	"math/bits"
	"sort"
)

func findMaximumNumber(k int64, x int) int64 {
	ans := sort.Search(int(k+1)<<x, func(num int) bool {
		num++
		n := bits.Len(uint(num)) // 多少位
		memo := make([][]int, n) // 记忆化搜索
		for i := range memo {
			memo[i] = make([]int, n+1)
			for j := range memo[i] {
				memo[i][j] = -1
			}
		}
		var dfs func(int, int, bool) int
		// i: 第i个二进制位
		// cnt1: 有多少满足条件的1的总和
		// limitHigh: 是否限制当前位置, 实际上只有0/1两种
		dfs = func(i, cnt1 int, limitHigh bool) (res int) {
			if i < 0 {
				return cnt1
			}
			if !limitHigh {
				p := &memo[i][cnt1]
				if *p >= 0 {
					return *p
				}
				defer func() { *p = res }()
			}
			up := 1
			if limitHigh {
				up = num >> i & 1
			}
			// 0是必然存在的
			res += dfs(i-1, cnt1, limitHigh && 0 == up)
			if up == 1 {
				// 1是可能存在的
				if (i+1)%x == 0 {
					cnt1++
				}
				res += dfs(i-1, cnt1, limitHigh)
			}
			return
		}
		// true  说明大了, 应该向左
		// false 说明小了, 应该往右
		// k 是额定的最大Accumulated Price, 需要不停的逼近
		// 数位dp是从高位向低位演进的.
		return dfs(n-1, 0, true) > int(k)
	})
	return int64(ans)
}
