# 0833 — Find And Replace In String

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func FindAndReplaceInString(s string, indices []int, sources []string, targets []string) string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n + m) where n = len(s), m = len(indices)  
**Kompleksitas Ruang:** O(n + m)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #833: Find And Replace in String
// https://leetcode.com/problems/find-and-replace-in-string/
// Difficulty: Medium

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(FindAndReplaceInString("abcd", []int{0, 2}, []string{"a", "cd"}, []string{"eee", "ffff"}))
	fmt.Println(FindAndReplaceInString("abcd", []int{0, 2}, []string{"ab", "ec"}, []string{"eee", "ffff"}))
	fmt.Println(FindAndReplaceInString("jjievdtjfb", []int{4, 6}, []string{"md", "tjfb"}, []string{"foe", "oov"}))
}

// Time: O(n + m) where n = len(s), m = len(indices) | Space: O(n + m)
func FindAndReplaceInString(s string, indices []int, sources []string, targets []string) string {
	n := len(s)
  // Alokasi slice integer
	replace := make([]int, n)
  // Range loop: iterasi dengan indeks + nilai
	for i := range replace {
		replace[i] = -1
	}

	for k, idx := range indices {
		if idx+len(sources[k]) <= n && s[idx:idx+len(sources[k])] == sources[k] {
			replace[idx] = k
		}
	}

	var sb strings.Builder
	for i := 0; i < n; {
		if replace[i] >= 0 {
			sb.WriteString(targets[replace[i]])
			i += len(sources[replace[i]])
		} else {
			sb.WriteByte(s[i])
			i++
		}
	}

	return sb.String()
}
```
