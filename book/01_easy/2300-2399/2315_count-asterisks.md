# 2315 — Count Asterisks

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func CountAsterisks(s string) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** —  |  **Ruang:** —


## 💻 Solusi Go

```go
package main

// LeetCode #2315: Count Asterisks
// https://leetcode.com/problems/count-asterisks/
// Difficulty: Easy
// Time O(n) | Space O(1)

import "fmt"

func main() {
	fmt.Println(CountAsterisks("l|*e*et|c**o|*de|")) // 2
	fmt.Println(CountAsterisks("iamprogrammer"))      // 0
}

func CountAsterisks(s string) int {
	count := 0
	bar := false
  // Linear scan O(n)
	for i := 0; i < len(s); i++ {
		if s[i] == '|' {
			bar = !bar
		} else if s[i] == '*' && !bar {
			count++
		}
	}
	return count
}
```
