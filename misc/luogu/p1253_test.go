// Code generated by copypasta/template/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://www.luogu.com.cn/problem/P1253
func Test_p1253(t *testing.T) {
	testCases := [][2]string{
		{
			`6 6
1 1 4 5 1 4
1 1 2 6
2 3 4 2
3 1 4
3 2 3
1 1 6 -1
3 1 6`,
			`7
6
-1`,
		},
		{
			`4 4
10 4 -3 -7
1 1 3 0
2 3 4 -4
1 2 4 -9
3 1 4`,
			`0`,
		},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, p1253)
}
