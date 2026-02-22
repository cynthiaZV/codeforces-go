package main

// https://space.bilibili.com/206214
func scoreDifference(nums []int) int {
	score := [2]int{}
	active := 0 // 主动玩家一开始是第一位玩家
	for i, x := range nums {
		active ^= x % 2 // 如果 x 是奇数，主动玩家换人
		if i%6 == 5 {
			active ^= 1 // 主动玩家换人
		}
		score[active] += x
	}
	return score[0] - score[1]
}
