# 2273 — Find Resultant Array After Removing Anagrams

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func FindResultantArrayAfterRemovingAnagrams(words []string) []string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2273: Find Resultant Array After Removing Anagrams
// https://leetcode.com/problems/find-resultant-array-after-removing-anagrams/
// Difficulty: Easy
// Time O(n * m log m) | Space O(n)

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(FindResultantArrayAfterRemovingAnagrams([]string{"abba", "baba", "bbaa", "cd", "cd"})) // ["abba","cd"]
	fmt.Println(FindResultantArrayAfterRemovingAnagrams([]string{"a", "b", "c", "d", "e"}))            // ["a","b","c","d","e"]
}

func FindResultantArrayAfterRemovingAnagrams(words []string) []string {
	result := []string{words[0]}
	prev := sortWord(words[0])

	for i := 1; i < len(words); i++ {
		curr := sortWord(words[i])
		if curr != prev {
			result = append(result, words[i])
			prev = curr
		}
	}
	return result
}

func sortWord(w string) string {
	b := []byte(w)
  // Custom sort dengan comparator
	sort.Slice(b, func(i, j int) bool { return b[i] < b[j] })
	return string(b)
}
```
