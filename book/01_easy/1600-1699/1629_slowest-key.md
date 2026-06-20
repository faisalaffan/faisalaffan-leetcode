# 1629 — Slowest Key

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func SlowestKey(releaseTimes []int, keysPressed string) byte`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n), Space: O(1)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #1629: Slowest Key
// https://leetcode.com/problems/slowest-key/
// Difficulty: Easy

import "fmt"

// Time: O(n), Space: O(1)
func SlowestKey(releaseTimes []int, keysPressed string) byte {
	maxDuration := releaseTimes[0]
	result := keysPressed[0]
	for i := 1; i < len(releaseTimes); i++ {
		duration := releaseTimes[i] - releaseTimes[i-1]
		if duration > maxDuration || (duration == maxDuration && keysPressed[i] > result) {
			maxDuration = duration
			result = keysPressed[i]
		}
	}
	return result
}

func main() {
	fmt.Printf("%c\n", SlowestKey([]int{9, 29, 49, 50}, "cbcd"))
	fmt.Printf("%c\n", SlowestKey([]int{12, 23, 36, 46, 62}, "spuda"))
}
```
