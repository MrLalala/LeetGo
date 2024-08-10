package main

import "sort"

func leftmostBuildingQueries(heights []int, queries [][]int) []int {
	ans := make([]int, len(queries))
	type pair struct{ h, i int }
	qs := make([][]pair, len(heights))
	for i, q := range queries {
		a, b := q[0], q[1]
		if a > b {
			a, b = b, a // 保证 a <= b
		}
		if a == b || heights[a] < heights[b] {
			ans[i] = b // a 直接跳到 b
		} else {
			// 入队的前提是 heights[a] >= heights[b]
			// 所以进行离线查询时, 只需要保证 target > heights[a] 即可
			qs[b] = append(qs[b], pair{heights[a], i}) // 离线询问
		}
	}

	// 递减栈,
	var st []int
	// 从后往前遍历, 栈中的元素始终是在要查询的元素的右边
	for i := len(heights) - 1; i >= 0; i-- {
		for _, q := range qs[i] {
			// 此时存储的值中, heights[a] > heights[b]
			// 同时存储的q.h也是 heights[a], 需要从当前的递减栈中, 找出首个 >= heights[a]的值
			// st[i] <= q.h 向左, 左边更大
			// st[i] >  q.h 向右, 右边更小
			j := sort.Search(len(st), func(i int) bool { return heights[st[i]] <= q.h }) - 1
			if j >= 0 {
				ans[q.i] = st[j]
			} else {
				ans[q.i] = -1
			}
		}
		for len(st) > 0 && heights[i] >= heights[st[len(st)-1]] {
			st = st[:len(st)-1]
		}
		st = append(st, i)
	}
	return ans
}
