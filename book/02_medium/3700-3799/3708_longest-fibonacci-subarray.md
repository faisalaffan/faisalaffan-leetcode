# 3708 — Longest Fibonacci Subarray

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func longestFibonacciSubarray(nums []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3708: Longest Fibonacci Subarray
// https://leetcode.com/problems/longest-fibonacci-subarray/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func longestFibonacciSubarray(nums []int) int {
	n := len(nums)
	if n <= 2 {
		return n
	}
	ans := 2
	cnt := 2
	for i := 2; i < n; i++ {
		if nums[i] == nums[i-1]+nums[i-2] {
			cnt++
		} else {
			cnt = 2
		}
		if cnt > ans {
			ans = cnt
		}
	}
	return ans
}

func main() {
	fmt.Println(longestFibonacciSubarray([]int{1, 1, 2, 3, 5, 8}))
	fmt.Println(longestFibonacciSubarray([]int{1, 3, 5, 8}))
	fmt.Println(longestFibonacciSubarray([]int{1, 2, 3, 4, 5, 6}))
}
```
