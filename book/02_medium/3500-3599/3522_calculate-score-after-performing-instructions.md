# 3522 — Calculate Score After Performing Instructions

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func CalculateScoreAfterPerformingInstructions(ops []string, vals []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3522: Calculate Score After Performing Instructions
// https://leetcode.com/problems/calculate-score-after-performing-instructions/
// Difficulty: Medium
// Complexity: O(n) time, O(1) space

import "fmt"

func main() {
	// Test case 1
	fmt.Println("Test 1:", CalculateScoreAfterPerformingInstructions([]string{"add", "add", "sub"}, []int{5, 3, 2}))
	// Test case 2
	fmt.Println("Test 2:", CalculateScoreAfterPerformingInstructions([]string{"add", "mul", "add"}, []int{1, 2, 3}))
	// Test case 3
	fmt.Println("Test 3:", CalculateScoreAfterPerformingInstructions([]string{"add"}, []int{10}))
}

func CalculateScoreAfterPerformingInstructions(ops []string, vals []int) int {
	score := 0
  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(ops) && i < len(vals); i++ {
		switch ops[i] {
		case "add":
			score += vals[i]
		case "sub":
			score -= vals[i]
		case "mul":
			score *= vals[i]
		case "div":
			if vals[i] != 0 {
				score /= vals[i]
			}
		}
	}
	return score
}
```
