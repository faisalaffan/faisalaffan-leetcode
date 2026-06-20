# 2897 — Apply Operations On Array To Maximize Sum Of Squares

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func maxSum2897(nums []int, k int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2897: Apply Operations on Array to Maximize Sum of Squares
// https://leetcode.com/problems/apply-operations-to-maximize-sum-of-squares/
// Difficulty: Hard
//
// The operation (a,b) -> (a&b, a|b) preserves total bit count per position but
// concentrates bits. Sum of squares is convex: we maximize by concentrating bits
// into as few numbers as possible. For each bit position, count available bits,
// then greedily build k largest possible numbers by taking one available bit from
// each position (highest first) for each of the k selections.
// O((N + k) * 32) time, O(32) space.

import "fmt"

const mod2897 = 1000000007

func maxSum2897(nums []int, k int) int {
	// Count bits at each position
  // Alokasi slice integer
	cnt := make([]int, 32)
	for _, v := range nums {
		for b := 0; b < 32; b++ {
			if v>>b&1 == 1 {
				cnt[b]++
			}
		}
	}

	// Build k largest numbers greedily
	result := int64(0)
	for i := 0; i < k; i++ {
		cur := int64(0)
		for b := 31; b >= 0; b-- {
			if cnt[b] > 0 {
				cur |= (1 << b)
				cnt[b]--
			}
		}
		result = (result + cur*cur) % mod2897
	}

	return int(result)
}

func main() {
	// Example: nums=[2,6,5,8], k=2 => 261
	fmt.Println(maxSum2897([]int{2, 6, 5, 8}, 2))
	// Example 2: nums=[4,5,4,7], k=3 => 90
	fmt.Println(maxSum2897([]int{4, 5, 4, 7}, 3))
	// k=1 (just take largest possible number)
	fmt.Println(maxSum2897([]int{2, 3, 4, 5}, 1))
	// All zeros
	fmt.Println(maxSum2897([]int{0, 0, 0}, 2))
	// Single element
	fmt.Println(maxSum2897([]int{7}, 1))
	// Simple
	fmt.Println(maxSum2897([]int{5, 6, 3}, 2))
	// All same values
	fmt.Println(maxSum2897([]int{3, 3, 3}, 2))
	// k larger than count of bits
	fmt.Println(maxSum2897([]int{8, 4, 2}, 5))
}
```
