# 1317 — Convert Integer To The Sum Of Two No Zero Integers

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan array integer dan sebuah target. Tugasmu adalah mencari **dua angka** yang jika dijumlahkan menghasilkan target. Kembalikan **indeks** (posisi) kedua angka.

Contoh: `nums=[2,7,11,15], target=9` → `2+7=9` → `[0,1]`.

**Cara berpikir:** Gunakan HashMap. Untuk setiap angka, cek apakah `target-angka` sudah ada di map. Kalau sudah → ketemu pasangan. Kalau belum → simpan angka ke map.

**Fungsi Solusi:** `func getNoZeroIntegers(n int) []int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n log n), Space: O(1)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #1317: Convert Integer to the Sum of Two No-Zero Integers
// https://leetcode.com/problems/convert-integer-to-the-sum-of-two-no-zero-integers/
// Difficulty: Easy
//
// LeetCode submission: func getNoZeroIntegers(n int) []int

import "fmt"

func main() {
	fmt.Println(ConvertIntegerToTheSumOfTwoNoZeroIntegers(2))    // [1 1]
	fmt.Println(ConvertIntegerToTheSumOfTwoNoZeroIntegers(11))   // [2 9]
	fmt.Println(ConvertIntegerToTheSumOfTwoNoZeroIntegers(1010)) // [122 888]
}

// Time: O(n log n), Space: O(1)
func ConvertIntegerToTheSumOfTwoNoZeroIntegers(n int) []int {
	for a := 1; a < n; a++ {
		b := n - a
		if !hasZero(a) && !hasZero(b) {
			return []int{a, b}
		}
	}
	return []int{}
}

func hasZero(x int) bool {
	for x > 0 {
		if x%10 == 0 {
			return true
		}
		x /= 10
	}
	return false
}
```
