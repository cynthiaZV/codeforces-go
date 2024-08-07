// Generated by copypasta/template/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1037/C
// https://codeforces.com/problemset/status/1037/problem/C
func Test_cf1037C(t *testing.T) {
	testCases := [][2]string{
		{
			`3
100
001`,
			`2`,
		},
		{
			`4
0101
0011`,
			`1`,
		},
		{
			`15
101010101010101
010101010101010`,
			`8`,
		},
		{
			`6
110110
000000`,
			`4`,
		},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, cf1037C)
}
