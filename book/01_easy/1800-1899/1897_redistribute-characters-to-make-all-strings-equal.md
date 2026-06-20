# 1897 — Redistribute Characters To Make All Strings Equal

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func MakeEqual(words []string) bool
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n * len), Space: O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1897: Redistribute Characters to Make All Strings Equal
// https://leetcode.com/problems/redistribute-characters-to-make-all-strings-equal/
// Difficulty: Easy

import "fmt"

// Time: O(n * len), Space: O(1)
func MakeEqual(words []string) bool {
  // Alokasi slice integer
	freq := make([]int, 26)
	for _, w := range words {
  // Loop linear O(n): iterasi setiap elemen
		for i := 0; i < len(w); i++ {
			freq[w[i]-'a']++
		}
	}
	n := len(words)
	for _, count := range freq {
		if count%n != 0 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(MakeEqual([]string{"abc", "aabc", "bc"}))
	fmt.Println(MakeEqual([]string{"ab", "a"}))
}
```
