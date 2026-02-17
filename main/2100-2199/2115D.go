package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf2115D(in io.Reader, out io.Writer) {
	var T, n int
	var s string
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}
		b := make([]int, n)
		ans := 0
		for i := range b {
			Fscan(in, &b[i])
			ans ^= b[i]
			a[i] ^= b[i]
		}

		Fscan(in, &s)
		for i := 60; i >= 0; i-- {
			x, y := 0, 0
			for j := n - 1; j >= 0; j-- {
				if a[j]>>i&1 == 0 {
					continue
				}
				if x > 0 {
					a[j] ^= x
				} else {
					y = int(s[j] - '0')
					x = a[j]
					a[j] ^= x
				}
			}
			ans ^= (ans>>i&1 ^ y) * x
		}
		Fprintln(out, ans)
	}
}

//func main() { cf2115D(bufio.NewReader(os.Stdin), os.Stdout) }
