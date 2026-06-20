# 0383 — Ransom Note

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func RansomNote(ransomNote, magazine string) bool`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n+m), Space: O(1)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #383: Ransom Note
// https://leetcode.com/problems/ransom-note/
// Difficulty: Easy

import "fmt"

// Time: O(n+m), Space: O(1)
func RansomNote(ransomNote, magazine string) bool {
	count := [26]int{}
	for _, c := range magazine {
		count[c-'a']++
	}
	for _, c := range ransomNote {
		count[c-'a']--
		if count[c-'a'] < 0 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(RansomNote("a", "b"))
	fmt.Println(RansomNote("aa", "ab"))
	fmt.Println(RansomNote("aa", "aab"))
}
```
