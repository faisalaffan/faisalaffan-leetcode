# 1427 — Perform String Shifts

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func stringShift(s string, shift [][]int) string

import "fmt"

func main()
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n + m), Space: O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1427: Perform String Shifts
// https://leetcode.com/problems/perform-string-shifts/
// Difficulty: Easy [Paid]
//
// LeetCode submission: func stringShift(s string, shift [][]int) string

import "fmt"

func main() {
	fmt.Println(PerformStringShifts("abc", [][]int{{0, 1}, {1, 2}}))             // "cab"
	fmt.Println(PerformStringShifts("abcdefg", [][]int{{1, 1}, {1, 1}, {0, 2}, {1, 3}})) // "efgabcd"
}

// Time: O(n + m), Space: O(n)
func PerformStringShifts(s string, shift [][]int) string {
	total := 0
	for _, sh := range shift {
		if sh[0] == 0 {
			total -= sh[1]
		} else {
			total += sh[1]
		}
	}
	n := len(s)
	total %= n
	if total < 0 {
		total += n
	}
	return s[n-total:] + s[:n-total]
}
```
