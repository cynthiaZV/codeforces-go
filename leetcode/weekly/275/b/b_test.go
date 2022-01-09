// Code generated by copypasta/template/leetcode/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
	"testing"
)

func Test(t *testing.T) {
	t.Log("Current test is [b]")
	examples := [][]string{
		{
			`[0,1,0,1,1,0,0]`, 
			`1`,
		},
		{
			`[0,1,1,1,0,0,1,1,0]`, 
			`2`,
		},
		{
			`[1,1,0,0,1]`, 
			`0`,
		},
		
	}
	targetCaseNum := 0 // -1
	if err := testutil.RunLeetCodeFuncWithExamples(t, minSwaps, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
// https://leetcode-cn.com/contest/weekly-contest-275/problems/minimum-swaps-to-group-all-1s-together-ii/
