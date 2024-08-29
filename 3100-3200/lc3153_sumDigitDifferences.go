package main

func sumDigitDifferences(nums []int) int64 {
	// 统计每个位置上的数量, 然后总数-数量=剩余的数字
	var cnt [10][10]int

	var ret int
	for k, num := range nums {
		bit := 0
		for num > 0 {
			s := num % 10
			// 这个减法用的喵啊
			ret += k - cnt[bit][s]
			cnt[bit][s]++
			bit++
			num /= 10
		}
	}
	return int64(ret)
}
