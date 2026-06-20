# 3168 — Minimum Number Of Chairs In A Waiting Room

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func MinimumNumberOfChairsInAWaitingRoom(s string) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #3168: Minimum Number of Chairs in a Waiting Room
// https://leetcode.com/problems/minimum-number-of-chairs-in-a-waiting-room/
// Difficulty: Easy

import "fmt"

func main() {
	// LeetCode name: minimumChairs
	fmt.Println(MinimumNumberOfChairsInAWaitingRoom("EEEE"))   // 4
	fmt.Println(MinimumNumberOfChairsInAWaitingRoom("ELELEL")) // 1
}

// Time: O(n) | Space: O(1)
// LeetCode submission name: minimumChairs
func MinimumNumberOfChairsInAWaitingRoom(s string) int {
	current := 0
	maxChairs := 0
  // Linear scan O(n)
	for i := 0; i < len(s); i++ {
		if s[i] == 'E' {
			current++
			if current > maxChairs {
				maxChairs = current
			}
		} else {
			current--
		}
	}
	return maxChairs
}
```
