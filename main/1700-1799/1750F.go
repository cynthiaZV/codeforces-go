package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1750F(in io.Reader, out io.Writer) {
	var n, m int
	Fscan(in, &n, &m)
	f := make([]int, n+1)
	g := make([]int, n+1)
	g[0] = 1
	h := make([]int, n+1)
	pow2 := 1
	for i := 1; i <= n; i++ {
		g[i] = g[i-1]
		for j := 1; j < i; j++ {
			if j*3 < i {
				g[i] = (g[i] + g[i-j*3-1]*f[j]) % m
			}
			h[i] = (h[i] + f[j]*f[i-j]) % m
			if j*2 < i {
				f[i] = (f[i] + h[j]*g[i-j*2-1]) % m
			}
		}
		if i > 2 {
			pow2 = pow2 * 2 % m
		}
		f[i] = (pow2 - f[i] + m) % m
	}
	Fprint(out, f[n])
}

//func main() { cf1750F(os.Stdin, os.Stdout) }
