# 1540 — Can Convert String In K Moves

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func CanConvertString(s string, t string, k int) bool
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(N), Space: O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1540: Can Convert String in K Moves
// https://leetcode.com/problems/can-convert-string-in-k-moves/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(CanConvertString("input", "ouput", 9))
	fmt.Println(CanConvertString("abc", "bcd", 10))
	fmt.Println(CanConvertString("aab", "bbb", 27))
}

func CanConvertString(s string, t string, k int) bool {
	// Time: O(N), Space: O(1)
	if len(s) != len(t) {
		return false
	}

	// Count how many times each shift value is needed
  // Alokasi slice integer
	shiftCount := make([]int, 26)

  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(s); i++ {
		if s[i] == t[i] {
			continue
		}
		// Compute needed shift (positive modulo 26)
		shift := (int(t[i]) - int(s[i]) + 26) % 26
		if shift == 0 {
			continue
		}

		// For each subsequent time we need this shift, add 26
		shiftCount[shift]++
		needed := shift + (shiftCount[shift]-1)*26
		if needed > k {
			return false
		}
	}

	return true
}
```
