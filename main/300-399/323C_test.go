// Generated by copypasta/template/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/323/C
// https://codeforces.com/problemset/status/323/problem/C
func Test_cf323C(t *testing.T) {
	testCases := [][2]string{
		{
			`3
3 1 2
3 2 1
1
1 2 3 3`,
			`1`,
		},
		{
			`4
4 3 2 1
2 3 4 1
3
1 2 3 4
1 3 2 1
1 4 2 3`,
			`1
1
2`,
		},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, cf323C)
}
