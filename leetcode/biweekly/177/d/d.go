package main

// https://space.bilibili.com/206214
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

func sumOfNumbers(l, r, k int) int {
	m := r - l + 1
	return (l + r) * m * (pow(10, k) - 1 + mod) % mod * pow(18, mod-2) % mod * pow(m, k-1) % mod
}
