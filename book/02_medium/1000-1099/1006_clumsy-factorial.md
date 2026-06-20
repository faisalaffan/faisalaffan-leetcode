# 1006 — Clumsy Factorial

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func clumsy(n int) int
```

> **💡 Hint:** Math - use stack to handle operator precedence (*, / before +, -)

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Stack

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **Stack** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1006: Clumsy Factorial
// https://leetcode.com/problems/clumsy-factorial/
// Difficulty: Medium
//
// Approach: Math - use stack to handle operator precedence (*, / before +, -)
// Time: O(n)
// Space: O(n)

import "fmt"

func main() {
	fmt.Println(clumsy(4))  // 7
	fmt.Println(clumsy(10)) // 12
	fmt.Println(clumsy(1))  // 1
}

func clumsy(n int) int {
	stack := []int{n}
	op := 0

	for i := n - 1; i >= 1; i-- {
		switch op % 4 {
		case 0: // *
			stack[len(stack)-1] *= i
		case 1: // /
			stack[len(stack)-1] /= i
		case 2: // +
			stack = append(stack, i)
		case 3: // -
			stack = append(stack, -i)
		}
		op++
	}

	result := 0
	for _, v := range stack {
		result += v
	}
	return result
}
```
