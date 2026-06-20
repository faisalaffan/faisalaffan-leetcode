# 3295 — Report Spam Message

## Deskripsi

**Soal:** [3295. Report Spam Message](https://leetcode.com/problems/report-spam-message/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n + m) Space: O(m)  
**Kompleksitas Ruang:** O(m)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #3295: Report Spam Message
// https://leetcode.com/problems/report-spam-message/
// Difficulty: Medium
// Time: O(n + m) Space: O(m)

import "fmt"

func main() {
	fmt.Println(reportSpam([]string{"hello", "world", "leetcode"}, []string{"world", "hello"})) // true
	fmt.Println(reportSpam([]string{"hello", "programming", "fun"}, []string{"world", "hello"})) // false
	fmt.Println(reportSpam([]string{"a", "b", "c", "d"}, []string{"a", "b", "x"}))               // true
}

func reportSpam(message []string, bannedWords []string) bool {
  // Membuat map untuk pencarian O(1): key → value
	banned := make(map[string]struct{}, len(bannedWords))
	for _, w := range bannedWords {
		banned[w] = struct{}{}
	}
	count := 0
	for _, w := range message {
		if _, ok := banned[w]; ok {
			count++
			if count >= 2 {
				return true
			}
		}
	}
	return false
}
```
