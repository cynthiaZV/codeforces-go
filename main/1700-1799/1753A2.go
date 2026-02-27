package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1753A2(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]int, n)
		s := 0
		for i := range a {
			Fscan(in, &a[i])
			s += a[i]
		}

		tar := 1
		if s < 0 {
			tar = -1
		}

		k := n
		flip := make([]bool, n)
		for i := 1; i < n; {
			if a[i] != tar || s == 0 {
				i++
				continue
			}
			st := i
			for i < n && a[i] == tar {
				if s != 0 && (i-st)%2 == 0 {
					flip[i] = true
					s -= tar * 2
					k--
				}
				i++
			}
		}

		if s != 0 {
			Fprintln(out, -1)
			continue
		}
		Fprintln(out, k)
		for i := 0; i < n; i++ {
			if i < n-1 && flip[i+1] {
				Fprintln(out, i+1, i+2)
				i++
			} else {
				Fprintln(out, i+1, i+1)
			}
		}
	}
}

//func main() { cf1753A2(bufio.NewReader(os.Stdin), os.Stdout) }
