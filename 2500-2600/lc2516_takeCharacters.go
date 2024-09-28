package main

func takeCharacters(s string, k int) int {
	// 外部窗口
	cnt := [3]int{}
	for _, c := range s {
		cnt[c-'a']++ // 一开始，把所有字母都取走
	}
	if cnt[0] < k || cnt[1] < k || cnt[2] < k {
		return -1 // 字母个数不足 k
	}
	// 起始情况下, 外部窗口的数量是肯定满足k的
	// 但是在移动右指针的过程中, 一旦出现了某一个字符不满足k个, 就需要移动左指针来补全
	// 因为在这一刻之前, 其他字符都是 >= k的, 之后再移动左指针只会让该字符越来越多.

	mx, left := 0, 0
	for right, c := range s {
		c -= 'a'
		cnt[c]--         // 移入窗口，相当于不取走 c
		for cnt[c] < k { // 窗口之外的 c 不足 k
			cnt[s[left]-'a']++ // 移出窗口，相当于取走 s[left]
			left++
		}
		mx = max(mx, right-left+1)
	}
	return len(s) - mx
}
