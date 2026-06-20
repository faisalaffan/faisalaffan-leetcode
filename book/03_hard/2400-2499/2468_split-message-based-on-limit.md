# 2468 — Split Message Based On Limit

## Deskripsi

**Soal:** [2468. Split Message Based On Limit](https://leetcode.com/problems/split-message-based-on-limit/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** Binary Search (pencarian biner)

**Fungsi Solusi:** `func splitMessage(message string, limit int) []string`

> **Ide Kunci:** Binary search on the number of parts. For a given part count p,

## Solusi Go

```go
package main

// LeetCode #2468: Split Message Based on Limit
// https://leetcode.com/problems/split-message-based-on-limit/
// Difficulty: Hard
//
// Given a message string and a limit per part, split the message into parts
// where each part has a suffix " <i/total>" appended. The suffix length counts
// towards the limit. Find the minimum number of parts (and the actual split)
// such that each part (including suffix) is <= limit characters.
//
// Approach: Binary search on the number of parts. For a given part count p,
// compute the suffix length for each part: len("<i/p>"). Check if the total
// available characters (p * limit - total_suffix_length) >= message length.
// Then construct the parts.

import (
	"fmt"
	"strconv"
)

func splitMessage(message string, limit int) []string {
	n := len(message)

	// Try different part counts
	for parts := 1; parts <= n; parts++ {
		suffixLen := len("</>") // <, /, >
		suffixLen += len(strconv.Itoa(parts))

		// Check total capacity
		totalAvailable := parts * (limit - suffixLen)
		if totalAvailable < n {
			continue
		}

		// Check each part's capacity individually (templates vary)
		// Re-check with exact suffixes
		var totalCap int
		feasible := true
		msgIdx := 0

		for i := 1; i <= parts; i++ {
			tag := "<" + strconv.Itoa(i) + "/" + strconv.Itoa(parts) + ">"
			tagLen := len(tag)
			available := limit - tagLen
			if available <= 0 {
				feasible = false
				break
			}
			totalCap += available
		}

		if !feasible || totalCap < n {
			continue
		}

		// Construct the result
  // Membuat slice untuk menyimpan hasil
		result := make([]string, parts)
		msgIdx = 0
		for i := 1; i <= parts; i++ {
			tag := "<" + strconv.Itoa(i) + "/" + strconv.Itoa(parts) + ">"
			available := limit - len(tag)
			end := msgIdx + available
			if end > n {
				end = n
			}
			result[i-1] = message[msgIdx:end] + tag
			msgIdx = end
		}

		return result
	}

	return []string{}
}

func main() {
	// Example 1
	fmt.Println(splitMessage("this is really a very awesome message", 9))
	// Example 2
	fmt.Println(splitMessage("short message", 15))
	// Single part
	fmt.Println(splitMessage("hello", 10))
	// Edge: exact fit
	fmt.Println(splitMessage("abc", 5))
	// Longer message
	fmt.Println(splitMessage("the quick brown fox jumps over the lazy dog", 11))
}
```
