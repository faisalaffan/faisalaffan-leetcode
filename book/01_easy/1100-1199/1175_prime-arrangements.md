# 1175 — Prime Arrangements

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func numPrimeArrangements(n int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n log log n)  |  **Ruang:** O(n)


## 💻 Solusi Go

```go
package main

// LeetCode #1175: Prime Arrangements
// https://leetcode.com/problems/prime-arrangements/
// Difficulty: Easy
// Time: O(n log log n) | Space: O(n)

import "fmt"

const mod = 1_000_000_007

func main() {
	fmt.Println(numPrimeArrangements(5))  // 12
	fmt.Println(numPrimeArrangements(100)) // 682289015
}

// LeetCode submission: numPrimeArrangements
func numPrimeArrangements(n int) int {
	primeCount := countPrimes(n)
	nonPrimeCount := n - primeCount
	ans := 1
	for i := 1; i <= primeCount; i++ {
		ans = (ans * i) % mod
	}
	for i := 1; i <= nonPrimeCount; i++ {
		ans = (ans * i) % mod
	}
	return ans
}

func countPrimes(n int) int {
	isPrime := make([]bool, n+1)
	for i := 2; i <= n; i++ {
		isPrime[i] = true
	}
	for i := 2; i*i <= n; i++ {
		if isPrime[i] {
			for j := i * i; j <= n; j += i {
				isPrime[j] = false
			}
		}
	}
	count := 0
	for i := 2; i <= n; i++ {
		if isPrime[i] {
			count++
		}
	}
	return count
}
```
