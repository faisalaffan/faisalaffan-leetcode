# 0161 — One Edit Distance

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func isOneEditDistance(s string, t string) bool
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n), Space: O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #161: One Edit Distance
// https://leetcode.com/problems/one-edit-distance/
// Difficulty: Medium [Paid]
// Time: O(n), Space: O(1)

import "fmt"

func isOneEditDistance(s string, t string) bool {
	ns, nt := len(s), len(t)
	if abs(ns-nt) > 1 {
		return false
	}

	if ns > nt {
		s, t = t, s
		ns, nt = nt, ns
	}

	for i := 0; i < ns; i++ {
		if s[i] != t[i] {
			if ns == nt {
				return s[i+1:] == t[i+1:]
			}
			return s[i:] == t[i+1:]
		}
	}

	return ns+1 == nt
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	fmt.Println(isOneEditDistance("ab", "acb"))
	fmt.Println(isOneEditDistance("", ""))
	fmt.Println(isOneEditDistance("a", ""))
}
```
