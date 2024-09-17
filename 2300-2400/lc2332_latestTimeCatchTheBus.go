package main

import "slices"

func latestTimeCatchTheBus(buses, passengers []int, capacity int) (ans int) {
	slices.Sort(buses)
	slices.Sort(passengers)

	// 让所有人都上车
	var pi int
	var cnt int
	for _, bus := range buses {
		cnt = capacity
		for pi < len(passengers) && passengers[pi] <= bus && cnt > 0 {
			cnt--
			pi++
		}
	}

	var ret int
	if cnt > 0 {
		// 如果还有剩余的座位, 那就发车时再来?
		ret = buses[len(buses)-1]
	} else {
		// 没有剩余位置了, 啥时候来呢..?
		// 卡着最后一名乘客
		ret = passengers[pi-1]
	}

	// 倒序迭代[pi-1, 0],
	for pi--; pi >= 0 && passengers[pi] == ret; pi-- {
		ret--
	}
	return ret
}
