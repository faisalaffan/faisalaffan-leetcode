# 0254 — Factor Combinations

## Deskripsi

**Soal:** [0254. Factor Combinations](https://leetcode.com/problems/factor-combinations/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n log n), Space: O(log n)  
**Kompleksitas Ruang:** O(log n)

**Algoritma:** —

**Fungsi Solusi:** `func getFactors(n int) [][]int`

## Solusi Go

```go
package main

// LeetCode #254: Factor Combinations
// https://leetcode.com/problems/factor-combinations/
// Difficulty: Medium [Paid]
// Time: O(n log n), Space: O(log n)

import "fmt"

func getFactors(n int) [][]int {
	result := [][]int{}
	var backtrack func(start, remaining int, path []int)
	backtrack = func(start, remaining int, path []int) {
		if len(path) > 0 {
  // Membuat slice untuk menyimpan hasil
			combo := make([]int, len(path)+1)
			copy(combo, path)
			combo[len(combo)-1] = remaining
			result = append(result, combo)
		}

		for i := start; i*i <= remaining; i++ {
			if remaining%i == 0 {
				path = append(path, i)
				backtrack(i, remaining/i, path)
				path = path[:len(path)-1]
			}
		}
	}

	backtrack(2, n, []int{})
	return result
}

func main() {
	fmt.Println(getFactors(12))
	fmt.Println(getFactors(37))
	fmt.Println(getFactors(32))
}
```
