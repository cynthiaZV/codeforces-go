package main

import "math/bits"

// github.com/EndlessCheng/codeforces-go
func concatenatedBinary1(n int) (ans int) {
	const mod = 1_000_000_007
	for i := 1; i <= n; i++ {
		w := bits.Len(uint(i))
		ans = (ans<<w | i) % mod
	}
	return
}

const mod = 1_000_000_007

func pow(x, n int) int {
	res := 1
	for ; n > 0; n /= 2 {
		if n%2 > 0 {
			res = res * x % mod
		}
		x = x * x % mod
	}
	return res
}

func concatenatedBinary(n int) (ans int) {
	for w := 1; 1<<(w-1) <= n; w++ {
		l := 1 << (w - 1)
		r := min(1<<w-1, n)
		m := r - l + 1
		q := 1 << w
		powQ := pow(q, m)
		invQ1 := pow(q-1, mod-2)
		s := r*(powQ-1)%mod*invQ1 - (q-m*powQ+(m-1)*powQ%mod*q)%mod*invQ1%mod*invQ1
		ans = (ans*powQ + s) % mod
	}
	return (ans + mod) % mod // 保证结果非负
}
