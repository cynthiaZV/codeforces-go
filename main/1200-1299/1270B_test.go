// Generated by copypasta/template/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1270/B
// https://codeforces.com/problemset/status/1270/problem/B?friends=on
func Test_cf1270B(t *testing.T) {
	testCases := [][2]string{
		{
			`3
5
1 2 3 4 5
4
2 0 1 9
2
2019 2020`,
			`NO
YES
1 4
NO`,
		},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, cf1270B)
}
