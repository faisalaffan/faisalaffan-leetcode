# 3146 — Permutation Difference Between Two Strings

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func PermutationDifferenceBetweenTwoStrings(s string, t string) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3146: Permutation Difference between Two Strings
// https://leetcode.com/problems/permutation-difference-between-two-strings/
// Difficulty: Easy

import "fmt"

func main() {
	// LeetCode name: findPermutationDifference
	fmt.Println(PermutationDifferenceBetweenTwoStrings("abc", "bac")) // 2
	fmt.Println(PermutationDifferenceBetweenTwoStrings("abcde", "edcba")) // 12
}

// Time: O(n) | Space: O(1)
// LeetCode submission name: findPermutationDifference
func PermutationDifferenceBetweenTwoStrings(s string, t string) int {
  // Membuat map (HashMap) — pencarian O(1)
	pos := make(map[byte]int)
  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(t); i++ {
		pos[t[i]] = i
	}
	diff := 0
  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(s); i++ {
		d := i - pos[s[i]]
		if d < 0 {
			d = -d
		}
		diff += d
	}
	return diff
}
```
