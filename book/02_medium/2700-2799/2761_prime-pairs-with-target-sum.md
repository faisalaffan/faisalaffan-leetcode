# 2761 — Prime Pairs With Target Sum

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func PrimePairsWithTargetSum(target int) [][]int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n log log n)  |  **Ruang:** O(n)


## 💻 Solusi Go

```go
package main

// LeetCode #2761: Prime Pairs With Target Sum
// https://leetcode.com/problems/prime-pairs-with-target-sum/
// Difficulty: Medium
// Time: O(n log log n) | Space: O(n)

import "fmt"

func PrimePairsWithTargetSum(target int) [][]int {
	// Sieve of Eratosthenes
	isPrime := make([]bool, target+1)
	for i := 2; i <= target; i++ {
		isPrime[i] = true
	}
	for i := 2; i*i <= target; i++ {
		if isPrime[i] {
			for j := i * i; j <= target; j += i {
				isPrime[j] = false
			}
		}
	}

  // Matriks 2D
	result := make([][]int, 0)
	for i := 2; i <= target/2; i++ {
		if isPrime[i] && isPrime[target-i] {
			result = append(result, []int{i, target - i})
		}
	}
	return result
}

func main() {
	fmt.Println(PrimePairsWithTargetSum(10))
	fmt.Println(PrimePairsWithTargetSum(2))
	fmt.Println(PrimePairsWithTargetSum(18))
}
```
