// Code generated by copypasta/template/atcoder/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// 题目：https://atcoder.jp/contests/arc075/tasks/arc075_c
// 提交：https://atcoder.jp/contests/arc075/submit?taskScreenName=arc075_c
// 对拍：https://atcoder.jp/contests/arc075/submissions?f.LanguageName=Go&f.Status=AC&f.Task=arc075_c&orderBy=source_length
// 最短：https://atcoder.jp/contests/arc075/submissions?f.Status=AC&f.Task=arc075_c&orderBy=source_length
func Test_c(t *testing.T) {
	testCases := [][2]string{
		{
			`3 6
7
5
7`,
			`5`,
		},
		{
			`1 2
1`,
			`0`,
		},
		{
			`7 26
10
20
30
40
30
20
10`,
			`13`,
		},
		
	}
	testutil.AssertEqualStringCase(t, testCases, 0, run)
}
