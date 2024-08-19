package main

func checkRecord(n int) int {
	const MOD = 1e9 + 7

	if n == 0 {
		return 1
	}
	//  对于某个长度为i的记录序列，以下变量表示对应情况的序列数量
	//  P:  序列中没有A，最新的一个记录不是L
	//  AP: 序列中有过A，最新的一个记录不是L
	//  L:  序列中没有A，最新的一个记录是L
	//  AL: 序列中有过A，最新的一个记录是L
	//  LL: 序列中没有A，最新的两个记录是LL
	//  ALL:序列中有过A，最新的两个记录是LL
	//  A:  最新的一个记录是A

	// 长度为1时, P, L, A 都算作1次
	// 穷举所有基础情况, 然后进行DP
	var P, AP, L, LL, AL, ALL, A = 1, 0, 1, 0, 0, 0, 1

	for i := 2; i <= n; i++ {
		P, AP, L, LL, AL, ALL, A =
			(P+L+LL)%MOD, // i位置是P的时候, 并且前边没有A的情况, 可以由P, L, LL转移过来
			(AP+AL+ALL+A)%MOD, // i位置是P的时候, 并且前边有A的情况, 可以由AP, AL, ALL, A转移过来
			P, // i位置是L的时候, 前边不能有A和L, 只能由P转移过来
			L, // i位置是L的时候, 连续两个L, 只能由L转移过来
			(AP+A)%MOD, // i位置是L的时候, 并且前边有一个A的情况, 可以由AP, A转移过来
			AL, // i位置是L的时候, ALL只能由AL转移过来
			(P+L+LL)%MOD // i位置是A的时候, 前边不能存在A, 只能由P, L, LL转移过来
	}

	return (P + AP + L + LL + AL + ALL + A) % MOD
}
