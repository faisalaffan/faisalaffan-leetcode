# 2649 — Nested Array Generator

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func nestedArrayGenerator(arr []any) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** DFS

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **DFS** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2649: Nested Array Generator
// https://leetcode.com/problems/nested-array-generator/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import "fmt"

func nestedArrayGenerator(arr []any) []int {
	result := []int{}
	var dfs func(item any)
	dfs = func(item any) {
		switch v := item.(type) {
		case []any:
			for _, sub := range v {
				dfs(sub)
			}
		case int:
			result = append(result, v)
		}
	}
	for _, item := range arr {
		dfs(item)
	}
	return result
}

func main() {
	// Test case 1: flat array
	arr1 := []any{1, 2, 3}
	fmt.Println("Test 1:", nestedArrayGenerator(arr1))
	// Expected: [1, 2, 3]

	// Test case 2: nested array
	arr2 := []any{1, []any{2, 3}, 4}
	fmt.Println("Test 2:", nestedArrayGenerator(arr2))
	// Expected: [1, 2, 3, 4]

	// Test case 3: deeply nested
	arr3 := []any{1, []any{2, []any{3, 4}, 5}, 6}
	fmt.Println("Test 3:", nestedArrayGenerator(arr3))
	// Expected: [1, 2, 3, 4, 5, 6]
}
```
