# 2665 — Counter Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func CounterIi(init int) *Counter
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2665: Counter II
// https://leetcode.com/problems/counter-ii/
// Difficulty: Easy
// Time: O(1) | Space: O(1)
// Note: JavaScript problem, adapted to Go. Returns an object with increment/decrement/reset.

import "fmt"

func main() {
	counter := CounterIi(5)
	fmt.Println(counter.increment())
	fmt.Println(counter.reset())
	fmt.Println(counter.decrement())
}

type Counter struct {
	init  int
	value int
}

func (c *Counter) increment() int {
	c.value++
	return c.value
}

func (c *Counter) decrement() int {
	c.value--
	return c.value
}

func (c *Counter) reset() int {
	c.value = c.init
	return c.value
}

func CounterIi(init int) *Counter {
	return &Counter{init: init, value: init}
}
```
