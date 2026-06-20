# 3411 — Maximum Subarray With Equal Products

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func MaximumSubarrayWithEqualProducts(nums []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** GCD / Matematika

**Kompleksitas Waktu:** O(n^2). Space: O(1).  
**Kompleksitas Ruang:** O(1).

> **Untuk fresh graduate:** Kuasai dulu teknik **GCD / Matematika** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3411: Maximum Subarray With Equal Products
// https://leetcode.com/problems/maximum-subarray-with-equal-products/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(MaximumSubarrayWithEqualProducts([]int{1, 2, 1, 2, 1, 1, 1}))
	fmt.Println(MaximumSubarrayWithEqualProducts([]int{2, 3, 4, 5, 6}))
}

// gcd returns the greatest common divisor of a and b.
func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// MaximumSubarrayWithEqualProducts returns the length of the longest subarray where product == lcm * gcd.
// Time: O(n^2). Space: O(1).
func MaximumSubarrayWithEqualProducts(nums []int) int {
	n := len(nums)
	maxLen := 0
	for i := 0; i < n; i++ {
		p := 1
		g := nums[i]
		l := nums[i]
		for j := i; j < n; j++ {
			p *= nums[j]
			g = gcd(g, nums[j])
			l = l * nums[j] / gcd(l, nums[j])
			if p == l*g {
				if j-i+1 > maxLen {
					maxLen = j - i + 1
				}
			}
		}
	}
	return maxLen
}
```
