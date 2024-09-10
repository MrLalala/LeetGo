package main

func maximizeWin(prizePositions []int, k int) (ans int) {
	n := len(prizePositions)
	// 如果区间恰好覆盖了所有的奖品
	if k*2+1 >= prizePositions[n-1]-prizePositions[0] {
		return n
	}
	// 之所以可以这么计算, 是因为允许重叠!
	mx, left, right := 0, 0, 0
	for mid, p := range prizePositions {
		// 把 prizePositions[mid] 视作第二条线段的左端点，计算第二条线段可以覆盖的最大奖品下标
		for right < n && prizePositions[right]-p <= k {
			right++
		}
		// 循环结束后，right-1 是第二条线段可以覆盖的最大奖品下标
		ans = max(ans, mx+right-mid)
		// 把 prizePositions[mid] 视作第一条线段的右端点，计算第一条线段可以覆盖的最小奖品下标
		for p-prizePositions[left] > k {
			left++
		}
		// 循环结束后，left 是第一条线段可以覆盖的最小奖品下标
		// 实际上, 在 p 不超过窗口之前, 这个值一直是1
		// 这也是为什么需要将前半部分后置的原因, 可以避免在第一轮循环中的重复计算
		mx = max(mx, mid-left+1)
	}
	return
}
