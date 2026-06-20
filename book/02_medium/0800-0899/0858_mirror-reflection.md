# 0858 — Mirror Reflection

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func MirrorReflection(p int, q int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** GCD / Matematika

**Kompleksitas Waktu:** O(log min(p,q))  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Kuasai dulu teknik **GCD / Matematika** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #858: Mirror Reflection
// https://leetcode.com/problems/mirror-reflection/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(MirrorReflection(2, 1))
	fmt.Println(MirrorReflection(3, 1))
	fmt.Println(MirrorReflection(4, 3))
}

// Time: O(log min(p,q)) | Space: O(1)
func MirrorReflection(p int, q int) int {
	g := gcd(p, q)
	p /= g
	q /= g

	if p%2 == 0 {
		return 2
	}
	if q%2 == 0 {
		return 0
	}
	return 1
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
```
