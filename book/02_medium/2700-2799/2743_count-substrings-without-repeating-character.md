# 2743 — Count Substrings Without Repeating Character

## Deskripsi

**Soal:** [2743. Count Substrings Without Repeating Character](https://leetcode.com/problems/count-substrings-without-repeating-character/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

**Fungsi Solusi:** `func CountSubstringsWithoutRepeatingCharacter(s string) int`

## Solusi Go

```go
package main

// LeetCode #2743: Count Substrings Without Repeating Character
// https://leetcode.com/problems/count-substrings-without-repeating-character/
// Difficulty: Medium [Paid]
// Time: O(n) | Space: O(1)

import "fmt"

func CountSubstringsWithoutRepeatingCharacter(s string) int {
  // Membuat map untuk pencarian O(1): key → value
	lastPos := make(map[byte]int)
	left := 0
	count := 0

	for right := 0; right < len(s); right++ {
		if pos, ok := lastPos[s[right]]; ok && pos >= left {
			left = pos + 1
		}
		lastPos[s[right]] = right
		count += right - left + 1
	}

	return count
}

func main() {
	fmt.Println(CountSubstringsWithoutRepeatingCharacter("abcabc"))
	fmt.Println(CountSubstringsWithoutRepeatingCharacter("aaaa"))
	fmt.Println(CountSubstringsWithoutRepeatingCharacter(""))
}
```
