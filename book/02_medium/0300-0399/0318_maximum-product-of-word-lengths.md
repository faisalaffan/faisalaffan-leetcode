# 0318 — Maximum Product Of Word Lengths

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func maxProduct(words []string) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n^2)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #318: Maximum Product of Word Lengths
// https://leetcode.com/problems/maximum-product-of-word-lengths/
// Difficulty: Medium
// Time: O(n^2) | Space: O(n)

import "fmt"

func maxProduct(words []string) int {
  // Alokasi slice integer
	bits := make([]int, len(words))
	for i, w := range words {
		mask := 0
		for _, ch := range w {
			mask |= 1 << (ch - 'a')
		}
		bits[i] = mask
	}

	maxProd := 0
  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(words); i++ {
		for j := i + 1; j < len(words); j++ {
			if bits[i]&bits[j] == 0 {
				prod := len(words[i]) * len(words[j])
				if prod > maxProd {
					maxProd = prod
				}
			}
		}
	}
	return maxProd
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", maxProduct([]string{"abcw", "baz", "foo", "bar", "xtfn", "abcdef"}))
	// Expected: 16 (abcw * xtfn)

	// Test case 2
	fmt.Println("Test 2:", maxProduct([]string{"a", "ab", "abc", "d", "cd", "bcd", "abcd"}))
	// Expected: 4 (ab * cd)

	// Test case 3
	fmt.Println("Test 3:", maxProduct([]string{"a", "aa", "aaa", "aaaa"}))
	// Expected: 0
}
```
