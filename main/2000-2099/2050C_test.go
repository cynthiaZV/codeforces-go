// Generated by copypasta/template/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/2050/C
// https://codeforces.com/problemset/status/2050/problem/C?friends=on
func Test_cf2050C(t *testing.T) {
	testCases := [][2]string{
		{
			`9
123
322
333333333333
9997
5472778912773
1234567890
23
33
52254522632`,
			`NO
YES
YES
NO
NO
YES
NO
YES
YES`,
		},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, cf2050C)
}
