# 3479 — Fruits Into Baskets Iii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func FruitsIntoBasketsIii(fruits []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap, Two Pointer

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3479: Fruits Into Baskets III
// https://leetcode.com/problems/fruits-into-baskets-iii/
// Difficulty: Medium
// Complexity: O(n) time, O(n) space

import "fmt"

func main() {
	// Test case 1
	fmt.Println("Test 1:", FruitsIntoBasketsIii([]int{1, 2, 1}))
	// Test case 2
	fmt.Println("Test 2:", FruitsIntoBasketsIii([]int{0, 1, 2, 2}))
	// Test case 3
	fmt.Println("Test 3:", FruitsIntoBasketsIii([]int{1, 2, 3, 2, 2}))
}

func FruitsIntoBasketsIii(fruits []int) int {
  // Membuat map (HashMap) — pencarian O(1)
	freq := make(map[int]int)
	left := 0
	maxLen := 0
	for right := 0; right < len(fruits); right++ {
		freq[fruits[right]]++
		for len(freq) > 2 {
			freq[fruits[left]]--
			if freq[fruits[left]] == 0 {
				delete(freq, fruits[left])
			}
			left++
		}
		if right-left+1 > maxLen {
			maxLen = right - left + 1
		}
	}
	return maxLen
}
```
