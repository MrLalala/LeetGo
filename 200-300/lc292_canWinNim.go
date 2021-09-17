package main

func canWinNim(n int) bool {

	/*
	   简单的论证一下:
	   先手可拿[1,3]个包, 这就意味着 当有4个包时, 后手一定赢
	   所以, 当包的数量不是4的整数倍的时候, 可以留给后手4N个包, 这样就能保证最终一定可以胜利

	   举个例子: 7个包的时候, 现手拿3个, 剩余4个, 那么后手无论拿多少, 我方再次出手时一定能拿到最后一个


	*/

	return n%4 != 0
}
