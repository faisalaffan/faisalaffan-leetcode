# 2124 — Check If All As Appears Before All Bs

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func CheckIfAllAsAppearsBeforeAllBs(s string) bool
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n), Space: O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2124: Check if All A's Appears Before All B's
// https://leetcode.com/problems/check-if-all-as-appears-before-all-bs/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(CheckIfAllAsAppearsBeforeAllBs("aaabbb")) // true
	fmt.Println(CheckIfAllAsAppearsBeforeAllBs("abab"))   // false
	fmt.Println(CheckIfAllAsAppearsBeforeAllBs("bbb"))    // true
}

// Time: O(n), Space: O(1)
func CheckIfAllAsAppearsBeforeAllBs(s string) bool {
	foundB := false
  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(s); i++ {
		if s[i] == 'b' {
			foundB = true
		} else if s[i] == 'a' && foundB {
			return false
		}
	}
	return true
}
```
