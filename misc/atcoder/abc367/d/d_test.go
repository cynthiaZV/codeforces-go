// Code generated by copypasta/template/atcoder/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// 题目：https://atcoder.jp/contests/abc367/tasks/abc367_d
// 提交：https://atcoder.jp/contests/abc367/submit?taskScreenName=abc367_d
// 对拍：https://atcoder.jp/contests/abc367/submissions?f.LanguageName=Go&f.Status=AC&f.Task=abc367_d&orderBy=source_length
// 最短：https://atcoder.jp/contests/abc367/submissions?f.Status=AC&f.Task=abc367_d&orderBy=source_length
func Test_d(t *testing.T) {
	testCases := [][2]string{
		{
			`4 3
2 1 4 3`,
			`4`,
		},
		{
			`2 1000000
1 1`,
			`0`,
		},
		{
			`9 5
9 9 8 2 4 4 3 5 3`,
			`11`,
		},
		
	}
	testutil.AssertEqualStringCase(t, testCases, 0, run)
}
