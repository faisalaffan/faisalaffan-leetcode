# 2629 — Function Composition

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func functionComposition(functions []func(int) int) func(int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Two Pointer

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Kuasai dulu teknik **Two Pointer** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2629: Function Composition
// https://leetcode.com/problems/function-composition/
// Difficulty: Easy
// Time: O(n) | Space: O(1)
// Note: JavaScript problem, adapted to Go. Composes functions right-to-left.

import "fmt"

func main() {
	double := func(x int) int { return x * 2 }
	addOne := func(x int) int { return x + 1 }
	square := func(x int) int { return x * x }

	fns := []func(int) int{square, double, addOne}
	composed := functionComposition(fns)
	fmt.Println(composed(5)) // (5+1)*2 squared = 144
}

func functionComposition(functions []func(int) int) func(int) int {
	return func(x int) int {
		result := x
		for i := len(functions) - 1; i >= 0; i-- {
			result = functions[i](result)
		}
		return result
	}
}
```
