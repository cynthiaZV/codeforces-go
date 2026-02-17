package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1739B(in io.Reader, out io.Writer) {
	var T, n int
o:
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}
		for i := 1; i < n; i++ {
			if a[i] > 0 && a[i] <= a[i-1] {
				Fprintln(out, -1)
				continue o
			}
			a[i] += a[i-1]
		}
		for _, v := range a {
			Fprint(out, v, " ")
		}
		Fprintln(out)
	}
}

//func main() { cf1739B(bufio.NewReader(os.Stdin), os.Stdout) }
