package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf2154D(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		g := make([][]int, n)
		for range n - 1 {
			var v, w int
			Fscan(in, &v, &w)
			v--
			w--
			g[v] = append(g[v], w)
			g[w] = append(g[w], v)
		}

		dep := make([]int8, n)
		for i := range dep {
			dep[i] = -1
		}
		dep[n-1] = 0
		q := make([]int, 1, n)
		q[0] = n - 1
		orders := q
		for len(q) > 0 {
			v := q[0]
			q = q[1:]
			for _, w := range g[v] {
				if dep[w] < 0 {
					dep[w] = dep[v] ^ 1
					q = append(q, w)
				}
			}
		}

		ans := []int{}
		cur := dep[0]
		orders = orders[1:n]
		for i := len(orders) - 1; i >= 0; i-- {
			v := orders[i]
			if dep[v] == cur {
				ans = append(ans, 0)
				cur ^= 1
			}
			ans = append(ans, v+1, 0)
			cur ^= 1
		}

		Fprintln(out, len(ans))
		for _, v := range ans {
			if v > 0 {
				Fprintln(out, 2, v)
			} else {
				Fprintln(out, 1)
			}
		}
	}
}

//func main() { cf2154D(bufio.NewReader(os.Stdin), os.Stdout) }
