package main

func timeRequiredToBuy2(tickets []int, k int) int {
	var tick int
	// 迭代数组, 找到第一个>0的位置
	ln := len(tickets)
	for i := 0; ; i++ {
		ni := i % ln
		if n := tickets[ni]; n > 0 {
			tick++
			tickets[ni]--
			if n == 1 && ni == k {
				return tick
			}
		}
	}
	return -1
}

func timeRequiredToBuy(tickets []int, k int) (ans int) {
	tk := tickets[k]
	// 类似于贪心? 计算前后最多会
	for i, t := range tickets {
		if i <= k {
			// k前边的人, 最多买票tk张
			ans += min(t, tk)
		} else {
			// k后边的人, 最多买票tk-1张
			ans += min(t, tk-1)
		}
	}
	return
}
