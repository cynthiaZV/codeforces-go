// Generated by copypasta/template/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/2062/D
// https://codeforces.com/problemset/status/2062/problem/D?friends=on
func Test_cf2062D(t *testing.T) {
	testCases := [][2]string{
		{
			`6
4
0 11
6 6
0 0
5 5
2 1
3 1
4 3
7
1 1
0 5
0 5
2 2
2 2
2 2
2 2
1 2
1 3
2 4
2 5
3 6
3 7
4
1 1
1 1
1 1
0 0
1 4
2 4
3 4
7
0 20
0 20
0 20
0 20
3 3
4 4
5 5
1 2
1 3
1 4
2 5
3 6
4 7
5
1000000000 1000000000
0 0
1000000000 1000000000
0 0
1000000000 1000000000
3 2
2 1
1 4
4 5
6
21 88
57 81
98 99
61 76
15 50
23 67
2 1
3 2
4 3
5 3
6 4`,
			`11
3
3
5
3000000000
98`,
		},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, cf2062D)
}
