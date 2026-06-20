# 3697 — Compute Decimal Representation

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func ComputeDecimalRepresentation(n int) []int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(log n)  |  **Ruang:** O(log n)


## 💻 Solusi Go

```go
package main

// LeetCode #3697: Compute Decimal Representation
// https://leetcode.com/problems/compute-decimal-representation/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(ComputeDecimalRepresentation(537))
	fmt.Println(ComputeDecimalRepresentation(102))
	fmt.Println(ComputeDecimalRepresentation(6))
}

// Time: O(log n)
// Space: O(log n)
func ComputeDecimalRepresentation(n int) []int {
  // Alokasi slice
	res := make([]int, 0)
	place := 1
	for n > 0 {
		d := n % 10
		if d != 0 {
			res = append(res, d*place)
		}
		place *= 10
		n /= 10
	}

	// Reverse to descending order
	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return res
}
```
