// Generated by copypasta/template/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1006/C
// https://codeforces.com/problemset/status/1006/problem/C?friends=on
func Test_cf1006C(t *testing.T) {
	testCases := [][2]string{
		{
			`5
1 3 1 1 4`,
			`5`,
		},
		{
			`5
1 3 2 1 4`,
			`4`,
		},
		{
			`3
4 1 2`,
			`0`,
		},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, cf1006C)
}
