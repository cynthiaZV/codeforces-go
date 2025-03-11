// Code generated by copypasta/template/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://www.luogu.com.cn/problem/P3373
func Test_p3373(t *testing.T) {
	testCases := [][2]string{
		{
			`5 5 38
1 5 4 2 3
2 1 4 1
3 2 5
1 2 4 2
2 3 5 5
3 1 4`,
			`17
2`,
		},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, p3373)
}
