# 2011 — Final Value Of Variable After Performing Operations

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func FinalValueOfVariableAfterPerformingOperations(operations []string) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n), Space: O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2011: Final Value of Variable After Performing Operations
// https://leetcode.com/problems/final-value-of-variable-after-performing-operations/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(FinalValueOfVariableAfterPerformingOperations([]string{"--X", "X++", "X++"}))     // 1
	fmt.Println(FinalValueOfVariableAfterPerformingOperations([]string{"++X", "++X", "X++"}))     // 3
	fmt.Println(FinalValueOfVariableAfterPerformingOperations([]string{"X++", "++X", "--X", "X--"})) // 0
}

// Time: O(n), Space: O(1)
func FinalValueOfVariableAfterPerformingOperations(operations []string) int {
	x := 0
	for _, op := range operations {
		if op[1] == '+' {
			x++
		} else {
			x--
		}
	}
	return x
}
```
