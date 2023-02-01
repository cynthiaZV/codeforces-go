// Code generated by copypasta/template/atcoder/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// 提交：https://atcoder.jp/contests/arc148/submit?taskScreenName=arc148_c
// 对拍：https://atcoder.jp/contests/arc148/submissions?f.LanguageName=Go&f.Status=AC&f.Task=arc148_c&orderBy=source_length
func Test_run(t *testing.T) {
	t.Log("Current test is [c]")
	testCases := [][2]string{
		{
			`6 6
1 1 2 2 5
6 1 2 3 4 5 6
3 2 5 6
1 3
3 1 2 3
3 4 5 6
4 2 3 4 5`,
			`1
2
1
3
2
3`,
		},
		
	}
	testutil.AssertEqualStringCase(t, testCases, 0, run)
}
// https://atcoder.jp/contests/arc148/tasks/arc148_c
