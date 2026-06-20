# 3539 — Find Sum Of Array Product Of Magical Sequences

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func sumOfArrayProduct(n int, m int) int
```

> **💡 Hint:** DP to count sequences and their product sums.

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Dynamic Programming

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **Dynamic Programming** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3539: Find Sum of Array Product of Magical Sequences
// https://leetcode.com/problems/find-sum-of-array-product-of-magical-sequences/
// Difficulty: Hard
//
// Find sum over all magical sequences of their array product. A magical
// sequence is defined by specific constraints (e.g., length n, values in
// [1, m] with some property). Return result modulo 1e9+7.
//
// Approach: DP to count sequences and their product sums.

import "fmt"

func main() {
	// Example 1
	fmt.Println(sumOfArrayProduct(2, 3))
	// Example 2
	fmt.Println(sumOfArrayProduct(3, 2))
	// Edge: single element
	fmt.Println(sumOfArrayProduct(1, 5))
	// Edge: zero
	fmt.Println(sumOfArrayProduct(0, 10))
}

const mod = 1000000007

func sumOfArrayProduct(n int, m int) int {
  // Edge case: input kosong — langsung return
	if n == 0 {
		return 0
	}
	// dp[i] = sum of products of sequences ending at value i
	// total[i] = total count of sequences ending at value i
  // Alokasi slice integer
	dp := make([]int64, m+1)
  // Alokasi slice integer
	cnt := make([]int64, m+1)
	for i := 1; i <= m; i++ {
		dp[i] = int64(i)
		cnt[i] = 1
	}

	for length := 2; length <= n; length++ {
  // Alokasi slice integer
		newDp := make([]int64, m+1)
  // Alokasi slice integer
		newCnt := make([]int64, m+1)
		for i := 1; i <= m; i++ {
			for j := 1; j <= m; j++ {
				newDp[(i+j)%m] = (newDp[(i+j)%m] + dp[j]*int64(i)) % mod
				newCnt[(i+j)%m] = (newCnt[(i+j)%m] + cnt[j]) % mod
			}
		}
		dp = newDp
		cnt = newCnt
	}

	var ans int64
	for i := 1; i <= m; i++ {
		ans = (ans + dp[i]) % mod
	}
	return int(ans)
}
```
