# 3388 — Count Beautiful Splits In An Array

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func beautifulSplits(nums []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Prefix Sum

**Kompleksitas Waktu:** O(n^2) Space: O(n^2)  
**Kompleksitas Ruang:** O(n^2)

> **Untuk fresh graduate:** Kuasai dulu teknik **Prefix Sum** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3388: Count Beautiful Splits in an Array
// https://leetcode.com/problems/count-beautiful-splits-in-an-array/
// Difficulty: Medium
// Time: O(n^2) Space: O(n^2)

import "fmt"

func main() {
	fmt.Println(beautifulSplits([]int{1, 1, 2, 1})) // 2
	fmt.Println(beautifulSplits([]int{1, 2, 3, 4})) // 0
}

func beautifulSplits(nums []int) int {
	n := len(nums)

	// lcp[i][j] = longest common prefix of nums[i:] and nums[j:]
  // Membuat matriks/slice 2D untuk DP
	lcp := make([][]int, n+1)
  // Range loop: iterasi dengan indeks + nilai
	for i := range lcp {
		lcp[i] = make([]int, n+1)
	}
	for i := n - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if nums[i] == nums[j] {
				lcp[i][j] = lcp[i+1][j+1] + 1
			}
		}
	}

	ans := 0
	for i := 1; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			// nums1 = [0,i), nums2 = [i,j), nums3 = [j,n)
			ok := false
			// Check if nums1 is prefix of nums2
			if lcp[0][i] >= i {
				ok = true
			}
			// Check if nums2 is prefix of nums3
			if lcp[i][j] >= j-i {
				ok = true
			}
			if ok {
				ans++
			}
		}
	}
	return ans
}
```
