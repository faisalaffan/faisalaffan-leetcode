# 2834 — Find The Minimum Possible Sum Of A Beautiful Array

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func FindTheMinimumPossibleSumOfABeautifulArray(n int, target int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2834: Find the Minimum Possible Sum of a Beautiful Array
// https://leetcode.com/problems/find-the-minimum-possible-sum-of-a-beautiful-array/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import "fmt"

func FindTheMinimumPossibleSumOfABeautifulArray(n int, target int) int {
  // Membuat map (HashMap) — pencarian O(1)
	used := make(map[int]bool)
	sum := 0
	for i := 1; len(used) < n; i++ {
		if used[target-i] {
			continue
		}
		used[i] = true
		sum += i
	}
	return sum
}

func main() {
	fmt.Println(FindTheMinimumPossibleSumOfABeautifulArray(2, 3))
	fmt.Println(FindTheMinimumPossibleSumOfABeautifulArray(3, 3))
	fmt.Println(FindTheMinimumPossibleSumOfABeautifulArray(5, 5))
}
```
