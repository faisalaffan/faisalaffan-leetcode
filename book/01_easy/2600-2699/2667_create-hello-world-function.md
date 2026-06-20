# 2667 — Create Hello World Function

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func CreateHelloWorldFunction() func() string`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(1)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #2667: Create Hello World Function
// https://leetcode.com/problems/create-hello-world-function/
// Difficulty: Easy
// Time: O(1) | Space: O(1)
// Note: JavaScript problem, adapted to Go. Returns a function that always returns "Hello World".

import "fmt"

func main() {
	f := CreateHelloWorldFunction()
	fmt.Println(f())
	fmt.Println(f())
}

func CreateHelloWorldFunction() func() string {
	return func() string {
		return "Hello World"
	}
}
```
