package main

import (
	. "fmt"
	"io"
	"math/bits"
)

// https://github.com/EndlessCheng
func cf2027E1(in io.Reader, out io.Writer) {
	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}
		x := make([]int, n)
		for i := range x {
			Fscan(in, &x[i])
		}

		sg := 0
	o:
		for i, v := range a {
			m := 1 << bits.Len(uint(v))
			w := x[i] % m
			for j := v + 1; j < m; j += j & -j {
				if j == w {
					continue o
				}
			}

			res := bits.OnesCount(uint(w))
			d := w - m/2
			if d >= 0 && v < m/2+d&-d && bits.OnesCount(uint(d))%2 > 0 {
				res -= 2
			}
			sg ^= res
		}

		if sg == 0 {
			Fprintln(out, "Bob")
		} else {
			Fprintln(out, "Alice")
		}
	}
}

//func main() { cf2027E1(bufio.NewReader(os.Stdin), os.Stdout) }
