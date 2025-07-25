// Generated by copypasta/template/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/2101/C
// https://codeforces.com/problemset/status/2101/problem/C?friends=on
func Test_cf2101C(t *testing.T) {
	testCases := [][2]string{
		{
			`4
4
1 2 1 2
2
2 2
10
1 2 1 5 1 2 2 1 1 2
8
1 5 2 8 4 1 4 2`,
			`4
1
16
16`,
		},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, cf2101C)
}
