# 2629 — Function Composition

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func functionComposition(functions []func(int) int) func(int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Two Pointer

**Waktu:** O(n)  |  **Ruang:** O(1)

> 🎓 **Fresh Grad Tips:** Kuasai **Two Pointer** — sering muncul di interview!

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
