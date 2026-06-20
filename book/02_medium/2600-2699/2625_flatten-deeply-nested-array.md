# 2625 — Flatten Deeply Nested Array

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func flatten(arr []any, depth int) []any
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** DFS

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **DFS** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2625: Flatten Deeply Nested Array
// https://leetcode.com/problems/flatten-deeply-nested-array/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import "fmt"

func flatten(arr []any, depth int) []any {
	result := []any{}

	var dfs func(item any, currDepth int)
	dfs = func(item any, currDepth int) {
		switch v := item.(type) {
		case []any:
			if currDepth < depth {
				for _, sub := range v {
					dfs(sub, currDepth+1)
				}
			} else {
				result = append(result, v)
			}
		default:
			result = append(result, v)
		}
	}

	for _, item := range arr {
		dfs(item, 0)
	}
	return result
}

func main() {
	// Test case 1: flatten to depth 1
	arr1 := []any{1, []any{2, []any{3, 4}}, 5}
	fmt.Println("Test 1:", flatten(arr1, 1))
	// Expected: [1, 2, [3, 4], 5]

	// Test case 2: flatten to depth 2
	fmt.Println("Test 2:", flatten(arr1, 2))
	// Expected: [1, 2, 3, 4, 5]

	// Test case 3: depth 0 - no flattening
	fmt.Println("Test 3:", flatten(arr1, 0))
	// Expected: [1, [2, [3, 4]], 5]
}
```
