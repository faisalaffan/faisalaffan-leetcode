# 0383 — Ransom Note

## Deskripsi

**Soal:** [0383. Ransom Note](https://leetcode.com/problems/ransom-note/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n+m), Space: O(1)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

**Fungsi Solusi:** `func RansomNote(ransomNote, magazine string) bool`

## Solusi Go

```go
package main

// LeetCode #383: Ransom Note
// https://leetcode.com/problems/ransom-note/
// Difficulty: Easy

import "fmt"

// Time: O(n+m), Space: O(1)
func RansomNote(ransomNote, magazine string) bool {
	count := [26]int{}
	for _, c := range magazine {
		count[c-'a']++
	}
	for _, c := range ransomNote {
		count[c-'a']--
		if count[c-'a'] < 0 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(RansomNote("a", "b"))
	fmt.Println(RansomNote("aa", "ab"))
	fmt.Println(RansomNote("aa", "aab"))
}
```
