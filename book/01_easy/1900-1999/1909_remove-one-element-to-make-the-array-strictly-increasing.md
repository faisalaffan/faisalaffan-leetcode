# 1909 — Remove One Element To Make The Array Strictly Increasing

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func RemoveOneElementToMakeTheArrayStrictlyIncreasing(nums []int) bool
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n), Space: O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1909: Remove One Element to Make the Array Strictly Increasing
// https://leetcode.com/problems/remove-one-element-to-make-the-array-strictly-increasing/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(RemoveOneElementToMakeTheArrayStrictlyIncreasing([]int{1, 2, 10, 5, 7}))       // true
	fmt.Println(RemoveOneElementToMakeTheArrayStrictlyIncreasing([]int{2, 3, 1, 2}))           // false
	fmt.Println(RemoveOneElementToMakeTheArrayStrictlyIncreasing([]int{1, 1, 1}))              // false
}

// Time: O(n), Space: O(1)
func RemoveOneElementToMakeTheArrayStrictlyIncreasing(nums []int) bool {
	removed := false
	for i := 1; i < len(nums); i++ {
		if nums[i] <= nums[i-1] {
			if removed {
				return false
			}
			removed = true
			// Try removing nums[i-1] or nums[i]
			if i-2 < 0 || nums[i] > nums[i-2] {
				// Removing nums[i-1] works
			} else if i+1 >= len(nums) || nums[i+1] > nums[i-1] {
				// Removing nums[i] works, skip it
				nums[i] = nums[i-1] // adjust for next comparison
			} else {
				return false
			}
		}
	}
	return true
}
```
