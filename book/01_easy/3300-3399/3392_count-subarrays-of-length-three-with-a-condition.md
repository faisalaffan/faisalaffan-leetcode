# 3392 — Count Subarrays Of Length Three With A Condition

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func CountSubarraysOfLengthThreeWithACondition(nums []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n). Space: O(1).  
**Kompleksitas Ruang:** O(1).

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3392: Count Subarrays of Length Three With a Condition
// https://leetcode.com/problems/count-subarrays-of-length-three-with-a-condition/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(CountSubarraysOfLengthThreeWithACondition([]int{1, 2, 1, 2, 1}))
	fmt.Println(CountSubarraysOfLengthThreeWithACondition([]int{1, 3, 5, 7, 9}))
}

// CountSubarraysOfLengthThreeWithACondition counts subarrays of length 3 where the sum of first and last equals the middle.
// Time: O(n). Space: O(1).
func CountSubarraysOfLengthThreeWithACondition(nums []int) int {
	count := 0
	for i := 0; i <= len(nums)-3; i++ {
		if nums[i]+nums[i+2] == nums[i+1] {
			count++
		}
	}
	return count
}
```
