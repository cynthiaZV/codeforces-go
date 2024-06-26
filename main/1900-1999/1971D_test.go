// Generated by copypasta/template/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/1971/problem/D
// https://codeforces.com/problemset/status/1971/problem/D
func Test_cf1971D(t *testing.T) {
	testCases := [][2]string{
		{
			`6
11010
00000000
1
10
0001111
0110`,
			`3
1
1
2
1
2`,
		},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, cf1971D)
}
