package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf2045F(in io.Reader, out io.Writer) {
	var T, n, m, k, r, c, a int
o:
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &m, &k)
		k++
		xors := make([]int, k)
		for range m {
			Fscan(in, &r, &c, &a)
			xors[r%k] ^= a % k
		}
		for _, v := range xors {
			if v > 0 {
				Fprintln(out, "Anda")
				continue o
			}
		}
		Fprintln(out, "Kamu")
	}
}

//func main() { cf2045F(bufio.NewReader(os.Stdin), os.Stdout) }
