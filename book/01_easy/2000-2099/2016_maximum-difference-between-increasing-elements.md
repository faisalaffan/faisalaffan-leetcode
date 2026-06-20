# 2016 — Maximum Difference Between Increasing Elements

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func MaximumDifferenceBetweenIncreasingElements(nums []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n), Space: O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2016: Maximum Difference Between Increasing Elements
// https://leetcode.com/problems/maximum-difference-between-increasing-elements/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(MaximumDifferenceBetweenIncreasingElements([]int{7, 1, 5, 4}))   // 4
	fmt.Println(MaximumDifferenceBetweenIncreasingElements([]int{9, 4, 3, 2}))   // -1
	fmt.Println(MaximumDifferenceBetweenIncreasingElements([]int{1, 5, 2, 10}))  // 9
}

// Time: O(n), Space: O(1)
func MaximumDifferenceBetweenIncreasingElements(nums []int) int {
	minSoFar := nums[0]
	maxDiff := -1
	for i := 1; i < len(nums); i++ {
		if nums[i] > minSoFar {
			diff := nums[i] - minSoFar
			if diff > maxDiff {
				maxDiff = diff
			}
		}
		if nums[i] < minSoFar {
			minSoFar = nums[i]
		}
	}
	return maxDiff
}
```
