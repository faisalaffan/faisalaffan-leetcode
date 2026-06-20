# 3746 — Minimum String Length After Balanced Removals

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func minimumStringLengthAfterBalancedRemovals(s string) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3746: Minimum String Length After Balanced Removals
// https://leetcode.com/problems/minimum-string-length-after-balanced-removals/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func minimumStringLengthAfterBalancedRemovals(s string) int {
	a, b := 0, 0
  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(s); i++ {
		if s[i] == 'a' {
			a++
		} else {
			b++
		}
	}
	if a > b {
		return a - b
	}
	return b - a
}

func main() {
	fmt.Println(minimumStringLengthAfterBalancedRemovals("aabbab"))
	fmt.Println(minimumStringLengthAfterBalancedRemovals("aaaa"))
	fmt.Println(minimumStringLengthAfterBalancedRemovals("aaabb"))
}
```
