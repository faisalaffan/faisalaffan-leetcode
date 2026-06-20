# 3314 — Construct The Minimum Bitwise Array I

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func ConstructTheMinimumBitwiseArrayI(nums []int) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n * min_val). Space: O(n).  
**Kompleksitas Ruang:** O(n).

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3314: Construct the Minimum Bitwise Array I
// https://leetcode.com/problems/construct-the-minimum-bitwise-array-i/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(ConstructTheMinimumBitwiseArrayI([]int{2, 3, 5, 7}))
	fmt.Println(ConstructTheMinimumBitwiseArrayI([]int{11, 13, 31}))
}

// ConstructTheMinimumBitwiseArrayI returns an array where ans[i] is the smallest number such that ans[i] | (ans[i]+1) == nums[i].
// Time: O(n * min_val). Space: O(n).
func ConstructTheMinimumBitwiseArrayI(nums []int) []int {
  // Alokasi slice integer
	result := make([]int, len(nums))
	for i, num := range nums {
		found := false
		for candidate := 0; candidate < num; candidate++ {
			if candidate|(candidate+1) == num {
				result[i] = candidate
				found = true
				break
			}
		}
		if !found {
			result[i] = -1
		}
	}
	return result
}
```
