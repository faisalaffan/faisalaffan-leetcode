# 2726 — Calculator With Method Chaining

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func CalculatorWithMethodChaining(initialValue int) *Calculator
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2726: Calculator with Method Chaining
// https://leetcode.com/problems/calculator-with-method-chaining/
// Difficulty: Easy
// Time: O(1) | Space: O(1)
// Note: JavaScript problem, adapted to Go. Calculator with method chaining.

import "fmt"

func main() {
	result := CalculatorWithMethodChaining(10).add(5).subtract(7).getResult()
	fmt.Println(result)
}

type Calculator struct {
	value int
}

func (c *Calculator) add(val int) *Calculator {
	c.value += val
	return c
}

func (c *Calculator) subtract(val int) *Calculator {
	c.value -= val
	return c
}

func (c *Calculator) multiply(val int) *Calculator {
	c.value *= val
	return c
}

func (c *Calculator) divide(val int) *Calculator {
	c.value /= val
	return c
}

func (c *Calculator) power(val int) *Calculator {
	result := 1
	for i := 0; i < val; i++ {
		result *= c.value
	}
	c.value = result
	return c
}

func (c *Calculator) getResult() int {
	return c.value
}

func CalculatorWithMethodChaining(initialValue int) *Calculator {
	return &Calculator{value: initialValue}
}
```
