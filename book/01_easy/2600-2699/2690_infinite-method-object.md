# 2690 — Infinite Method Object

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func InfiniteMethodObject() func(string) string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2690: Infinite Method Object
// https://leetcode.com/problems/infinite-method-object/
// Difficulty: Easy [Paid]
// Time: O(1) | Space: O(1)
// Note: JavaScript Proxy problem, adapted to Go. Returns an object that returns "methodName" for any method.

import "fmt"

func main() {
	obj := InfiniteMethodObject()
	fmt.Println(obj("abc"))
	fmt.Println(obj("xyz"))
}

func InfiniteMethodObject() func(string) string {
	return func(methodName string) string {
		return methodName
	}
}
```
