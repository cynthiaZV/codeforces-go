package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1740F(in io.Reader, out io.Writer) {
	const mod = 998244353
	var n int
	Fscan(in, &n)
	cnt := make([]int, n+1)
	s := make([]int, n+1)
	for range n {
		var v int
		Fscan(in, &v)
		cnt[v]++
		s[cnt[v]]++
	}
	for i := 1; i <= n; i++ {
		s[i] += s[i-1]
	}

	f := make([][]int, n+1)
	for i := range f {
		f[i] = make([]int, n+1)
	}
	f[0][0] = 1
	for v := n; v > 0; v-- {
		for j := 1; j*v <= n; j++ {
			for k := v * j; k <= s[j]; k++ {
				f[j][k] = (f[j][k] + f[j-1][k-v]) % mod
			}
		}
	}

	ans := 0
	for i := 1; i <= n; i++ {
		ans += f[i][n]
	}
	Fprint(out, ans%mod)
}

//func main() { cf1740F(bufio.NewReader(os.Stdin), os.Stdout) }
