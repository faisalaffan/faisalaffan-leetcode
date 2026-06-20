# 1663 — Smallest String With A Given Numeric Value

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func GetSmallestString(n int, k int) string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(N), Space: O(N)  
**Kompleksitas Ruang:** O(N)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1663: Smallest String With A Given Numeric Value
// https://leetcode.com/problems/smallest-string-with-a-given-numeric-value/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(GetSmallestString(3, 27))
	fmt.Println(GetSmallestString(5, 73))
	fmt.Println(GetSmallestString(1, 26))
}

func GetSmallestString(n int, k int) string {
	// Time: O(N), Space: O(N)
	// Greedy: fill from the end with 'z' as much as possible
	result := make([]byte, n)
  // Range loop: iterasi dengan indeks + nilai
	for i := range result {
		result[i] = 'a'
	}

	k -= n // All positions have at least 'a' (value 1)

	for i := n - 1; i >= 0 && k > 0; i-- {
		add := 25 // 'z' - 'a' = 25
		if k < add {
			add = k
		}
		result[i] = byte('a' + add)
		k -= add
	}

	return string(result)
}
```
