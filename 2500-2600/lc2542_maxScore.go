package main

import (
	"container/heap"
	"sort"
)

func maxScore(nums1, nums2 []int, k int) int64 {
	// x:nums1[i], y:nums2[i]
	type pair struct{ x, y int }
	a := make([]pair, len(nums1))
	sum := 0
	for i, x := range nums1 {
		a[i] = pair{x, nums2[i]}
	}
	// 按照nums2从大到小排序
	sort.Slice(a, func(i, j int) bool { return a[i].y > a[j].y })

	h := hp{nums2[:k]} // 复用内存
	for i, p := range a[:k] {
		sum += p.x
		// 堆按照nums1从小到大排序
		h.IntSlice[i] = p.x
	}
	ans := sum * a[k-1].y
	heap.Init(&h)
	for _, p := range a[k:] {
		if p.x > h.IntSlice[0] {
			// sum表示的是替换了已选择的nums1中的最小值后的和
			sum += p.x - h.replace(p.x)
			// 为什么要乘以p.y呢? 因为a是按照nums2从大到小排序的, 所以p.y就是当前选择的nums2中的最小值
			ans = max(ans, sum*p.y)
		}
	}
	return int64(ans)
}

type hp struct{ sort.IntSlice }

func (h hp) Pop() (_ interface{}) { return }
func (h hp) Push(interface{})     {}
func (h hp) replace(v int) int {
	top := h.IntSlice[0]
	h.IntSlice[0] = v
	heap.Fix(&h, 0)
	return top
}
func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}