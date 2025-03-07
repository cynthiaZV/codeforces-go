// Generated by copypasta/template/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/2029/C
// https://codeforces.com/problemset/status/2029/problem/C?friends=on
func Test_cf2029C(t *testing.T) {
	testCases := [][2]string{
		{
			`5
6
1 2 3 4 5 6
7
1 2 1 1 1 3 4
1
1
9
9 9 8 2 4 4 3 5 3
10
1 2 3 4 1 3 2 1 1 10`,
			`5
4
0
4
5`,
		},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, cf2029C)
}
