# 3871 — Count Commas In Range Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func CountCommasInRangeIi(n int64) int64
```

> **💡 Hint:** Iterate powers of 1000 starting from 1000. For each power x <= n,

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(log N)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3871: Count Commas in Range II
// https://leetcode.com/problems/count-commas-in-range-ii/
// Difficulty: Medium
// Time: O(log N) | Space: O(1)
// Approach: Iterate powers of 1000 starting from 1000. For each power x <= n,
// add n - x + 1 (numbers that gain a comma at this magnitude).

import "fmt"

func CountCommasInRangeIi(n int64) int64 {
	var ans int64 = 0
	for x := int64(1000); x <= n; x *= 1000 {
		ans += n - x + 1
	}
	return ans
}

func main() {
	// Example 1
	fmt.Println(CountCommasInRangeIi(1002)) // Expected: 3

	// Example 2
	fmt.Println(CountCommasInRangeIi(998)) // Expected: 0

	// Extra
	fmt.Println(CountCommasInRangeIi(1000000)) // Expected: 999001 + 1 = 999002
}
```
