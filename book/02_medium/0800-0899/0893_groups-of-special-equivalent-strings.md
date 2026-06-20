# 0893 — Groups Of Special Equivalent Strings

## Deskripsi

**Soal:** [0893. Groups Of Special Equivalent Strings](https://leetcode.com/problems/groups-of-special-equivalent-strings/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n * m) where n = len(words), m = avg word length  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #893: Groups of Special-Equivalent Strings
// https://leetcode.com/problems/groups-of-special-equivalent-strings/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(GroupsOfSpecialEquivalentStrings([]string{"abcd", "cdab", "cbad", "xyzz", "zzxy", "zzyx"}))
	fmt.Println(GroupsOfSpecialEquivalentStrings([]string{"abc", "acb", "bac", "bca", "cab", "cba"}))
	fmt.Println(GroupsOfSpecialEquivalentStrings([]string{"a"}))
}

// Time: O(n * m) where n = len(words), m = avg word length | Space: O(n)
func GroupsOfSpecialEquivalentStrings(words []string) int {
  // Membuat map untuk pencarian O(1): key → value
	groups := make(map[[52]int]bool)

	for _, word := range words {
		var key [52]int
		for i, c := range word {
			// Even indices: 0-25, Odd indices: 26-51
			key[int(c-'a')+26*(i%2)]++
		}
		groups[key] = true
	}

	return len(groups)
}
```
