package main

import "math/bits"

// https://space.bilibili.com/206214
func primeFactorization(k int) ([3]int, bool) {
	e2 := bits.TrailingZeros(uint(k))
	k >>= e2

	e3 := 0
	for k%3 == 0 {
		e3++
		k /= 3
	}

	e5 := 0
	for k%5 == 0 {
		e5++
		k /= 5
	}

	return [3]int{e2, e3, e5}, k == 1
}

func countSequences1(nums []int, k int64) int {
	e, ok := primeFactorization(int(k))
	if !ok { // k 有大于 5 的质因子
		return 0
	}

	n := len(nums)
	es := make([][3]int, n)
	for i, x := range nums {
		es[i], _ = primeFactorization(x)
	}

	type args struct{ i, e2, e3, e5 int }
	memo := map[args]int{}
	var dfs func(int, int, int, int) int
	dfs = func(i, e2, e3, e5 int) int {
		if i < 0 {
			if e2 == 0 && e3 == 0 && e5 == 0 { // k 变成 1
				return 1
			}
			return 0
		}
		p := args{i, e2, e3, e5}
		if res, ok := memo[p]; ok {
			return res
		}

		e := es[i]
		res1 := dfs(i-1, e2-e[0], e3-e[1], e5-e[2]) // k 除以 nums[i]
		res2 := dfs(i-1, e2+e[0], e3+e[1], e5+e[2]) // k 乘以 nums[i]
		res3 := dfs(i-1, e2, e3, e5)                // k 不变
		res := res1 + res2 + res3

		memo[p] = res
		return res
	}
	return dfs(n-1, e[0], e[1], e[2]) // 从 k 开始，目标是变成 1
}

func countSequences(nums []int, k int64) int {
	type args struct{ i, p, q int }
	memo := map[args]int{}
	var dfs func(int, int, int) int
	dfs = func(i, p, q int) int {
		if i < 0 {
			if p == int(k) && q == 1 {
				return 1
			}
			return 0
		}
		t := args{i, p, q}
		if res, ok := memo[t]; ok {
			return res
		}

		x := nums[i]
		g := gcd(p, q*x)
		res1 := dfs(i-1, p/g, q*x/g) // 除以 nums[i]
		g = gcd(p*x, q)
		res2 := dfs(i-1, p*x/g, q/g) // 乘以 nums[i]
		res3 := dfs(i-1, p, q)       // k 不变
		res := res1 + res2 + res3

		memo[t] = res
		return res
	}
	return dfs(len(nums)-1, 1, 1) // 从 1/1 开始，目标是变成 k/1
}

func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}
