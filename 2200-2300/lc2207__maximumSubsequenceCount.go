package main

func maximumSubsequenceCount2(text string, pattern string) int64 {
	// 正向计算 pattern[0], 逆向计算 pattern[1]
	// 根据乘法原理, 计算最大值?

	// ababab = a*3+a*2+a*1
	// 开头放a: a*3*2+a*2+a*1
	// 结尾放b: a*4+a*3+a*2
	// 所以, 要么放开头, 要么放结尾呗?

	var s, e = pattern[0], pattern[1]
	var bs = []byte(text)[:0]
	for i := range text {
		c := text[i]
		if c == s || c == e {
			bs = append(bs, c)
		}
	}
	vlen := len(bs)
	if vlen <= 1 {
		return int64(vlen)
	}
	if s == e {
		return int64(vlen * (vlen + 1) / 2)
	}
	var rightCnt = make([]int, len(bs)+1)
	for i := len(bs) - 1; i >= 0; i-- {
		rightCnt[i] = rightCnt[i+1]
		if bs[i] == e {
			rightCnt[i]++
		}
	}
	// fmt.Println(string(bs), rightCnt)
	// 计算开头+s和结尾+e的区别
	var addS, addE int
	addS += rightCnt[0]
	for i := range bs {
		c := bs[i]
		if c == s {
			addS += rightCnt[i]
			addE += rightCnt[i] + 1
			// fmt.Println(addS, addE)
		}
	}
	return int64(max(addS, addE))
}

func maximumSubsequenceCount(text, pattern string) (ans int64) {
	x, y := pattern[0], pattern[1]
	cntX, cntY := 0, 0
	for i := range text {
		c := text[i]
		if c == y {
			// 每出现一次y, 就增加cntX个子序列
			ans += int64(cntX)
			cntY++
		}
		if c == x {
			cntX++
		}
	}
	// x放最左边, 增加cntY个子序列;
	// y放最右边, 增加cntX个子序列;
	return ans + int64(max(cntX, cntY))
}
