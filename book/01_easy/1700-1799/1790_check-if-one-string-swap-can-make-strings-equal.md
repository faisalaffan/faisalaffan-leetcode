# 1790 — Check If One String Swap Can Make Strings Equal

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func AreAlmostEqual(s1 string, s2 string) bool
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n), Space: O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1790: Check if One String Swap Can Make Strings Equal
// https://leetcode.com/problems/check-if-one-string-swap-can-make-strings-equal/
// Difficulty: Easy

import "fmt"

// Time: O(n), Space: O(1)
func AreAlmostEqual(s1 string, s2 string) bool {
	var diff []int
  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			diff = append(diff, i)
			if len(diff) > 2 {
				return false
			}
		}
	}
	if len(diff) == 0 {
		return true
	}
	if len(diff) != 2 {
		return false
	}
	return s1[diff[0]] == s2[diff[1]] && s1[diff[1]] == s2[diff[0]]
}

func main() {
	fmt.Println(AreAlmostEqual("bank", "kanb"))
	fmt.Println(AreAlmostEqual("attack", "defend"))
	fmt.Println(AreAlmostEqual("kelb", "kelb"))
}
```
