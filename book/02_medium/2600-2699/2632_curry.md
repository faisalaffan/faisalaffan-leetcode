# 2632 — Curry

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func curry2New(fn func(int, int) int) *curry2
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(1) per call  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2632: Curry
// https://leetcode.com/problems/curry/
// Difficulty: Medium [Paid]
// Time: O(1) per call | Space: O(1)

import "fmt"

// curry2 implements currying for a 2-argument function
type curry2 struct {
	fn     func(int, int) int
	args   []int
	arity  int
}

func curry2New(fn func(int, int) int) *curry2 {
	return &curry2{fn: fn, arity: 2}
}

func (c *curry2) call(args ...int) *curry2 {
	c.args = append(c.args, args...)
	return c
}

func (c *curry2) done() int {
	return c.fn(c.args[0], c.args[1])
}

func add(a, b int) int {
	return a + b
}

func main() {
	fn := add

	// Test case 1: currying with two separate calls
	c := curry2New(fn)
	r1 := c.call(1).call(2).done()
	fmt.Println("Test 1:", r1)
	// Expected: 3

	// Test case 2: currying with one call
	c2 := curry2New(fn)
	r2 := c2.call(3, 4).done()
	fmt.Println("Test 2:", r2)
	// Expected: 7
}
```
