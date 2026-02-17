package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1995C(in io.Reader, out io.Writer) {
	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}

		ans := 0
		k := 0
		for i := 1; i < n; i++ {
			x, y := a[i-1], a[i]
			if x > y {
				if y == 1 {
					ans = -1
					break
				}
				for y < x {
					y *= y
					k++
				}
				ans += k
			} else {
				for k > 0 && x*x <= y {
					x *= x
					k--
				}
				ans += k
			}
		}
		Fprintln(out, ans)
	}
}

//func main() { cf1995C(bufio.NewReader(os.Stdin), os.Stdout) }
