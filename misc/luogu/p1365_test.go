// Code generated by copypasta/template/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://www.luogu.com.cn/problem/P1365
func Test_p1365(t *testing.T) {
	testCases := [][2]string{
		{
			`4
????`,
			`4.1250`,
		},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, p1365)
}
