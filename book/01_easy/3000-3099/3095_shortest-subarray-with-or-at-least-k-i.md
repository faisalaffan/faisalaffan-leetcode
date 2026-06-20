# 3095 — Shortest Subarray With Or At Least K I

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func ShortestSubarrayWithOrAtLeastKI(nums []int, k int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n^2)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3095: Shortest Subarray With OR at Least K I
// https://leetcode.com/problems/shortest-subarray-with-or-at-least-k-i/
// Difficulty: Easy

import "fmt"

func main() {
	// LeetCode name: minimumSubarrayLength
	fmt.Println(ShortestSubarrayWithOrAtLeastKI([]int{1, 2, 3}, 2)) // 1
	fmt.Println(ShortestSubarrayWithOrAtLeastKI([]int{2, 1, 8}, 10)) // 3
	fmt.Println(ShortestSubarrayWithOrAtLeastKI([]int{1, 2}, 10))    // -1
}

// Time: O(n^2) | Space: O(1)
// LeetCode submission name: minimumSubarrayLength
func ShortestSubarrayWithOrAtLeastKI(nums []int, k int) int {
	n := len(nums)
	minLen := n + 1
	for i := 0; i < n; i++ {
		orVal := 0
		for j := i; j < n; j++ {
			orVal |= nums[j]
			if orVal >= k {
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
	return minLen
}
```
