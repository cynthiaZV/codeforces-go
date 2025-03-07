// Generated by copypasta/template/generator_test.go
package main

import (
	. "fmt"
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"io"
	"slices"
	"strings"
	"testing"
)

// https://codeforces.com/problemset/problem/1980/D
// https://codeforces.com/problemset/status/1980/problem/D?friends=on
func Test_cf1980D(t *testing.T) {
	testCases := [][2]string{
		{
			`12
6
20 6 12 3 48 36
4
12 6 3 4
3
10 12 3
5
32 16 8 4 2
5
100 50 2 10 20
4
2 4 8 1
10
7 4 6 2 4 5 1 4 2 8
7
5 9 6 8 5 9 2
6
11 14 8 12 9 3
9
5 7 3 10 6 3 12 6 3
3
4 2 4
8
1 6 11 12 6 12 3 6`,
			`YES
NO
YES
NO
YES
YES
NO
YES
YES
YES
YES
YES`,
		},
		{
			`1
4
6 2 3 5`,
			`YES`,
		},
		{
			`1
4
2 4 8 1`,
			`YES`,
		},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, cf1980D)
}

func TestCompare_cf1980D(_t *testing.T) {
	gcd := func(a, b int) int {
		for a != 0 {
			a, b = b%a, a
		}
		return b
	}
	//return
	testutil.DebugTLE = 0
	rg := testutil.NewRandGenerator()
	inputGenerator := func() string {
		//return ``
		rg.Clear()
		rg.One()
		n := rg.Int(3,9)
		rg.NewLine()
		rg.IntSlice(n, 1, 12)
		return rg.String()
	}

	// 暴力算法
	runBF := func(in io.Reader, out io.Writer) {
		solve := func(Case int) {
			var n int
			Fscan(in, &n)
			a := make([]int, n)
			for i := range a {
				Fscan(in, &a[i])
			}

			o:for i := range a {
				b := slices.Clone(a)
				b = append(b[:i], b[i+1:]...)
				pre := 0
				for i := 1; i < len(b); i++ {
					v, w := b[i-1], b[i]
					g := gcd(v,w)
					if g < pre {
						continue o
					}
					pre = g
				}
				Fprintln(out, "YES")
				return
			}
			Fprintln(out, "NO")
		}

		T := 1
		Fscan(in, &T)
		for Case := 1; Case <= T; Case++ {
			solve(Case)
		}

		_leftData, _ := io.ReadAll(in)
		if _s := strings.TrimSpace(string(_leftData)); _s != "" {
			panic("有未读入的数据：\n" + _s)
		}
	}

	testutil.AssertEqualRunResultsInf(_t, inputGenerator, runBF, cf1980D)
}
