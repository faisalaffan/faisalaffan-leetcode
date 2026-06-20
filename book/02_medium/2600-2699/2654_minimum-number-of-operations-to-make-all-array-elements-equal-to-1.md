# 2654 — Minimum Number Of Operations To Make All Array Elements Equal To 1

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func minOperations(nums []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** GCD / Matematika

**Kompleksitas Waktu:** O(n^2)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Kuasai dulu teknik **GCD / Matematika** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2654: Minimum Number of Operations to Make All Array Elements Equal to 1
// https://leetcode.com/problems/minimum-number-of-operations-to-make-all-array-elements-equal-to-1/
// Difficulty: Medium
// Time: O(n^2) | Space: O(1)

import "fmt"

func minOperations(nums []int) int {
	n := len(nums)

	// If there's already a 1, answer is n - count(1)
	ones := 0
	for _, v := range nums {
		if v == 1 {
			ones++
		}
	}
	if ones > 0 {
		return n - ones
	}

	// Find shortest subarray with GCD=1
	minLen := n + 1
	for i := 0; i < n; i++ {
		g := nums[i]
		for j := i; j < n; j++ {
			g = gcd(g, nums[j])
			if g == 1 {
				if j-i+1 < minLen {
					minLen = j - i + 1
				}
				break
			}
		}
	}

	if minLen > n {
		return -1
	}

	// Need (minLen-1) operations to convert subarray to 1, then (n-1) to spread
	return (minLen - 1) + (n - 1)
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", minOperations([]int{2, 6, 3, 4}))
	// Expected: 4

	// Test case 2
	fmt.Println("Test 2:", minOperations([]int{2, 10, 6, 14}))
	// Expected: -1

	// Test case 3: already has 1
	fmt.Println("Test 3:", minOperations([]int{1, 5, 3}))
	// Expected: 2
}
```
