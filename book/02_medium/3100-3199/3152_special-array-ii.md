# 3152 — Special Array Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func isArraySpecial(nums []int, queries [][]int) []bool
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Prefix Sum

**Kompleksitas Waktu:** O(n + q)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **Prefix Sum** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3152: Special Array II
// https://leetcode.com/problems/special-array-ii/
// Difficulty: Medium
// Time: O(n + q) | Space: O(n)

import "fmt"

func isArraySpecial(nums []int, queries [][]int) []bool {
	n := len(nums)
  // Alokasi slice integer
	prefix := make([]int, n)
	for i := 1; i < n; i++ {
		prefix[i] = prefix[i-1]
		if nums[i]%2 == nums[i-1]%2 {
			prefix[i]++
		}
	}

	ans := make([]bool, len(queries))
	for i, q := range queries {
		from, to := q[0], q[1]
		ans[i] = prefix[from] == prefix[to]
	}
	return ans
}

func main() {
	fmt.Println(isArraySpecial([]int{3, 4, 1, 2, 6}, [][]int{{0, 4}}))        // Expected: [false]
	fmt.Println(isArraySpecial([]int{4, 3, 1, 6}, [][]int{{0, 2}, {2, 3}}))   // Expected: [false, true]
}
```
