# 0276 — Paint Fence

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func numWays(n int, k int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n), Space: O(1)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #276: Paint Fence
// https://leetcode.com/problems/paint-fence/
// Difficulty: Medium [Paid]
// Time: O(n), Space: O(1)

import "fmt"

func numWays(n int, k int) int {
	if n == 0 || k == 0 {
		return 0
	}
	if n == 1 {
		return k
	}

	same := k
	diff := k * (k - 1)

	for i := 3; i <= n; i++ {
		same, diff = diff, (same+diff)*(k-1)
	}

	return same + diff
}

func main() {
	fmt.Println(numWays(3, 2))
	fmt.Println(numWays(1, 1))
	fmt.Println(numWays(7, 2))
}
```
