// Code generated by copypasta/template/leetcode/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
	"testing"
)

func Test(t *testing.T) {
	examples := [][]string{
		{
			`11`, `2`, `4`,
			`[3,4]`,
		},
		{
			`5`, `1`, `5`,
			`[1,1]`,
		},
		{
			`5`, `1`, `4`,
			`[2,2]`,
		},
		{
			`14`, `2`, `5`,
			`[3,4]`,
		},
		{
			`8`, `3`, `7`,
			`[3,3]`,
		},
	}
	if err := testutil.RunLeetCodeFuncWithExamples(t, earliestAndLatest, examples, 0); err != nil {
		t.Fatal(err)
	}
}
// https://leetcode-cn.com/contest/weekly-contest-245/problems/the-earliest-and-latest-rounds-where-players-compete/
