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
		deg := make([]int, n)
		for range n - 1 {
			var v, w int
			Fscan(in, &v, &w)
			v--
			w--
			g[v] = append(g[v], w)
			g[w] = append(g[w], v)
			deg[v]++
			deg[w]++
		}

		dep := make([]int8, n)
		var dfs func(int, int)
		dfs = func(v, fa int) {
			for _, w := range g[v] {
				if w != fa {
					dep[w] = dep[v] ^ 1
					dfs(w, v)
				}
			}
		}
		dfs(n-1, -1)

		q := make([]int, 0, n)
		orders := q
		for i, d := range deg[:n-1] {
			if d == 1 {
				q = append(q, i)
			}
		}
		for len(q) > 0 {
			v := q[0]
			q = q[1:]
			for _, w := range g[v] {
				deg[w]--
				if deg[w] == 1 {
					q = append(q, w)
				}
			}
		}

		ans := []int{}
		cur := dep[0]
		for _, v := range orders[:n-1] {
			if dep[v] == cur {
				ans = append(ans, 0)
			}
			ans = append(ans, v+1, 0)
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
