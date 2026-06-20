# 1688 — Count Of Matches In Tournament

## Deskripsi

**Soal:** [1688. Count Of Matches In Tournament](https://leetcode.com/problems/count-of-matches-in-tournament/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(log n), Space: O(1)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

**Fungsi Solusi:** `func NumberOfMatches(n int) int`

## Solusi Go

```go
package main

// LeetCode #1688: Count of Matches in Tournament
// https://leetcode.com/problems/count-of-matches-in-tournament/
// Difficulty: Easy

import "fmt"

// Time: O(log n), Space: O(1)
func NumberOfMatches(n int) int {
	matches := 0
	for n > 1 {
		if n%2 == 0 {
			matches += n / 2
			n /= 2
		} else {
			matches += (n - 1) / 2
			n = (n-1)/2 + 1
		}
	}
	return matches
}

// Alternative O(1): return n-1 (each match eliminates one team, champion is last)

func main() {
	fmt.Println(NumberOfMatches(7))
	fmt.Println(NumberOfMatches(14))
}
```
