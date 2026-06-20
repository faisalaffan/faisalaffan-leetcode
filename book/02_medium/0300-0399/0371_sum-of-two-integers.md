# 0371 — Sum Of Two Integers

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array integer dan sebuah target. Tugasmu adalah mencari **dua angka** yang jika dijumlahkan menghasilkan target. Kembalikan **indeks** (posisi) kedua angka.

Contoh: `nums=[2,7,11,15], target=9` → `2+7=9` → `[0,1]`.

**Cara berpikir:** Gunakan HashMap. Untuk setiap angka, cek apakah `target-angka` sudah ada di map. Kalau sudah → ketemu pasangan. Kalau belum → simpan angka ke map.

**Fungsi Solusi:** `func getSum(a int, b int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(1)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #371: Sum of Two Integers
// https://leetcode.com/problems/sum-of-two-integers/
// Difficulty: Medium
// Time: O(1) | Space: O(1)

import "fmt"

func getSum(a int, b int) int {
	for b != 0 {
		carry := a & b
		a = a ^ b
		b = carry << 1
	}
	return a
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", getSum(1, 2))
	// Expected: 3

	// Test case 2
	fmt.Println("Test 2:", getSum(2, 3))
	// Expected: 5

	// Test case 3: Negative numbers
	fmt.Println("Test 3:", getSum(-1, 1))
	// Expected: 0
}
```
