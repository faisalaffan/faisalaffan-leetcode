# 0890 — Find And Replace Pattern

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan data terstruktur dan diminta untuk mencari elemen atau pola tertentu. Tugasmu adalah menemukan posisi, jumlah, atau keberadaan elemen dengan efisien.

Seperti mencari kata di kamus — kamu tidak membaca dari halaman 1, tapi langsung ke tengah (binary search), lalu maju/mundur. Teknik pencarian yang efisien sangat penting untuk interview.

**Konsep kunci:** linear search O(n), binary search O(log n), HashMap lookup O(1).

**Fungsi yang perlu kamu implementasikan:**
```go
func FindAndReplacePattern(words []string, pattern string) []string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n * m) where n = len(words), m = avg word length  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #890: Find and Replace Pattern
// https://leetcode.com/problems/find-and-replace-pattern/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(FindAndReplacePattern([]string{"abc", "deq", "mee", "aqq", "dkd", "ccc"}, "abb"))
	fmt.Println(FindAndReplacePattern([]string{"a", "b", "c"}, "a"))
	fmt.Println(FindAndReplacePattern([]string{"aa", "ab"}, "aa"))
}

// Time: O(n * m) where n = len(words), m = avg word length | Space: O(n)
func FindAndReplacePattern(words []string, pattern string) []string {
	var ans []string
	for _, word := range words {
		if isMatch(word, pattern) {
			ans = append(ans, word)
		}
	}
	return ans
}

func isMatch(word, pattern string) bool {
	if len(word) != len(pattern) {
		return false
	}
  // Membuat map (HashMap) — pencarian O(1)
	w2p := make(map[byte]byte)
  // Membuat map (HashMap) — pencarian O(1)
	p2w := make(map[byte]byte)

  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(word); i++ {
		wc, pc := word[i], pattern[i]
		if v, ok := w2p[wc]; ok && v != pc {
			return false
		}
		if v, ok := p2w[pc]; ok && v != wc {
			return false
		}
		w2p[wc] = pc
		p2w[pc] = wc
	}

	return true
}
```
