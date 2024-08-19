package main

func waysToReachStair(k int) int {
	type args struct {
		i    int  // 第几个位置
		op2  int  // 进行了几次操作2
		pre1 bool // 上一次是不是操作1
	}

	// 记忆化, 对应参数的次数
	var cache = make(map[args]int)

	var dfs func(int, int, bool) int
	dfs = func(i, op2 int, pre1 bool) int {
		if i > k+1 {
			// 没有连续的2次1, 往后只会越来越远
			// 提前枝减
			return 0
		}
		p := args{i: i, op2: op2, pre1: pre1}
		if n, ok := cache[p]; ok {
			return n
		}
		// 操作2: 跳1<<op2步, 并且op2+1
		res := dfs(i+1<<op2, op2+1, false)
		if i > 0 && !pre1 {
			res += dfs(i-1, op2, true)
		}
		if i == k {
			// 如果当前已经到达了k, 那么就是一种方法
			res++
		}
		cache[p] = res
		return res
	}
	return dfs(1, 0, false)
}
