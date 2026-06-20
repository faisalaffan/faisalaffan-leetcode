# 2932 — Maximum Strong Pair Xor I

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func MaximumStrongPairXorI(nums []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n^2)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2932: Maximum Strong Pair XOR I
// https://leetcode.com/problems/maximum-strong-pair-xor-i/
// Difficulty: Easy

import "fmt"

func main() {
	// LeetCode name: maximumStrongPairXor
	fmt.Println(MaximumStrongPairXorI([]int{1, 2, 3, 4, 5})) // 7
	fmt.Println(MaximumStrongPairXorI([]int{10, 100}))        // 0
	fmt.Println(MaximumStrongPairXorI([]int{5, 6, 25, 30}))   // 7
}

// Time: O(n^2) | Space: O(1)
// LeetCode submission name: maximumStrongPairXor
func MaximumStrongPairXorI(nums []int) int {
	n := len(nums)
	maxXor := 0
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			x, y := nums[i], nums[j]
			if abs(x-y) <= min(x, y) {
				if x^y > maxXor {
					maxXor = x ^ y
				}
			}
		}
	}
	return maxXor
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
```
