// Generated by copypasta/template/generator_test.go
package main

import (
	"fmt"
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"io"
	"testing"
)

// https://codeforces.com/problemset/problem/1422/E
// https://codeforces.com/problemset/status/1422/problem/E?friends=on
func Test_cf1422E(t *testing.T) {
	testCases := [][2]string{
		{
			`abcdd`,
			`3 abc
2 bc
1 c
0 
1 d`,
		},
		{
			`abbcdddeaaffdfouurtytwoo`,
			`18 abbcd...tw
17 bbcdd...tw
16 bcddd...tw
15 cddde...tw
14 dddea...tw
13 ddeaa...tw
12 deaad...tw
11 eaadf...tw
10 aadfortytw
9 adfortytw
8 dfortytw
9 fdfortytw
8 dfortytw
7 fortytw
6 ortytw
5 rtytw
6 urtytw
5 rtytw
4 tytw
3 ytw
2 tw
1 w
0 
1 o`,
		},
		{
			`adda`,
			`2 aa
1 a
2 da
1 a`,
		},
		{
			`aabba`,
			`1 a
2 aa
1 a
2 ba
1 a`,
		},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, cf1422E)
}

func TestCompare_cf1422E(_t *testing.T) {
	return
	testutil.DebugTLE = 0
	rg := testutil.NewRandGenerator()
	inputGenerator := func() string {
		//return ``
		rg.Clear()
		rg.Str(1, 20, 'a', 'e')
		return rg.String()
	}

	// 暴力算法
	runBF := func(in io.Reader, out io.Writer) {
		var s string
		fmt.Fscan(in, &s)

		n := len(s)
		el, ke := -1, -1
		sz := []struct {
			first, second int
		}{{-1, n}}
		sol := make([][]any, 0)
		var p string

		add := func(a int) {
			p += string(s[sz[a].second])
		}

		for i := n - 1; i >= 0; i-- {
			x := int(s[i] - 'a')
			y := sz[len(sz)-1].first
			t := sz[len(sz)-1].second
			if x == y && i+1 == t && (x > el || (x == el && x > ke)) {
				sz = sz[:len(sz)-1]
			} else {
				if x != y && i+1 == t && el != y {
					ke = el
					el = y
				}
				sz = append(sz, struct{ first, second int }{x, i})
			}

			si := len(sz) - 1
			p = ""
			if si > 10 {
				for j := si; j > si-5; j-- {
					add(j)
				}
				p += "..."
				add(2)
				add(1)
			} else {
				for j := si; j >= 1; j-- {
					add(j)
				}
			}
			sol = append(sol, []any{si, p})
		}

		for i := len(sol) - 1; i >= 0; i-- {
			v := sol[i]
			fmt.Fprintln(out, v...)
		}
	}
	
	testutil.AssertEqualRunResultsInf(_t, inputGenerator, runBF, cf1422E)
}
