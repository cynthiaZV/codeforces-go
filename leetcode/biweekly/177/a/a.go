package main

import (
	"math"
)

// https://space.bilibili.com/206214
func minDistinctFreqPair(nums []int) []int {
	cnt := map[int]int{}
	mn := math.MaxInt
	for _, x := range nums {
		cnt[x]++
		mn = min(mn, x)
	}

	cntMin := cnt[mn]
	minY := math.MaxInt
	for y, c := range cnt {
		if c != cntMin {
			minY = min(minY, y)
		}
	}

	if minY == math.MaxInt {
		return []int{-1, -1}
	}
	return []int{mn, minY}
}
