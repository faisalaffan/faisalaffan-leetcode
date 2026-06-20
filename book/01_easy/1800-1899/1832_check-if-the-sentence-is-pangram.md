# 1832 — Check If The Sentence Is Pangram

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func CheckIfPangram(sentence string) bool`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n), Space: O(1)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #1832: Check if the Sentence Is Pangram
// https://leetcode.com/problems/check-if-the-sentence-is-pangram/
// Difficulty: Easy

import "fmt"

// Time: O(n), Space: O(1)
func CheckIfPangram(sentence string) bool {
	seen := 0
  // Linear scan O(n)
	for i := 0; i < len(sentence); i++ {
		seen |= 1 << (sentence[i] - 'a')
	}
	return seen == (1<<26)-1
}

func main() {
	fmt.Println(CheckIfPangram("thequickbrownfoxjumpsoverthelazydog"))
	fmt.Println(CheckIfPangram("leetcode"))
}
```
