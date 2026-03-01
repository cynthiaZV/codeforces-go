package main

// https://space.bilibili.com/206214
func mergeCharacters(s string, k int) string {
	last := [26]int{}
	for i := range last {
		last[i] = -k - 1 // 保证首次遇到字母 i 时，len(ans)-last[i] > k 是 true
	}

	ans := []byte{}
	for _, ch := range s {
		// ch 在 ans 中的下标是 len(ans)
		if len(ans)-last[ch-'a'] > k {
			last[ch-'a'] = len(ans)
			ans = append(ans, byte(ch))
		}
	}
	return string(ans)
}
