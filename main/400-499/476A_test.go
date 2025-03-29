// Generated by copypasta/template/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/476/A
// https://codeforces.com/problemset/status/476/problem/A?friends=on
func Test_cf476A(t *testing.T) {
	testCases := [][2]string{
		{
			`10 2`,
			`6`,
		},
		{
			`3 5`,
			`-1`,
		},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, cf476A)
}
