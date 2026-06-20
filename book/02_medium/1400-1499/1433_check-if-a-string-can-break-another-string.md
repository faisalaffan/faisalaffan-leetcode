# 1433 — Check If A String Can Break Another String

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func checkIfCanBreak(s1 string, s2 string) bool
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n log n) for sorting  
**Kompleksitas Ruang:** O(n) for byte slices

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1433: Check If a String Can Break Another String
// https://leetcode.com/problems/check-if-a-string-can-break-another-string/
// Difficulty: Medium

import "fmt"
import "sort"

func main() {
	// Test case 1
	fmt.Println(checkIfCanBreak("abc", "xya")) // true

	// Test case 2
	fmt.Println(checkIfCanBreak("abe", "acd")) // false

	// Test case 3
	fmt.Println(checkIfCanBreak("leetcodee", "interview")) // true

	// Test case 4
	fmt.Println(checkIfCanBreak("a", "b")) // true
}

// Time: O(n log n) for sorting
// Space: O(n) for byte slices
func checkIfCanBreak(s1 string, s2 string) bool {
	b1 := []byte(s1)
	b2 := []byte(s2)
  // Custom sort dengan comparator
	sort.Slice(b1, func(i, j int) bool { return b1[i] < b1[j] })
  // Custom sort dengan comparator
	sort.Slice(b2, func(i, j int) bool { return b2[i] < b2[j] })

	// Check if s1 can break s2
	canBreak1 := true
  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(b1); i++ {
		if b1[i] < b2[i] {
			canBreak1 = false
			break
		}
	}

	// Check if s2 can break s1
	canBreak2 := true
  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(b1); i++ {
		if b2[i] < b1[i] {
			canBreak2 = false
			break
		}
	}

	return canBreak1 || canBreak2
}
```
