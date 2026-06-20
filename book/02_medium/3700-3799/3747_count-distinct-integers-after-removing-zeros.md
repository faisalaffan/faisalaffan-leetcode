# 3747 — Count Distinct Integers After Removing Zeros

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func countDistinctIntegersAfterRemovingZeros(n int64) int64`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(log n)  |  **Ruang:** O(log n)


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
  // Alokasi slice
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
