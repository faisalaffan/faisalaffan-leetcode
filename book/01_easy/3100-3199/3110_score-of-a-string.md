# 3110 — Score Of A String

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func ScoreOfAString(s string) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #3110: Score of a String
// https://leetcode.com/problems/score-of-a-string/
// Difficulty: Easy

import "fmt"

func main() {
	// LeetCode name: scoreOfString
	fmt.Println(ScoreOfAString("hello")) // 13
	fmt.Println(ScoreOfAString("zaz"))   // 50
}

// Time: O(n) | Space: O(1)
// LeetCode submission name: scoreOfString
func ScoreOfAString(s string) int {
	score := 0
	for i := 1; i < len(s); i++ {
		diff := int(s[i]) - int(s[i-1])
		if diff < 0 {
			diff = -diff
		}
		score += diff
	}
	return score
}
```
