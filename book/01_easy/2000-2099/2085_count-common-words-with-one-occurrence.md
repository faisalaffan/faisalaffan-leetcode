# 2085 — Count Common Words With One Occurrence

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func CountCommonWordsWithOneOccurrence(words1 []string, words2 []string) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n + m), Space: O(n + m)  
**Kompleksitas Ruang:** O(n + m)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2085: Count Common Words With One Occurrence
// https://leetcode.com/problems/count-common-words-with-one-occurrence/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(CountCommonWordsWithOneOccurrence(
		[]string{"leetcode", "is", "amazing", "as", "is"},
		[]string{"amazing", "leetcode", "is"},
	)) // 2
	fmt.Println(CountCommonWordsWithOneOccurrence(
		[]string{"a", "ab"},
		[]string{"a", "a", "a", "ab"},
	)) // 1
}

// Time: O(n + m), Space: O(n + m)
func CountCommonWordsWithOneOccurrence(words1 []string, words2 []string) int {
  // Membuat map (HashMap) — pencarian O(1)
	freq1 := make(map[string]int)
  // Membuat map (HashMap) — pencarian O(1)
	freq2 := make(map[string]int)

	for _, w := range words1 {
		freq1[w]++
	}
	for _, w := range words2 {
		freq2[w]++
	}

	count := 0
	for w, c := range freq1 {
		if c == 1 && freq2[w] == 1 {
			count++
		}
	}
	return count
}
```
