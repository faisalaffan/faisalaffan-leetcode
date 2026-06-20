# 1358 — Number Of Substrings Containing All Three Characters

## Deskripsi

**Soal:** [1358. Number Of Substrings Containing All Three Characters](https://leetcode.com/problems/number-of-substrings-containing-all-three-characters/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n) where n = length of string  
**Kompleksitas Ruang:** O(1) - fixed array of 3

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #1358: Number of Substrings Containing All Three Characters
// https://leetcode.com/problems/number-of-substrings-containing-all-three-characters/
// Difficulty: Medium

import "fmt"

func main() {
	// Test case 1
	fmt.Println(numberOfSubstrings("abcabc")) // 10

	// Test case 2
	fmt.Println(numberOfSubstrings("aaacb")) // 3

	// Test case 3
	fmt.Println(numberOfSubstrings("abc")) // 1
}

// Time: O(n) where n = length of string
// Space: O(1) - fixed array of 3
func numberOfSubstrings(s string) int {
	count := 0
	lastPos := [3]int{-1, -1, -1}

  // Loop standar: indeks 0 sampai n-1
	for i := 0; i < len(s); i++ {
		lastPos[s[i]-'a'] = i
		// Find the minimum of the last positions of 'a', 'b', 'c'
		// This is the leftmost boundary of the substring containing all three
		minPos := lastPos[0]
		if lastPos[1] < minPos {
			minPos = lastPos[1]
		}
		if lastPos[2] < minPos {
			minPos = lastPos[2]
		}
		// If all three have been seen, count substrings ending at i
		if minPos >= 0 {
			count += minPos + 1
		}
	}

	return count
}
```
