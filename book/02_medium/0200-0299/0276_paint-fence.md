# 0276 — Paint Fence

## Deskripsi

**Soal:** [0276. Paint Fence](https://leetcode.com/problems/paint-fence/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n), Space: O(1)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

**Fungsi Solusi:** `func numWays(n int, k int) int`

## Solusi Go

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
