package main

func maximumRobots(chargeTimes, runningCosts []int, budget int64) (ans int) {
	var q []int
	sum := int64(0)
	left := 0
	for right, t := range chargeTimes {
		// 1. 入, 类似于滑动窗口的最大值
		for len(q) > 0 && t >= chargeTimes[q[len(q)-1]] {
			q = q[:len(q)-1]
		}
		q = append(q, right)
		// 窗口的区间和
		sum += int64(runningCosts[right])

		// 2. 出, 如果超标了, 就需要递增左指针
		for len(q) > 0 && int64(chargeTimes[q[0]])+int64(right-left+1)*sum > budget {
			if q[0] == left {
				q = q[1:]
			}
			sum -= int64(runningCosts[left])
			left++
		}

		// 3. 更新答案
		ans = max(ans, right-left+1)
	}
	return
}
