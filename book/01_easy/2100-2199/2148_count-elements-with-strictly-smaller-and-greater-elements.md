# 2148 — Count Elements With Strictly Smaller And Greater Elements

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func CountElementsWithStrictlySmallerAndGreaterElements(nums []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n), Space: O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2148: Count Elements With Strictly Smaller and Greater Elements
// https://leetcode.com/problems/count-elements-with-strictly-smaller-and-greater-elements/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(CountElementsWithStrictlySmallerAndGreaterElements([]int{11, 7, 2, 15}))   // 2
	fmt.Println(CountElementsWithStrictlySmallerAndGreaterElements([]int{-3, 3, 3, 90}))   // 2
	fmt.Println(CountElementsWithStrictlySmallerAndGreaterElements([]int{1, 2, 3}))         // 1
}

// Time: O(n), Space: O(1)
func CountElementsWithStrictlySmallerAndGreaterElements(nums []int) int {
	if len(nums) < 3 {
		return 0
	}

	min, max := nums[0], nums[0]
	for _, v := range nums {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}

	if min == max {
		return 0
	}

	count := 0
	for _, v := range nums {
		if v > min && v < max {
			count++
		}
	}
	return count
}
```
