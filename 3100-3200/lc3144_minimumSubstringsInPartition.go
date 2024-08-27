package main

func minimumSubstringsInPartition(s string) int {
	var dp = make([]int, len(s)+1)
	for i := range dp {
		dp[i] = i
	}
	for i := 1; i <= len(s); i++ {
		var cnt [26]int
		var k, maxCnt int
		// s[:i] = s[:j] + s[j:i], 让s[j:i]尽可能的长, 那么s[:j]就尽可能的短
		for j := i - 1; j >= 0; j-- {
			idx := s[j] - 'a'
			if cnt[idx] == 0 {
				k++ // 首次出现
			}
			cnt[idx]++
			maxCnt = max(maxCnt, cnt[idx]) // 这一组中的最大值
			if i-j == k*maxCnt {
				dp[i] = min(dp[i], dp[j]+1) // [j:i]可以合并
			}
		}
	}
	return dp[len(s)]
}
