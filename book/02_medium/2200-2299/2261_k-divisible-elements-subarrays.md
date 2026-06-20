# 2261 — K Divisible Elements Subarrays

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func countDistinct(nums []int, k int, p int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n^2 * k)  
**Kompleksitas Ruang:** O(n^2)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2261: K Divisible Elements Subarrays
// https://leetcode.com/problems/k-divisible-elements-subarrays/
// Difficulty: Medium
// Time: O(n^2 * k) | Space: O(n^2)

import (
	"fmt"
	"strconv"
	"strings"
)

func countDistinct(nums []int, k int, p int) int {
	n := len(nums)
  // Membuat map (HashMap) — pencarian O(1)
	seen := make(map[string]bool)

	for i := 0; i < n; i++ {
		count := 0
		var sb strings.Builder
		for j := i; j < n; j++ {
			if nums[j]%p == 0 {
				count++
			}
			if count > k {
				break
			}
			if sb.Len() > 0 {
				sb.WriteString(",")
			}
			sb.WriteString(strconv.Itoa(nums[j]))
			seen[sb.String()] = true
		}
	}
	return len(seen)
}

func main() {
	// Test case 1
	fmt.Println(countDistinct([]int{2, 3, 3, 2, 2}, 2, 2))
	// Expected: 11

	// Test case 2
	fmt.Println(countDistinct([]int{1, 2, 3, 4}, 4, 5))
	// Expected: 10
}
```
