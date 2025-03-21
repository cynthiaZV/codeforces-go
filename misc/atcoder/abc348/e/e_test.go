// Code generated by copypasta/template/atcoder/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// 题目：https://atcoder.jp/contests/abc348/tasks/abc348_e
// 提交：https://atcoder.jp/contests/abc348/submit?taskScreenName=abc348_e
// 对拍：https://atcoder.jp/contests/abc348/submissions?f.LanguageName=Go&f.Status=AC&f.Task=abc348_e&orderBy=source_length
// 最短：https://atcoder.jp/contests/abc348/submissions?f.Status=AC&f.Task=abc348_e&orderBy=source_length
func Test_e(t *testing.T) {
	testCases := [][2]string{
		{
			`4
1 2
1 3
2 4
1 1 1 2`,
			`5`,
		},
		{
			`2
2 1
1 1000000000`,
			`1`,
		},
		{
			`7
7 3
2 5
2 4
3 1
3 6
2 1
2 7 6 9 3 4 6`,
			`56`,
		},
		
	}
	testutil.AssertEqualStringCase(t, testCases, 0, run)
}
