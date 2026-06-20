# 3404 — Count Special Subsequences

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func numberOfSubsequences(nums []int) int64
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n^2) Space: O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3404: Count Special Subsequences
// https://leetcode.com/problems/count-special-subsequences/
// Difficulty: Medium
// Time: O(n^2) Space: O(n)

import "fmt"

func numberOfSubsequences(nums []int) int64 {
	n := len(nums)
	var ans int64
  // Membuat map (HashMap) — pencarian O(1)
	cnt := make(map[float64]int)

	// For each r, q = r-2. Accumulate (p,q) pairs as r increases.
	for r := 4; r < n-2; r++ {
		q := r - 2
		b := float64(nums[q])
		for _, aVal := range nums[:q-1] {
			ratio := float64(aVal) / b
			cnt[ratio]++
		}

		c := float64(nums[r])
		for _, dVal := range nums[r+2:] {
			ratio := float64(dVal) / c
			ans += int64(cnt[ratio])
		}
	}
	return ans
}

func main() {
	fmt.Println(numberOfSubsequences([]int{1, 2, 3, 4, 3, 6, 1})) // 1
	fmt.Println(numberOfSubsequences([]int{3, 4, 3, 4, 3, 4, 3, 4})) // 3
}
```
