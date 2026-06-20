# 3747 — Count Distinct Integers After Removing Zeros

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func countDistinctIntegersAfterRemovingZeros(n int64) int64
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(log n)  
**Kompleksitas Ruang:** O(log n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3747: Count Distinct Integers After Removing Zeros
// https://leetcode.com/problems/count-distinct-integers-after-removing-zeros/
// Difficulty: Medium
// Time: O(log n) | Space: O(log n)

import "fmt"

func countDistinctIntegersAfterRemovingZeros(n int64) int64 {
	s := fmt.Sprintf("%d", n)
	m := len(s)

	// Precompute powers of 9
  // Alokasi slice integer
	pow9 := make([]int64, m+1)
	pow9[0] = 1
	for i := 1; i <= m; i++ {
		pow9[i] = pow9[i-1] * 9
	}

	// Count numbers with fewer digits (all non-zero digits)
	var ans int64
	for length := 1; length < m; length++ {
		ans += pow9[length]
	}

	// Count numbers with same length as n, but <= n
	for idx := 0; idx < m; idx++ {
		d := int(s[idx] - '0')
		if d == 0 {
			return ans
		}
		ans += int64(d-1) * pow9[m-idx-1]
	}
	return ans + 1
}

func main() {
	fmt.Println(countDistinctIntegersAfterRemovingZeros(10))
	fmt.Println(countDistinctIntegersAfterRemovingZeros(100))
	fmt.Println(countDistinctIntegersAfterRemovingZeros(1))
}
```
