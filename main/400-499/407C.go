package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf407C(in io.Reader, _w io.Writer) {
	const mod = 1_000_000_007
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n, m, l, r, k int
	Fscan(in, &n, &m)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}

	c := make([][101]int, n+101)
	for i := range c {
		c[i][0] = 1
		for j := 1; j <= min(i, 100); j++ {
			c[i][j] = (c[i-1][j] + c[i-1][j-1]) % mod
		}
	}

	f := [102][]int{}
	for i := range f {
		f[i] = make([]int, n+2)
	}
	for range m {
		Fscan(in, &l, &r, &k)
		f[k][l]++
		for j := range k + 1 {
			f[j][r+1] = (f[j][r+1] - c[r-l+k-j][k-j]) % mod
		}
	}
	for i := 100; i >= 0; i-- {
		for j := 1; j <= n; j++ {
			f[i][j] = (f[i][j] + f[i][j-1] + f[i+1][j]) % mod
		}
	}

	for i, v := range a {
		Fprint(out, (f[0][i+1]+v+mod)%mod, " ")
	}
}

//func main() { cf407C(bufio.NewReader(os.Stdin), os.Stdout) }
