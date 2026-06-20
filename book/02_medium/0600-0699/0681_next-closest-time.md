# 0681 — Next Closest Time

## Deskripsi

**Soal:** [0681. Next Closest Time](https://leetcode.com/problems/next-closest-time/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(1) since there are at most 4^4 = 256 combinations  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #681: Next Closest Time
// https://leetcode.com/problems/next-closest-time/
// Difficulty: Medium [Paid]
// Time: O(1) since there are at most 4^4 = 256 combinations
// Space: O(1)

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(nextClosestTime("19:34"))
	fmt.Println(nextClosestTime("23:59"))
	fmt.Println(nextClosestTime("13:55"))
}

func nextClosestTime(time string) string {
  // Membuat map untuk pencarian O(1): key → value
	digits := make(map[byte]bool)
  // Loop standar: indeks 0 sampai n-1
	for i := 0; i < len(time); i++ {
		if time[i] != ':' {
			digits[time[i]] = true
		}
	}

	hours, _ := strconv.Atoi(time[:2])
	minutes, _ := strconv.Atoi(time[3:])

	current := hours*60 + minutes

	for elapsed := 1; elapsed <= 24*60; elapsed++ {
		t := (current + elapsed) % (24 * 60)
		h := t / 60
		m := t % 60

		hs := fmt.Sprintf("%02d%02d", h, m)
		valid := true
		for i := 0; i < 4; i++ {
			if !digits[hs[i]] {
				valid = false
				break
			}
		}

		if valid {
			return fmt.Sprintf("%02d:%02d", h, m)
		}
	}

	return ""
}
```
