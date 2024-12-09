// Generated by copypasta/template/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/2047/problem/A
// https://codeforces.com/problemset/status/2047/problem/A?friends=on
func Test_cf2047A(t *testing.T) {
	testCases := [][2]string{
		{
			`5
1
1
2
1 8
5
1 3 2 1 2
7
1 2 1 10 2 7 2
14
1 10 10 100 1 1 10 1 10 2 10 2 10 1`,
			`1
2
2
2
3`,
		},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, cf2047A)
}
