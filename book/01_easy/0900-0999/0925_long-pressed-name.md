# 0925 — Long Pressed Name

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func isLongPressedName(name string, typed string) bool
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n + m). Space: O(1).  
**Kompleksitas Ruang:** O(1).

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #925: Long Pressed Name
// https://leetcode.com/problems/long-pressed-name/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(isLongPressedName("alex", "aaleex"))    // true
	fmt.Println(isLongPressedName("saeed", "ssaaedd"))  // false
	fmt.Println(isLongPressedName("leelee", "lleeelee")) // true
	fmt.Println(isLongPressedName("alex", "aaleexa"))   // false
}

// isLongPressedName checks if typed is a long-pressed version of name.
// Time: O(n + m). Space: O(1).
func isLongPressedName(name string, typed string) bool {
	if len(typed) < len(name) {
		return false
	}
	i, j := 0, 0
	for i < len(name) && j < len(typed) {
		if name[i] != typed[j] {
			return false
		}
		// Count occurrences in name
		c1 := 0
		ch := name[i]
		for i < len(name) && name[i] == ch {
			i++
			c1++
		}
		// Count occurrences in typed
		c2 := 0
		for j < len(typed) && typed[j] == ch {
			j++
			c2++
		}
		if c2 < c1 {
			return false
		}
	}
	return i == len(name) && j == len(typed)
}
```
