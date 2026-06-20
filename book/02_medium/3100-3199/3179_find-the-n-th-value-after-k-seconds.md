# 3179 — Find The N Th Value After K Seconds

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func valueAfterKSeconds(n int, k int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n * k)  |  **Ruang:** O(n)


## 💻 Solusi Go

```go
package main

// LeetCode #3179: Find the N-th Value After K Seconds
// https://leetcode.com/problems/find-the-n-th-value-after-k-seconds/
// Difficulty: Medium
// Time: O(n * k) | Space: O(n)

import "fmt"

func valueAfterKSeconds(n int, k int) int {
	const mod = 1_000_000_007
  // Alokasi slice
	arr := make([]int, n)
  // Range loop
	for i := range arr {
		arr[i] = 1
	}

	for s := 0; s < k; s++ {
		for i := 1; i < n; i++ {
			arr[i] = (arr[i] + arr[i-1]) % mod
		}
	}
	return arr[n-1]
}

func main() {
	fmt.Println(valueAfterKSeconds(4, 5)) // Expected: 56
	fmt.Println(valueAfterKSeconds(5, 3)) // Expected: 35
}
```
