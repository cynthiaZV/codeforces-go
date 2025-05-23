// Code generated by copypasta/template/atcoder/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// 题目：https://atcoder.jp/contests/arc051/tasks/arc051_c
// 提交：https://atcoder.jp/contests/arc051/submit?taskScreenName=arc051_c
// 对拍：https://atcoder.jp/contests/arc051/submissions?f.LanguageName=Go&f.Status=AC&f.Task=arc051_c&orderBy=source_length
// 最短：https://atcoder.jp/contests/arc051/submissions?f.Status=AC&f.Task=arc051_c&orderBy=source_length
func Test_c(t *testing.T) {
	testCases := [][2]string{
		{
			`3 10 3
1 99 10`,
			`99
100
100`,
		},
		{
			`2 100001 2
1 200000`,
			`200000
199931`,
		},
		{
			`10 123 1000000000
394632992 714094234 84420 5 3439891 3395 35 58 5001 730`,
			`954804718
385989482
948741792
268211139
100694402
492858064
955116743
68100851
154525400
479209143`,
		},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, run)
}
