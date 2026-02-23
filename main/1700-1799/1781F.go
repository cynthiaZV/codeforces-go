package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1781F(in io.Reader, out io.Writer) {
	const mod = 998244353
	const mx = 10000
	inv := [mx + 1]int{}
	inv[1] = 1
	for i := 2; i <= mx; i++ {
		inv[i] = (mod - mod/i) % mod * inv[mod%i] % mod
	}

	var n, p int
	Fscan(in, &n, &p)
	p = p * inv[mx] % mod

	f := make([][]int, n+1)
	for i := range f {
		f[i] = make([]int, n+1)
	}
	for i := range f[0] {
		f[0][i] = 1
	}
	for i := 1; i <= n; i++ {
		for j := range n - i + 1 {
			for k := 1; k <= i; k++ {
				val := p * f[k-1][j+1] % mod
				if j > 0 {
					val = (val + (1-p+mod)%mod*f[k-1][j-1]%mod) % mod
				}
				f[i][j] = (f[i][j] + val*f[i-k][j]%mod*inv[k]%mod) % mod
			}
		}
	}

	ans := f[n][0]
	for i := 2; i <= n; i++ {
		ans = ans * i % mod * inv[i*2-1] % mod
	}
	Fprint(out, ans)
}

//func main() { cf1781F(os.Stdin, os.Stdout) }
