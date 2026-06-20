# 3014 — Minimum Number Of Pushes To Type Word I

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func MinimumNumberOfPushesToTypeWordI(word string) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3014: Minimum Number of Pushes to Type Word I
// https://leetcode.com/problems/minimum-number-of-pushes-to-type-word-i/
// Difficulty: Easy

import "fmt"

func main() {
	// LeetCode name: minimumPushes
	fmt.Println(MinimumNumberOfPushesToTypeWordI("abcde")) // 5
	fmt.Println(MinimumNumberOfPushesToTypeWordI("xycdefghij")) // 12
}

// Time: O(n) | Space: O(1)
// LeetCode submission name: minimumPushes
// Each key can hold up to 4 distinct letters. First 8 letters cost 1 push each,
// next 8 cost 2 pushes, etc.
func MinimumNumberOfPushesToTypeWordI(word string) int {
	n := len(word)
	pushes := 0
	// First 8 distinct chars: 1 push each
	// Next 8: 2 pushes each, etc.
	for i := 0; i < n; i++ {
		pushes += (i / 8) + 1
	}
	return pushes
}
```
