// Generated by copypasta/template/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1042/D
// https://codeforces.com/problemset/status/1042/problem/D
func Test_cf1042D(t *testing.T) {
	testCases := [][2]string{
		{
			`5 4
5 -1 3 4 -1`,
			`5`,
		},
		{
			`3 0
-1 2 -3`,
			`4`,
		},
		{
			`4 -1
-2 1 -2 3`,
			`3`,
		},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, cf1042D)
}