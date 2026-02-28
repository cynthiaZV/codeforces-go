package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf733E(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n int
	var s string
	Fscan(in, &n, &s)

	var q1, q2 []int
	for i, ch := range s {
		if ch == 'D' {
			q1 = append(q1, i)
		} else {
			q2 = append(q2, n-1-i)
		}
	}

	sum := 0
	for i, v := range q1 {
		sum += (v-i)*2 + 1
		Fprint(out, sum, " ")
	}

	ans := make([]int, len(q2))
	sum = 0
	for i := range ans {
		sum += (q2[len(q2)-1-i]-i)*2 + 1
		ans[i] = sum
	}
	for i := len(ans) - 1; i >= 0; i-- {
		Fprint(out, ans[i], " ")
	}
}

//func main() { cf733E(bufio.NewReader(os.Stdin), os.Stdout) }
