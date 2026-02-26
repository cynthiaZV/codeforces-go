package main

import (
	"fmt"
	"strconv"
)

// github.com/EndlessCheng/codeforces-go
func findDifferentBinaryString1(nums []string) string {
	n := len(nums)
	has := make(map[int]bool, n)
	for _, s := range nums {
		x, _ := strconv.ParseInt(s, 2, 64)
		has[int(x)] = true
	}

	ans := 0
	for has[ans] {
		ans++
	}

	return fmt.Sprintf("%0*b", n, ans)
}

func findDifferentBinaryString(nums []string) string {
	ans := make([]byte, len(nums))
	for i, s := range nums {
		ans[i] = s[i] ^ 1
	}
	return string(ans)
}
