# 2150 — Find All Lonely Numbers In The Array

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func findLonely(nums []int) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2150: Find All Lonely Numbers in the Array
// https://leetcode.com/problems/find-all-lonely-numbers-in-the-array/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import "fmt"

func findLonely(nums []int) []int {
  // Membuat map (HashMap) — pencarian O(1)
	freq := make(map[int]int)
	for _, v := range nums {
		freq[v]++
	}

	result := []int{}
	for _, v := range nums {
		if freq[v] == 1 && freq[v-1] == 0 && freq[v+1] == 0 {
			result = append(result, v)
		}
	}
	return result
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", findLonely([]int{10, 6, 5, 8}))
	// Expected: [10, 8]

	// Test case 2
	fmt.Println("Test 2:", findLonely([]int{1, 3, 5, 3}))
	// Expected: [1, 5]
}
```
