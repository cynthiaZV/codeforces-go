package main

// https://space.bilibili.com/206214
var fac = [10]int{1}

func init() {
	// 预处理阶乘
	for i := 1; i < len(fac); i++ {
		fac[i] = fac[i-1] * i
	}
}

func isDigitorialPermutation(n int) bool {
	sumFac := 0
	cnt := [10]int{}
	for ; n > 0; n /= 10 {
		d := n % 10
		sumFac += fac[d]
		cnt[d]++
	}

	for ; sumFac > 0; sumFac /= 10 {
		cnt[sumFac%10]--
	}

	// cnt[i] == 0
	return cnt == [10]int{}
}
