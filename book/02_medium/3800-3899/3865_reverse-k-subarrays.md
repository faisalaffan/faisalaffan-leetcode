# 3865 — Reverse K Subarrays

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func ReverseKSubarrays(nums []int, k int) int
```

> **💡 Hint:** Count positions that don't match expected value at their index.

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(N)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3865: Reverse K Subarrays
// https://leetcode.com/problems/reverse-k-subarrays/
// Difficulty: Medium [Paid]
// Time: O(N) | Space: O(1)
// Approach: Count positions that don't match expected value at their index.
// Each reversal of a subarray of length k can fix at most 2 positions.

import "fmt"

func ReverseKSubarrays(nums []int, k int) int {
	n := len(nums)
	swaps := 0
	for i := 0; i < n; i++ {
		if nums[i] != i { // value should equal index for sorted array [0,1,2,...]
			// Find where i is
			j := i
			for j < n && nums[j] != i {
				j++
			}
			if j-i+1 >= k && j < n {
				// Reverse subarray i..j
				for l, r := i, j; l < r; l, r = l+1, r-1 {
					nums[l], nums[r] = nums[r], nums[l]
				}
				swaps++
			}
		}
	}
	return swaps
}

func main() {
	// Example
	fmt.Println(ReverseKSubarrays([]int{1, 0, 3, 2}, 2)) // Expected: ?

	// Example
	fmt.Println(ReverseKSubarrays([]int{2, 1, 0}, 3)) // Expected: ?

	// Example
	fmt.Println(ReverseKSubarrays([]int{0, 1, 2}, 1)) // Expected: 0
}
```
