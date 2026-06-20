# 1805 — Number Of Different Integers In A String

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func NumDifferentIntegers(word string) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n), Space: O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1805: Number of Different Integers in a String
// https://leetcode.com/problems/number-of-different-integers-in-a-string/
// Difficulty: Easy

import "fmt"

// Time: O(n), Space: O(n)
func NumDifferentIntegers(word string) int {
  // Membuat map (HashMap) — pencarian O(1)
	seen := make(map[string]bool)
	i := 0
	for i < len(word) {
		if word[i] >= '0' && word[i] <= '9' {
			j := i
			for j < len(word) && word[j] >= '0' && word[j] <= '9' {
				j++
			}
			for i < j && word[i] == '0' {
				i++
			}
			num := word[i:j]
			if !seen[num] {
				seen[num] = true
			}
			i = j
		} else {
			i++
		}
	}
	return len(seen)
}

func main() {
	fmt.Println(NumDifferentIntegers("a123bc34d8ef34"))
	fmt.Println(NumDifferentIntegers("leet1234code234"))
	fmt.Println(NumDifferentIntegers("a1b01c001"))
}
```
