# 2754 — Bind Function To Context

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func BindFunctionToContext(fn func(Context, ...interface{}) interface
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2754: Bind Function to Context
// https://leetcode.com/problems/bind-function-to-context/
// Difficulty: Medium [Paid]
// Time: O(1) | Space: O(1)

import "fmt"

type Context map[string]interface{}

type BoundFunc struct {
	fn      func(ctx Context, args ...interface{}) interface{}
	ctx     Context
}

func BindFunctionToContext(fn func(Context, ...interface{}) interface{}, ctx Context) func(...interface{}) interface{} {
	return func(args ...interface{}) interface{} {
		return fn(ctx, args...)
	}
}

func main() {
	ctx := Context{"multiplier": 3}
	bound := BindFunctionToContext(func(ctx Context, args ...interface{}) interface{} {
		m := ctx["multiplier"].(int)
		return args[0].(int) * m
	}, ctx)
	fmt.Println(bound(5))
	fmt.Println(bound(10))
}
```
