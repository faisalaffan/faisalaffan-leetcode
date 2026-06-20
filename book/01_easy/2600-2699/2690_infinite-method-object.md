# 2690 — Infinite Method Object

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func InfiniteMethodObject() func(string) string`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(1)  |  **Ruang:** O(1)


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
