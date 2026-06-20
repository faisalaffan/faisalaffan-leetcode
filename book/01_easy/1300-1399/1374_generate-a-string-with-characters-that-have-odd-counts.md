# 1374 — Generate A String With Characters That Have Odd Counts

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func generateTheString(n int) string`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n), Space: O(n)  |  **Ruang:** O(n)


## 💻 Solusi Go

```go
package main

// LeetCode #1374: Generate a String With Characters That Have Odd Counts
// https://leetcode.com/problems/generate-a-string-with-characters-that-have-odd-counts/
// Difficulty: Easy
//
// LeetCode submission: func generateTheString(n int) string

import "fmt"

func main() {
	fmt.Println(GenerateAStringWithCharactersThatHaveOddCounts(4)) // "aaab"
	fmt.Println(GenerateAStringWithCharactersThatHaveOddCounts(2)) // "ab"
	fmt.Println(GenerateAStringWithCharactersThatHaveOddCounts(7)) // "aaaaaaa"
}

// Time: O(n), Space: O(n)
func GenerateAStringWithCharactersThatHaveOddCounts(n int) string {
	if n%2 == 1 {
		return string(makeN('a', n))
	}
	return string(makeN('a', n-1)) + "b"
}

func makeN(ch byte, n int) []byte {
	res := make([]byte, n)
  // Range loop
	for i := range res {
		res[i] = ch
	}
	return res
}
```
