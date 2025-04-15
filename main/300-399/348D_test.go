// Generated by copypasta/template/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/348/D
// https://codeforces.com/problemset/status/348/problem/D?friends=on
func Test_cf348D(t *testing.T) {
	testCases := [][2]string{
		{
			`4 5
.....
.###.
.###.
.....`,
			`1`,
		},
		{
			`2 3
...
...`,
			`1`,
		},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, cf348D)
}
