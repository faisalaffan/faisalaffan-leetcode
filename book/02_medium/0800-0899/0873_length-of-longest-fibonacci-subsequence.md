# 0873 — Length Of Longest Fibonacci Subsequence

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func LengthOfLongestFibonacciSubsequence(arr []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap, Dynamic Programming

**Kompleksitas Waktu:** O(n^2)  
**Kompleksitas Ruang:** O(n^2)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #873: Length of Longest Fibonacci Subsequence
// https://leetcode.com/problems/length-of-longest-fibonacci-subsequence/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(LengthOfLongestFibonacciSubsequence([]int{1, 2, 3, 4, 5, 6, 7, 8}))
	fmt.Println(LengthOfLongestFibonacciSubsequence([]int{1, 3, 7, 11, 12, 14, 18}))
	fmt.Println(LengthOfLongestFibonacciSubsequence([]int{1, 3, 5}))
}

// Time: O(n^2) | Space: O(n^2)
func LengthOfLongestFibonacciSubsequence(arr []int) int {
	n := len(arr)
  // Membuat map (HashMap) — pencarian O(1)
	index := make(map[int]int, n)
	for i, v := range arr {
		index[v] = i
	}

  // Membuat matriks/slice 2D untuk DP
	dp := make([][]int, n)
  // Range loop: iterasi dengan indeks + nilai
	for i := range dp {
		dp[i] = make([]int, n)
	}

	ans := 0
	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			prev := arr[i] - arr[j]
			if k, ok := index[prev]; ok && k < j {
				dp[i][j] = dp[j][k] + 1
				if dp[i][j] > ans {
					ans = dp[i][j]
				}
			} else {
				dp[i][j] = 2
			}
		}
	}

	if ans >= 3 {
		return ans
	}
	return 0
}
```
