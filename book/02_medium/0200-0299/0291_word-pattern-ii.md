# 0291 — Word Pattern Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func wordPatternMatch(pattern string, s string) bool
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Backtracking

**Kompleksitas Waktu:** O(2^n) worst case, Space: O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **Backtracking** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #291: Word Pattern II
// https://leetcode.com/problems/word-pattern-ii/
// Difficulty: Medium [Paid]
// Time: O(2^n) worst case, Space: O(n)

import "fmt"

func wordPatternMatch(pattern string, s string) bool {
  // Membuat map (HashMap) — pencarian O(1)
	pMap := make(map[byte]string)
  // Membuat map (HashMap) — pencarian O(1)
	sMap := make(map[string]byte)

	var backtrack func(patIdx, strIdx int) bool
	backtrack = func(patIdx, strIdx int) bool {
		if patIdx == len(pattern) && strIdx == len(s) {
			return true
		}
		if patIdx >= len(pattern) || strIdx >= len(s) {
			return false
		}

		ch := pattern[patIdx]
		if mapped, ok := pMap[ch]; ok {
			if strIdx+len(mapped) > len(s) || s[strIdx:strIdx+len(mapped)] != mapped {
				return false
			}
			return backtrack(patIdx+1, strIdx+len(mapped))
		}

		for end := strIdx + 1; end <= len(s); end++ {
			candidate := s[strIdx:end]
			if existing, ok := sMap[candidate]; ok && existing != ch {
				continue
			}

			pMap[ch] = candidate
			sMap[candidate] = ch
			if backtrack(patIdx+1, end) {
				return true
			}
			delete(pMap, ch)
			delete(sMap, candidate)
		}

		return false
	}

	return backtrack(0, 0)
}

func main() {
	fmt.Println(wordPatternMatch("abab", "redblueredblue"))
	fmt.Println(wordPatternMatch("aaaa", "asdasdasdasd"))
	fmt.Println(wordPatternMatch("ab", "aa"))
}
```
