# 3255 — Find The Power Of K Size Subarrays Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func resultsArray(nums []int, k int) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3255: Find the Power of K-Size Subarrays II
// https://leetcode.com/problems/find-the-power-of-k-size-subarrays-ii/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func resultsArray(nums []int, k int) []int {
	n := len(nums)
  // Alokasi slice integer
	ans := make([]int, n-k+1)
	consec := 1

  // Range loop: iterasi dengan indeks + nilai
	for i := range n {
		if i > 0 && nums[i] == nums[i-1]+1 {
			consec++
		} else {
			consec = 1
		}
		if i >= k-1 {
			if consec >= k {
				ans[i-k+1] = nums[i]
			} else {
				ans[i-k+1] = -1
			}
		}
	}
	return ans
}

func main() {
	fmt.Println(resultsArray([]int{1, 2, 3, 4, 3, 5}, 3)) // Expected: [3, 4, -1, -1]
	fmt.Println(resultsArray([]int{1, 4, 5, 2, 3}, 3))    // Expected: [-1, 5, -1]
	fmt.Println(resultsArray([]int{1}, 1))                 // Expected: [1]
}
```
