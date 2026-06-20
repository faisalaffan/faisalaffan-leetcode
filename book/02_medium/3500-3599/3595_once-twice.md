# 3595 — Once Twice

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func OnceTwice(nums []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3595: Once Twice
// https://leetcode.com/problems/once-twice/
// Difficulty: Medium [Paid]
// Complexity: O(n) time, O(1) space

import "fmt"

func main() {
	// Test case 1
	fmt.Println("Test 1:", OnceTwice([]int{1, 1, 2, 2, 3}))
	// Test case 2
	fmt.Println("Test 2:", OnceTwice([]int{1, 1, 1, 2, 2}))
	// Test case 3
	fmt.Println("Test 3:", OnceTwice([]int{1, 2, 3}))
}

func OnceTwice(nums []int) int {
	// Count elements that appear twice vs once
  // Membuat map (HashMap) — pencarian O(1)
	freq := make(map[int]int)
	for _, v := range nums {
		freq[v]++
	}
	once, twice := 0, 0
	for _, c := range freq {
		if c == 1 {
			once++
		} else if c >= 2 {
			twice++
		}
	}
	return once * twice
}
```
