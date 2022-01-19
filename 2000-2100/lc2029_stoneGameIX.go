package main

func stoneGameIX(stones []int) bool {

	// 将每个石头的价值和3取余
	// cnt0: 不会对总和产生影响, 但是奇偶性会影响先后手
	//       比如, 如果发现移除价值为1/2的石头会失败, 那么就可以通过移除0来交换先后手
	//       当数量为偶数时, 胜负情况等同于为0的情况
	//       当数量为奇数时, 胜负情况等同于为1的情况
	// cnt1, cnt2
	//       先手1, 对应的序列应该是 1121212121
	//       先手2, 对应的序列应该是 2212121212

	// 当 cnt1 = 1, cnt2 >= 1时,(12.222) 直接获胜
	// 当 cnt1 >= 2, 此时和cnt2进行对比
	//      cnt1 - cnt2 == 1: (112.), 石子全被移除, 输
	//      cnt1 - cnt2 == 2: (1121.), 石子全被移除, 输
	//      cnt1 - cnt2 > 2: (11211.), 没有2可以移除, 因为取到了3的倍数所以对手输了
	//      cnt1 <= cnt2: (112122.), 对手必须要选择一个满足3的石子, 赢
	var cnt0, cnt1, cnt2 int
	for _, val := range stones {
		val %= 3
		if val == 0 {
			cnt0++
		} else if val == 1 {
			cnt1++
		} else {
			cnt2++
		}
	}
	if cnt0%2 == 0 {
		// 12.22 赢
		// 1121. 赢
		// 全2或者全1的情况一定会输的
		return cnt1 >= 1 && cnt2 >= 1
	}
	// 奇数个, 存在一个3可以进行换手
	// 对应上述的4种情况里3可以获胜(只能是因为对手取到了3的倍数失败才算是换手的胜利)
	return cnt1-cnt2 > 2 || cnt2-cnt1 > 2
}
