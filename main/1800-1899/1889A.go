package main

import (
	. "fmt"
	"io"
	"strings"
)

// https://github.com/EndlessCheng
func cf1889A(in io.Reader, out io.Writer) {
	var T, n int
	var s string
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &s)
		if strings.Count(s, "0")*2 != n {
			Fprintln(out, -1)
			continue
		}
		ans := []any{}
		i, j := 0, n-1
		for s != "" {
			if s[0] != s[len(s)-1] {
				s = s[1 : len(s)-1]
				j--
			} else if s[0] == '0' {
				s = s[1:] + "0"
				j++
				ans = append(ans, j)
			} else {
				s = "1" + s[:len(s)-1]
				j++
				ans = append(ans, i)
			}
			i++
		}
		Fprintln(out, len(ans))
		Fprintln(out, ans...)
	}
}

//func main() { cf1889A(os.Stdin, os.Stdout) }
