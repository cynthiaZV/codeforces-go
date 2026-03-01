package main

import "strings"

// https://space.bilibili.com/206214
func trimTrailingVowels(s string) string {
	return strings.TrimRight(s, "aeiou")
}
