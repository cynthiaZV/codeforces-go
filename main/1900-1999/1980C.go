package main

import (
	. "fmt"
	"io"
	"slices"
)

// https://github.com/EndlessCheng
func cf1980C(in io.Reader, out io.Writer) {
	var T, n, m int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}

		b := make([]int, n)
		need := map[int]int{}
		left := 0
		for i, v := range a {
			Fscan(in, &b[i])
			if v != b[i] {
				need[b[i]]++
				left++
			}
		}

		Fscan(in, &m)
		d := make([]int, m)
		for i := range d {
			Fscan(in, &d[i])
		}

		if !slices.Contains(b, d[m-1]) {
			Fprintln(out, "NO")
			continue
		}

		for i := m - 1; i >= 0; i-- {
			v := d[i]
			if need[v] > 0 {
				need[v]--
				left--
			}
		}
		if left == 0 {
			Fprintln(out, "YES")
		} else {
			Fprintln(out, "NO")
		}
	}
}

//func main() { cf1980C(bufio.NewReader(os.Stdin), os.Stdout) }
