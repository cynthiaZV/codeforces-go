// Generated by copypasta/template/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1862/E
// https://codeforces.com/problemset/status/1862/problem/E
func Test_cf1862E(t *testing.T) {
	testCases := [][2]string{
		{
			`6
5 2 2
3 2 5 4 6
4 3 2
1 1 1 1
6 6 6
-82 45 1 -77 39 11
5 2 2
3 2 5 4 8
2 1 1
-1 2
6 3 2
-8 8 -2 -1 9 0`,
			`2
0
60
3
0
7`,
		},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, cf1862E)
}
