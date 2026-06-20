# 3295 — Report Spam Message

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func reportSpam(message []string, bannedWords []string) bool`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap

**Waktu:** O(n + m) Space: O(m)  |  **Ruang:** O(m)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3295: Report Spam Message
// https://leetcode.com/problems/report-spam-message/
// Difficulty: Medium
// Time: O(n + m) Space: O(m)

import "fmt"

func main() {
	fmt.Println(reportSpam([]string{"hello", "world", "leetcode"}, []string{"world", "hello"})) // true
	fmt.Println(reportSpam([]string{"hello", "programming", "fun"}, []string{"world", "hello"})) // false
	fmt.Println(reportSpam([]string{"a", "b", "c", "d"}, []string{"a", "b", "x"}))               // true
}

func reportSpam(message []string, bannedWords []string) bool {
  // HashMap: O(1) lookup
	banned := make(map[string]struct{}, len(bannedWords))
	for _, w := range bannedWords {
		banned[w] = struct{}{}
	}
	count := 0
	for _, w := range message {
		if _, ok := banned[w]; ok {
			count++
			if count >= 2 {
				return true
			}
		}
	}
	return false
}
```
