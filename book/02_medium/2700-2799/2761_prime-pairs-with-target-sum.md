# 2761 — Prime Pairs With Target Sum

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sekumpulan bilangan dan diminta untuk menghitung penjumlahan dengan aturan tertentu. Tugasmu adalah menemukan kombinasi, subset, atau urutan yang memenuhi target penjumlahan.

Seperti menghitung kembalian belanja — kamu perlu kombinasi pecahan uang yang tepat. Soal penjumlahan seringnya diselesaikan dengan HashMap (two-sum pattern) atau Prefix Sum (jumlah kumulatif).

**Konsep kunci:** target sum, complement (pelengkap), prefix sum, cumulative sum.

**Fungsi yang perlu kamu implementasikan:**
```go
func PrimePairsWithTargetSum(target int) [][]int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n log log n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

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

  // Membuat matriks/slice 2D untuk DP
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
