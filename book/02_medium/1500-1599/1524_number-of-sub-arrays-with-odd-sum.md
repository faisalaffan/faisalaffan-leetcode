# 1524 — Number Of Sub Arrays With Odd Sum

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func NumOfSubarrays(arr []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Prefix Sum

**Kompleksitas Waktu:** O(N), Space: O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Kuasai dulu teknik **Prefix Sum** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1524: Number of Sub-arrays With Odd Sum
// https://leetcode.com/problems/number-of-sub-arrays-with-odd-sum/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(NumOfSubarrays([]int{1, 3, 5}))
	fmt.Println(NumOfSubarrays([]int{2, 4, 6}))
	fmt.Println(NumOfSubarrays([]int{1, 2, 3, 4, 5, 6, 7}))
}

func NumOfSubarrays(arr []int) int {
	// Time: O(N), Space: O(1)
	const mod = 1_000_000_007

	oddCount := 0
	evenCount := 1 // prefix sum = 0 is even
	prefixSum := 0
	result := 0

	for _, num := range arr {
		prefixSum += num

		if prefixSum%2 == 0 {
			// Current prefix is even
			result = (result + oddCount) % mod
			evenCount++
		} else {
			// Current prefix is odd
			result = (result + evenCount) % mod
			oddCount++
		}
	}

	return result
}
```
