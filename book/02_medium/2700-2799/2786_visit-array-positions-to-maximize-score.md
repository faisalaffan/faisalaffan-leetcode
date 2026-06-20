# 2786 — Visit Array Positions To Maximize Score

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func VisitArrayPositionsToMaximizeScore(nums []int, x int) int64
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Dynamic Programming

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Kuasai dulu teknik **Dynamic Programming** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2786: Visit Array Positions to Maximize Score
// https://leetcode.com/problems/visit-array-positions-to-maximize-score/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func VisitArrayPositionsToMaximizeScore(nums []int, x int) int64 {
	n := len(nums)
	// dp[0] = max score ending at even parity, dp[1] = max score ending at odd parity
	dp := [2]int64{int64(nums[0]), int64(nums[0])}
	parity0 := nums[0] & 1

	other := 1 - parity0
	dp[other] = int64(nums[0]) - int64(x)
	if dp[other] < int64(nums[0]) {
		dp[other] = int64(nums[0])
	}

	best := int64(nums[0])
	for i := 1; i < n; i++ {
		p := nums[i] & 1
		// Option 1: Start here
		cur := int64(nums[i])
		// Option 2: Extend from prev with same parity
		if int64(nums[i])+dp[p] > cur {
			cur = int64(nums[i]) + dp[p]
		}
		// Option 3: Extend from prev with different parity (pay x)
		otherP := 1 - p
		if int64(nums[i])-int64(x)+dp[otherP] > cur {
			cur = int64(nums[i]) - int64(x) + dp[otherP]
		}
		if cur > dp[p] {
			dp[p] = cur
		}
		if cur > best {
			best = cur
		}
	}

	return best
}

func main() {
	fmt.Println(VisitArrayPositionsToMaximizeScore([]int{2, 3, 6, 1, 9, 2}, 5))
	fmt.Println(VisitArrayPositionsToMaximizeScore([]int{2, 4, 6, 8}, 3))
}
```
