# 2527 — Find Xor Beauty Of Array

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func xorBeauty(nums []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2527: Find Xor-Beauty of Array
// https://leetcode.com/problems/find-xor-beauty-of-array/
// Difficulty: Medium
// Time: O(n) | Space: O(1)
// XOR of all (nums[i]|nums[j]) & nums[k] over all i,j,k = XOR of all nums[i].
// Because the expression simplifies: each bit appears in result iff it appears odd times in nums.

import "fmt"

func main() {
	fmt.Println(xorBeauty([]int{1, 4})) // 5
	fmt.Println(xorBeauty([]int{15, 45, 20, 2, 34, 35, 5, 44, 32, 30})) // 34
}

func xorBeauty(nums []int) int {
	ans := 0
	for _, v := range nums {
		ans ^= v
	}
	return ans
}
```
