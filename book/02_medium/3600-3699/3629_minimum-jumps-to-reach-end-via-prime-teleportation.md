# 3629 — Minimum Jumps To Reach End Via Prime Teleportation

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func isPrime(x int) bool
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Dynamic Programming

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **Dynamic Programming** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3629: Minimum Jumps to Reach End via Prime Teleportation
// https://leetcode.com/problems/minimum-jumps-to-reach-end-via-prime-teleportation/
// Difficulty: Medium
// Complexity: O(n) time, O(n) space

import "fmt"

func isPrime(x int) bool {
	if x < 2 {
		return false
	}
	for i := 2; i*i <= x; i++ {
		if x%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", MinimumJumpsToReachEndViaPrimeTeleportation([]int{1, 2, 3, 4, 5}))
	// Test case 2
	fmt.Println("Test 2:", MinimumJumpsToReachEndViaPrimeTeleportation([]int{4, 6, 8, 10}))
	// Test case 3
	fmt.Println("Test 3:", MinimumJumpsToReachEndViaPrimeTeleportation([]int{2, 3, 5, 7, 11}))
}

func MinimumJumpsToReachEndViaPrimeTeleportation(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return 0
	}
	// DP: min jumps to reach position i
  // Alokasi slice integer
	dp := make([]int, n)
  // Range loop: iterasi dengan indeks + nilai
	for i := range dp {
		dp[i] = n + 1
	}
	dp[0] = 0
	for i := 0; i < n; i++ {
		if isPrime(nums[i]) {
			// Can teleport to any position
			for j := 0; j < n; j++ {
				if j != i && dp[i]+1 < dp[j] {
					dp[j] = dp[i] + 1
				}
			}
		} else {
			// Can move to adjacent position
			if i+1 < n && dp[i]+1 < dp[i+1] {
				dp[i+1] = dp[i] + 1
			}
			if i-1 >= 0 && dp[i]+1 < dp[i-1] {
				dp[i-1] = dp[i] + 1
			}
		}
	}
	if dp[n-1] > n {
		return -1
	}
	return dp[n-1]
}
```
