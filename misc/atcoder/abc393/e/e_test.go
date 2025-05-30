// Code generated by copypasta/template/atcoder/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// 题目：https://atcoder.jp/contests/abc393/tasks/abc393_e
// 提交：https://atcoder.jp/contests/abc393/submit?taskScreenName=abc393_e
// 对拍：https://atcoder.jp/contests/abc393/submissions?f.LanguageName=Go&f.Status=AC&f.Task=abc393_e&orderBy=source_length
// 最短：https://atcoder.jp/contests/abc393/submissions?f.Status=AC&f.Task=abc393_e&orderBy=source_length
func Test_e(t *testing.T) {
	testCases := [][2]string{
		{
			`5 2
3 4 6 7 12`,
			`3
4
6
1
6`,
		},
		{
			`3 3
6 10 15`,
			`1
1
1`,
		},
		{
			`10 3
414003 854320 485570 52740 833292 625990 909680 885153 435420 221663`,
			`59
590
590
879
879
590
20
879
590
59`,
		},
		
	}
	testutil.AssertEqualStringCase(t, testCases, 0, run)
}
