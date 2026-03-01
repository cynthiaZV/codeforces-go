package main

import (
	"math"
	"slices"
)

// https://space.bilibili.com/206214
func makeParityAlternating(nums []int) []int {
	if len(nums) == 1 {
		return []int{0, 0}
	}

	gMin := slices.Min(nums)
	gMax := slices.Max(nums)

	f := func(target int) (int, int) {
		op, mn, mx := 0, math.MaxInt, math.MinInt
		for i, x := range nums {
			if (x-i)&1 != target { // 等价于 x&1 != target ^ i%2
				op++
				if x == gMin {
					x++
				} else if x == gMax {
					x--
				}
			}
			mn = min(mn, x)
			mx = max(mx, x)
		}
		return op, max(mx-mn, 1) // 在 n >= 2 的情况下，极差至少是 1
	}

	op1, minD1 := f(0)
	op2, minD2 := f(1)

	if op1 < op2 || op1 == op2 && minD1 < minD2 {
		return []int{op1, minD1}
	}
	return []int{op2, minD2}
}
