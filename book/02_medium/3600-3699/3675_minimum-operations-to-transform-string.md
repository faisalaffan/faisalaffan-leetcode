# 3675 — Minimum Operations To Transform String

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func minimumOperationsToTransformString(s string) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3675: Minimum Operations to Transform String
// https://leetcode.com/problems/minimum-operations-to-transform-string/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func minimumOperationsToTransformString(s string) int {
	minChar := byte('z' + 1)
  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(s); i++ {
		c := s[i]
		if c != 'a' && c < minChar {
			minChar = c
			if minChar == 'b' {
				break
			}
		}
	}
	if minChar > 'z' {
		return 0
	}
	return int('z' + 1 - minChar)
}

func main() {
	fmt.Println(minimumOperationsToTransformString("yz"))
	fmt.Println(minimumOperationsToTransformString("a"))
	fmt.Println(minimumOperationsToTransformString("abc"))
}
```
