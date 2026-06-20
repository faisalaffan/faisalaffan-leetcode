# 1528 — Shuffle String

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func restoreString(s string, indices []int) string

import "fmt"

func main()
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n), Space: O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1528: Shuffle String
// https://leetcode.com/problems/shuffle-string/
// Difficulty: Easy
//
// LeetCode submission: func restoreString(s string, indices []int) string

import "fmt"

func main() {
	fmt.Println(ShuffleString("codeleet", []int{4, 5, 6, 7, 0, 2, 1, 3})) // "leetcode"
	fmt.Println(ShuffleString("abc", []int{0, 1, 2}))                     // "abc"
}

// Time: O(n), Space: O(n)
func ShuffleString(s string, indices []int) string {
	res := make([]byte, len(s))
	for i, idx := range indices {
		res[idx] = s[i]
	}
	return string(res)
}
```
