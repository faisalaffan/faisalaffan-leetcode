# 0806 — Number Of Lines To Write String

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func numberOfLines(widths []int, s string) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n). Space: O(1).  
**Kompleksitas Ruang:** O(1).

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #806: Number of Lines To Write String
// https://leetcode.com/problems/number-of-lines-to-write-string/
// Difficulty: Easy

import "fmt"

func main() {
	widths := []int{10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10}
	fmt.Println(numberOfLines(widths, "abcdefghijklmnopqrstuvwxyz")) // [3, 60]

	widths2 := []int{4, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10}
	fmt.Println(numberOfLines(widths2, "bbbcccdddaaa")) // [2, 4]
}

// numberOfLines returns the lines and last line width needed to write the string.
// Time: O(n). Space: O(1).
func numberOfLines(widths []int, s string) []int {
	lines, currentWidth := 1, 0
	for _, c := range s {
		w := widths[c-'a']
		if currentWidth+w > 100 {
			lines++
			currentWidth = w
		} else {
			currentWidth += w
		}
	}
	return []int{lines, currentWidth}
}
```
