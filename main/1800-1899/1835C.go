package main

import (
	"bufio"
	. "fmt"
	"io"
	"slices"
)

// https://github.com/EndlessCheng
func cf1835C(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var T, k int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &k)
		n := 1 << (k + 1)
		s := make([]int, n+1)
		for i := 1; i <= n; i++ {
			Fscan(in, &s[i])
			s[i] ^= s[i-1]
		}

		half := n / 2
		last := make([]int, half)
		for i := range last {
			last[i] = -1
		}
		last[0] = 0
		pos := make([][2]int, half)

		for i := 1; ; i++ {
			low := s[i] & (half - 1)
			j := last[low]
			if j >= 0 {
				high := (s[i] ^ s[j]) >> k
				if pos[high][1] > 0 {
					p := append(pos[high][:], j)
					slices.Sort(p)
					Fprintln(out, p[0]+1, p[1], p[2]+1, i)
					break
				}
				pos[high] = [2]int{j, i}
			}
			last[low] = i
		}
	}
}

//func main() { cf1835C(bufio.NewReader(os.Stdin), os.Stdout) }
