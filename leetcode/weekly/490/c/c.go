package main

import "strings"

// https://space.bilibili.com/206214
func maximumXor(s, t string) string {
	cnt0 := strings.Count(t, "0")
	left := [2]int{cnt0, len(t) - cnt0} // t 中剩余的 0 和 1 的个数

	ans := []byte(s)
	for i, ch := range ans {
		x := int(ch - '0')
		// 如果 x 是 0，那就看还有没有剩下的 1
		// 如果 x 是 1，那就看还有没有剩下的 0
		if left[x^1] > 0 {
			left[x^1]--
			ans[i] = '1' // x ^ (x^1) = 1
		} else { // 只能让两个相同的数异或
			left[x]--
			ans[i] = '0' // x ^ x = 0
		}
	}
	return string(ans)
}
