# 0645 — Set Mismatch

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func findErrorNums(nums []int) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n). Space: O(1).  
**Kompleksitas Ruang:** O(1).

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #645: Set Mismatch
// https://leetcode.com/problems/set-mismatch/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(findErrorNums([]int{1, 2, 2, 4})) // [2, 3]
	fmt.Println(findErrorNums([]int{1, 1}))        // [1, 2]
	fmt.Println(findErrorNums([]int{2, 2}))        // [2, 1]
}

// findErrorNums finds the duplicated and missing number in the set.
// Time: O(n). Space: O(1).
func findErrorNums(nums []int) []int {
	n := len(nums)
	sum := 0
	sumSq := 0
	expectedSum := n * (n + 1) / 2
	expectedSumSq := n * (n + 1) * (2*n + 1) / 6

	for _, v := range nums {
		sum += v
		sumSq += v * v
	}

	// diff = duplicate - missing
	diff := sum - expectedSum
	// sqDiff = duplicate^2 - missing^2
	sqDiff := sumSq - expectedSumSq
	// duplicate + missing = sqDiff / diff
	plus := sqDiff / diff

	dup := (diff + plus) / 2
	miss := (plus - diff) / 2
	return []int{dup, miss}
}
```
