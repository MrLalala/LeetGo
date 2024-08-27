package main

func minEnd2(n, x int) int64 {
	n-- // 先把 n 减一，这样下面讨论的 n 就是原来的 n-1
	// 还是双指针, 这边相当于是找x为0的位置填n对应的位置
	// 没找到一个空位n的指针++
	i, j := 0, 0
	for n>>j > 0 {
		// x 的第 i 个比特值是 0，即「空位」
		if x>>i&1 == 0 {
			// 空位填入 n 的第 j 个比特值
			x |= n >> j & 1 << i
			j++
		}
		i++
	}
	return int64(x)
}

func minEnd(n, x int) int64 {
	n-- // 先把 n 减一，这样下面讨论的 n 就是原来的 n-1
	// 还是双指针, 这边相当于是找x为0的位置填n对应的位置
	// 没找到一个空位n的指针++
	i, j := 0, 0
	for n>>j > 0 {
		// x 的第 i 个比特值是 0，即「空位」
		if x>>i&1 == 0 {
			// 空位填入 n 的第 j 个比特值
			x |= n >> j & 1 << i
			j++
		}
		i++
	}
	return int64(x)
}
