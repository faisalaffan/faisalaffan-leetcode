# 2873 — Maximum Value Of An Ordered Triplet I

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func MaximumValueOfAnOrderedTripletI(nums []int) int64
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2873: Maximum Value of an Ordered Triplet I
// https://leetcode.com/problems/maximum-value-of-an-ordered-triplet-i/
// Difficulty: Easy
// Time: O(n) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(MaximumValueOfAnOrderedTripletI([]int{12, 6, 1, 2, 7}))
	fmt.Println(MaximumValueOfAnOrderedTripletI([]int{1, 10, 3, 2, 5}))
}

func MaximumValueOfAnOrderedTripletI(nums []int) int64 {
	n := len(nums)
	if n < 3 {
		return 0
	}

	maxNum := int64(nums[0])
	maxDiff := int64(nums[0] - nums[1])
	ans := int64(0)

	for i := 2; i < n; i++ {
		val := maxDiff * int64(nums[i])
		if val > ans {
			ans = val
		}
		if int64(nums[i-1]) > maxNum {
			maxNum = int64(nums[i-1])
		}
		if maxNum-int64(nums[i]) > maxDiff {
			maxDiff = maxNum - int64(nums[i])
		}
	}

	return ans
}
```
