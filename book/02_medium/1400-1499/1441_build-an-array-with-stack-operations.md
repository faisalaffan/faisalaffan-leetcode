# 1441 — Build An Array With Stack Operations

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func buildArray(target []int, n int) []string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Stack

**Kompleksitas Waktu:** O(n) where n = max number in target  
**Kompleksitas Ruang:** O(n) for result

> **Untuk fresh graduate:** Kuasai dulu teknik **Stack** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1441: Build an Array With Stack Operations
// https://leetcode.com/problems/build-an-array-with-stack-operations/
// Difficulty: Medium

import "fmt"

func main() {
	// Test case 1
	fmt.Println(buildArray([]int{1, 3}, 3)) // ["Push","Push","Pop","Push"]

	// Test case 2
	fmt.Println(buildArray([]int{1, 2, 3}, 3)) // ["Push","Push","Push"]

	// Test case 3
	fmt.Println(buildArray([]int{1, 2}, 4)) // ["Push","Push"]
}

// Time: O(n) where n = max number in target
// Space: O(n) for result
func buildArray(target []int, n int) []string {
	result := make([]string, 0, n*2)
	current := 1

	for _, t := range target {
		for current < t {
			result = append(result, "Push", "Pop")
			current++
		}
		result = append(result, "Push")
		current++
	}

	return result
}
```
