# 2295 — Replace Elements In An Array

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func arrayChange(nums []int, operations [][]int) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** O(n + m)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2295: Replace Elements in an Array
// https://leetcode.com/problems/replace-elements-in-an-array/
// Difficulty: Medium
// Time: O(n + m) | Space: O(n)

import "fmt"

func arrayChange(nums []int, operations [][]int) []int {
  // Membuat map (HashMap) — pencarian O(1)
	pos := make(map[int]int)
	for i, v := range nums {
		pos[v] = i
	}

	for _, op := range operations {
		oldVal, newVal := op[0], op[1]
		if idx, ok := pos[oldVal]; ok {
			nums[idx] = newVal
			delete(pos, oldVal)
			pos[newVal] = idx
		}
	}
	return nums
}

func main() {
	// Test case 1
	fmt.Println(arrayChange([]int{1, 2, 4, 6}, [][]int{{1, 3}, {4, 7}, {6, 1}}))
	// Expected: [3,2,7,1]

	// Test case 2
	fmt.Println(arrayChange([]int{1, 2}, [][]int{{1, 3}, {2, 1}, {3, 2}}))
	// Expected: [2,1]
}
```
