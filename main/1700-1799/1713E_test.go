// Generated by copypasta/template/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1713/E
// https://codeforces.com/problemset/status/1713/problem/E?friends=on
func Test_cf1713E(t *testing.T) {
	testCases := [][2]string{
		{
			`2
3
2 1 2
2 1 2
1 1 2
4
3 3 1 2
1 1 3 1
3 2 3 2
2 3 3 1`,
			`2 1 1
2 1 1
2 2 2
3 1 1 2
3 1 2 1
3 3 3 3
2 3 2 1`,
		},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, cf1713E)
}
