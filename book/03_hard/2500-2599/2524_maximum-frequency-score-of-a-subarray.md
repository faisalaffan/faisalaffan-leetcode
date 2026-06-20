# 2524 — Maximum Frequency Score Of A Subarray

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func maxFrequencyScore(nums []int, k int) int
```

> **💡 Hint:** Sliding window with frequency and contribution tracking.

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap, Sliding Window

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2524: Maximum Frequency Score of a Subarray
// https://leetcode.com/problems/maximum-frequency-score-of-a-subarray/
// Difficulty: Hard [Paid]

import "fmt"

const mod = 1000000007

// maxFrequencyScore computes the maximum frequency score of any subarray of length k.
// Frequency score = sum over distinct values v of (freq[v] * v^freq[v]) % MOD.
//
// Approach: Sliding window with frequency and contribution tracking.
// Maintain current score. When freq[v] changes from f to f+1:
//   subtract old contribution (v^f) and add new (v^(f+1)).
// Use fast exponentiation with precomputed powers for O(1) window updates.
//
// Complexity: O(n * log MOD) time, O(distinct values) space
func maxFrequencyScore(nums []int, k int) int {
	n := len(nums)
	if k > n || k <= 0 {
		return 0
	}

  // Membuat map (HashMap) — pencarian O(1)
	freq := make(map[int]int)
	score := 0

	// Initialize first window
	for i := 0; i < k; i++ {
		v := nums[i]
		f := freq[v]
		// Add v^(f+1), remove old v^f (if f > 0)
		if f > 0 {
			score = (score - modPow(v, f) + mod) % mod
		}
		score = (score + modPow(v, f+1)) % mod
		freq[v] = f + 1
	}

	maxScore := score

	// Slide window
	for i := k; i < n; i++ {
		// Remove outgoing element nums[i-k]
		out := nums[i-k]
		f := freq[out] // current freq before removal
		// Remove out^f contribution, add out^(f-1) if f > 1
		score = (score - modPow(out, f) + mod) % mod
		if f > 1 {
			score = (score + modPow(out, f-1)) % mod
		}
		freq[out] = f - 1
		if freq[out] == 0 {
			delete(freq, out)
		}

		// Add incoming element nums[i]
		in := nums[i]
		f = freq[in] // current freq before addition
		if f > 0 {
			score = (score - modPow(in, f) + mod) % mod
		}
		score = (score + modPow(in, f+1)) % mod
		freq[in] = f + 1

		if score > maxScore {
			maxScore = score
		}
	}

	return maxScore
}

// modPow computes (base^exp) % mod using binary exponentiation.
func modPow(base, exp int) int {
	result := 1
	b := base % mod
	e := exp
	for e > 0 {
		if e&1 == 1 {
			result = result * b % mod
		}
		b = b * b % mod
		e >>= 1
	}
	return result
}

func main() {
	// Test cases
	fmt.Println("Test 1: nums=[1,1,1,2,2], k=3 ->", maxFrequencyScore([]int{1, 1, 1, 2, 2}, 3))
	fmt.Println("Test 2: nums=[1,2,3,4], k=2 ->", maxFrequencyScore([]int{1, 2, 3, 4}, 2))
	fmt.Println("Test 3: nums=[5,5,5], k=1 ->", maxFrequencyScore([]int{5, 5, 5}, 1))
	fmt.Println("Test 4: nums=[1,2], k=3 ->", maxFrequencyScore([]int{1, 2}, 3)) // k > n
	fmt.Println("Test 5: nums=[1,1,1], k=3 ->", maxFrequencyScore([]int{1, 1, 1}, 3))
	fmt.Println("Test 6: nums=[4,4,4,4], k=2 ->", maxFrequencyScore([]int{4, 4, 4, 4}, 2))
}
```
