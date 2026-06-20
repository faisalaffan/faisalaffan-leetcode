# 1684 — Count The Number Of Consistent Strings

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func CountConsistentStrings(allowed string, words []string) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n + m*k), Space: O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1684: Count the Number of Consistent Strings
// https://leetcode.com/problems/count-the-number-of-consistent-strings/
// Difficulty: Easy

import "fmt"

// Time: O(n + m*k), Space: O(1)
func CountConsistentStrings(allowed string, words []string) int {
  // Membuat map (HashMap) — pencarian O(1)
	allowedSet := make(map[byte]bool)
  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(allowed); i++ {
		allowedSet[allowed[i]] = true
	}
	count := 0
	for _, word := range words {
		consistent := true
  // Loop linear O(n): iterasi setiap elemen
		for i := 0; i < len(word); i++ {
			if !allowedSet[word[i]] {
				consistent = false
				break
			}
		}
		if consistent {
			count++
		}
	}
	return count
}

func main() {
	fmt.Println(CountConsistentStrings("ab", []string{"ad", "bd", "aaab", "baa", "badab"}))
	fmt.Println(CountConsistentStrings("abc", []string{"a", "b", "c", "ab", "ac", "bc", "abc"}))
	fmt.Println(CountConsistentStrings("cad", []string{"cc", "acd", "b", "ba", "bac", "bad", "ac", "d"}))
}
```
