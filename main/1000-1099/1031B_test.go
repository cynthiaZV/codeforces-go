// Generated by copypasta/template/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1031/B
// https://codeforces.com/problemset/status/1031/problem/B?friends=on
func Test_cf1031B(t *testing.T) {
	testCases := [][2]string{
		{
			`4
3 3 2
1 2 0`,
			`YES
1 3 2 0`,
		},
		{
			`3
1 3
3 2`,
			`NO`,
		},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, cf1031B)
}
