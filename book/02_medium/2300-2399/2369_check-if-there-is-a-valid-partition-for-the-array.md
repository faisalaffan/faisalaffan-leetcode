# 2369 — Check If There Is A Valid Partition For The Array

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func validPartition(nums []int) bool
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Dynamic Programming

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Kuasai dulu teknik **Dynamic Programming** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2369: Check if There is a Valid Partition For The Array
// https://leetcode.com/problems/check-if-there-is-a-valid-partition-for-the-array/
// Difficulty: Medium
// Time: O(n) | Space: O(1)
// DP: dp[i] = valid partition for first i elements. Check three conditions.

import "fmt"

func main() {
	fmt.Println(validPartition([]int{4, 4, 4, 5, 6}))     // true
	fmt.Println(validPartition([]int{1, 1, 1, 2}))          // false
	fmt.Println(validPartition([]int{1, 2, 3, 4, 5, 6}))    // true
}

func validPartition(nums []int) bool {
	n := len(nums)
	dp := make([]bool, n+1)
	dp[0] = true
	for i := 1; i <= n; i++ {
		if i >= 2 && dp[i-2] && nums[i-2] == nums[i-1] {
			dp[i] = true
		}
		if i >= 3 && dp[i-3] {
			if (nums[i-3] == nums[i-2] && nums[i-2] == nums[i-1]) ||
				(nums[i-3]+1 == nums[i-2] && nums[i-2]+1 == nums[i-1]) {
				dp[i] = true
			}
		}
	}
	return dp[n]
}
```
