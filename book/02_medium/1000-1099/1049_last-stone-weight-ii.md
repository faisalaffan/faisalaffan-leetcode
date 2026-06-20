# 1049 — Last Stone Weight Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func lastStoneWeightII(stones []int) int
```

> **💡 Hint:** DP - subset sum. Partition stones into two groups.

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Dynamic Programming

**Kompleksitas Waktu:** O(n * sum)  
**Kompleksitas Ruang:** O(sum)

> **Untuk fresh graduate:** Kuasai dulu teknik **Dynamic Programming** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1049: Last Stone Weight II
// https://leetcode.com/problems/last-stone-weight-ii/
// Difficulty: Medium
//
// Approach: DP - subset sum. Partition stones into two groups.
//           Minimize |sum - 2*subsetSum|.
// Time: O(n * sum)
// Space: O(sum)

import "fmt"

func main() {
	fmt.Println(lastStoneWeightII([]int{2, 7, 4, 1, 8, 1})) // 1
	fmt.Println(lastStoneWeightII([]int{31, 26, 33, 21, 40})) // 5
}

func lastStoneWeightII(stones []int) int {
	total := 0
	for _, s := range stones {
		total += s
	}

	target := total / 2
	dp := make([]bool, target+1)
	dp[0] = true

	for _, s := range stones {
		for j := target; j >= s; j-- {
			if dp[j-s] {
				dp[j] = true
			}
		}
	}

	for j := target; j >= 0; j-- {
		if dp[j] {
			return total - 2*j
		}
	}

	return total
}
```
