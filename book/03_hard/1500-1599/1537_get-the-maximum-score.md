# 1537 — Get The Maximum Score

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func maxSum(nums1 []int, nums2 []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1537: Get the Maximum Score
// https://leetcode.com/problems/get-the-maximum-score/
// Difficulty: Hard
//
// Two-pointer + cumulative sum approach:
// 1. Traverse both sorted arrays with two pointers.
// 2. Keep running sums for each path.
// 3. When values equal, we can switch paths; take max of both sums.
// 4. Continue accumulating after the switch point.

import "fmt"

func main() {
	// Example: [2,4,5,8,10], [4,6,8,9] -> 30
	// Path: 2+4+6+8+10 = 30 (switch at 4, switch at 8)
	fmt.Println(maxSum([]int{2, 4, 5, 8, 10}, []int{4, 6, 8, 9}))

	// Additional tests
	fmt.Println(maxSum([]int{1, 3, 5, 7, 9}, []int{2, 4, 6, 8, 10}))
	fmt.Println(maxSum([]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}))
	fmt.Println(maxSum([]int{1, 4, 5, 8, 12}, []int{2, 3, 6, 10, 11}))
	fmt.Println(maxSum([]int{1}, []int{1}))
}

const mod = 1_000_000_007

func maxSum(nums1 []int, nums2 []int) int {
	i, j := 0, 0
	sum1, sum2 := 0, 0
	n, m := len(nums1), len(nums2)

	for i < n && j < m {
		if nums1[i] < nums2[j] {
			sum1 += nums1[i]
			i++
		} else if nums1[i] > nums2[j] {
			sum2 += nums2[j]
			j++
		} else {
			// Equal — switch point: take max of both paths
			best := max(sum1, sum2) + nums1[i]
			sum1 = best
			sum2 = best
			i++
			j++
		}
	}

	for i < n {
		sum1 += nums1[i]
		i++
	}
	for j < m {
		sum2 += nums2[j]
		j++
	}

	return max(sum1, sum2) % mod
}
```
