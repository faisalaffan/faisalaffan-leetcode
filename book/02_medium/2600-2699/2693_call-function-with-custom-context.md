# 2693 — Call Function With Custom Context

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func callWithContext(fn func(Context, ...int) int, ctx Context, args ...int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2693: Call Function with Custom Context
// https://leetcode.com/problems/call-function-with-custom-context/
// Difficulty: Medium
// Time: O(1) | Space: O(1)

import "fmt"

type Context map[string]any

func callWithContext(fn func(Context, ...int) int, ctx Context, args ...int) int {
	return fn(ctx, args...)
}

func main() {
	// Test case 1
	sumFn := func(ctx Context, args ...int) int {
		total := 0
		for _, v := range args {
			total += v
		}
		return total
	}
	fmt.Println("Test 1:", callWithContext(sumFn, Context{}, 1, 2, 3))
	// Expected: 6

	// Test case 2
	ctxFn := func(ctx Context, args ...int) int {
		if multiplier, ok := ctx["mult"]; ok {
			m := multiplier.(int)
			result := 0
			for _, v := range args {
				result += v * m
			}
			return result
		}
		return 0
	}
	fmt.Println("Test 2:", callWithContext(ctxFn, Context{"mult": 3}, 1, 2, 3))
	// Expected: 18

	// Test case 3
	fmt.Println("Test 3:", callWithContext(sumFn, Context{}))
	// Expected: 0
}
```
