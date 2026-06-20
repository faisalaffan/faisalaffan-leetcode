# 2640 — Find The Score Of All Prefixes Of An Array

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func findPrefixScore(nums []int) []int64
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Prefix Sum

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **Prefix Sum** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2640: Find the Score of All Prefixes of an Array
// https://leetcode.com/problems/find-the-score-of-all-prefixes-of-an-array/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import "fmt"

func findPrefixScore(nums []int) []int64 {
	n := len(nums)
  // Alokasi slice integer
	ans := make([]int64, n)
	maxSoFar := nums[0]
	prefixSum := int64(0)

	for i, v := range nums {
		if v > maxSoFar {
			maxSoFar = v
		}
		conver := int64(v + maxSoFar)
		prefixSum += conver
		ans[i] = prefixSum
	}
	return ans
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", findPrefixScore([]int{2, 3, 7, 5, 10}))
	// Expected: [4,10,24,36,56]

	// Test case 2
	fmt.Println("Test 2:", findPrefixScore([]int{1, 1, 1}))
	// Expected: [2,4,6]

	// Test case 3
	fmt.Println("Test 3:", findPrefixScore([]int{5}))
	// Expected: [10]
}
```
