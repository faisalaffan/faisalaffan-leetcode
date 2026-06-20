# 2772 — Apply Operations To Make All Array Elements Equal To Zero

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func ApplyOperationsToMakeAllArrayElementsEqualToZero(nums []int, k int) bool
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2772: Apply Operations to Make All Array Elements Equal to Zero
// https://leetcode.com/problems/apply-operations-to-make-all-array-elements-equal-to-zero/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import "fmt"

func ApplyOperationsToMakeAllArrayElementsEqualToZero(nums []int, k int) bool {
	n := len(nums)
  // Alokasi slice integer
	diff := make([]int, n+1)
	cur := 0

	for i := 0; i < n; i++ {
		cur += diff[i]
		val := nums[i] + cur
		if val < 0 {
			return false
		}
		val %= 2
		if val != 0 {
			if i+k > n {
				return false
			}
			diff[i] -= 1
			diff[i+k] += 1
			cur -= 1
		}
	}

	return true
}

func main() {
	fmt.Println(ApplyOperationsToMakeAllArrayElementsEqualToZero([]int{2, 0, 2}, 2))
	fmt.Println(ApplyOperationsToMakeAllArrayElementsEqualToZero([]int{1, 0, 1}, 2))
}
```
