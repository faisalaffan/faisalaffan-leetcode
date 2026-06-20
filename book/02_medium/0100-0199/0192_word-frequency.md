# 0192 — Word Frequency

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func wordFrequency(text string) []string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n log n), Space: O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #192: Word Frequency
// https://leetcode.com/problems/word-frequency/
// Difficulty: Medium
// Time: O(n log n), Space: O(n)

import (
	"fmt"
	"sort"
	"strings"
)

func wordFrequency(text string) []string {
	words := strings.Fields(text)
  // Membuat map (HashMap) — pencarian O(1)
	freq := make(map[string]int)

	for _, w := range words {
		freq[w]++
	}

	type kv struct {
		word  string
		count int
	}

	var sorted []kv
	for w, c := range freq {
		sorted = append(sorted, kv{w, c})
	}

  // Custom sort dengan comparator
	sort.Slice(sorted, func(i, j int) bool {
		if sorted[i].count != sorted[j].count {
			return sorted[i].count > sorted[j].count
		}
		return sorted[i].word < sorted[j].word
	})

	result := make([]string, len(sorted))
	for i, kv := range sorted {
		result[i] = fmt.Sprintf("%s %d", kv.word, kv.count)
	}
	return result
}

func main() {
	for _, line := range wordFrequency("the day is sunny the the the sunny is is") {
		fmt.Println(line)
	}
	fmt.Println("---")
	for _, line := range wordFrequency("hello world hello") {
		fmt.Println(line)
	}
	fmt.Println("---")
	for _, line := range wordFrequency("") {
		fmt.Println(line)
	}
}
```
