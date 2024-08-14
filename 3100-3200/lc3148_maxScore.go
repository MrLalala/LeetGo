package main

import "math"

func maxScore(grid [][]int) int {
	ans := math.MinInt
	colMin := make([]int, len(grid[0]))
	for i := range colMin {
		colMin[i] = math.MaxInt
	}
	for _, row := range grid {
		preMin := math.MaxInt // colMin[0..j] 的最小值, 可以理解为是left
		for j, x := range row {
			// colMin[j] 可以理解为是 top
			// 到当前位置上的最优解等同于当前位置的高度 - min(left, top)
			ans = max(ans, x-min(preMin, colMin[j]))
			// 更新在竖直方向的最小值 min(top, x)
			colMin[j] = min(colMin[j], x)
			// 更新在水平方向上的最小值, 也可以理解为是 min(left, top, x)
			preMin = min(preMin, colMin[j])
		}
	}
	return ans
}
