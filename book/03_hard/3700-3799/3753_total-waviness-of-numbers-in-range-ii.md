# 3753 — Total Waviness Of Numbers In Range Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sekumpulan bilangan dan diminta untuk menghitung penjumlahan dengan aturan tertentu. Tugasmu adalah menemukan kombinasi, subset, atau urutan yang memenuhi target penjumlahan.

Seperti menghitung kembalian belanja — kamu perlu kombinasi pecahan uang yang tepat. Soal penjumlahan seringnya diselesaikan dengan HashMap (two-sum pattern) atau Prefix Sum (jumlah kumulatif).

**Konsep kunci:** target sum, complement (pelengkap), prefix sum, cumulative sum.

**Fungsi yang perlu kamu implementasikan:**
```go
func totalWaviness(a int64, b int64) int64
```

> **💡 Hint:** Digit DP. Count numbers and accumulate waviness.

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Two Pointer, Dynamic Programming

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **Two Pointer** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3753: Total Waviness of Numbers in Range II
// https://leetcode.com/problems/total-waviness-of-numbers-in-range-ii/
// Difficulty: Hard
//
// Compute total waviness of all numbers in [a, b].
// Waviness of a number = sum of absolute differences between
// consecutive digits.
//
// Approach: Digit DP. Count numbers and accumulate waviness.

import "fmt"

func main() {
	// Example 1
	fmt.Println(totalWaviness(1, 10))
	// Example 2
	fmt.Println(totalWaviness(100, 200))
	// Edge: single number
	fmt.Println(totalWaviness(123, 123))
	// Edge: single digit
	fmt.Println(totalWaviness(5, 9))
}

func totalWaviness(a int64, b int64) int64 {
	if a > b {
		return 0
	}
	return sumWavy(b) - sumWavy(a-1)
}

// sumWavy returns total waviness of all numbers in [0, n]
func sumWavy(n int64) int64 {
	if n < 0 {
		return 0
	}
	if n < 10 {
		return 0
	}

	digits := getDigits(n)
	m := len(digits)

	// cnt[pos][tight][started] = count
	// sum[pos][tight][started] = total waviness sum
  // Membuat matriks/slice 2D untuk DP
	cnt := make([][][]int64, m+1)
  // Membuat matriks/slice 2D untuk DP
	sum := make([][][]int64, m+1)
  // Range loop: iterasi dengan indeks + nilai
	for i := range cnt {
		cnt[i] = make([][]int64, 2)
		sum[i] = make([][]int64, 2)
		for j := range cnt[i] {
			cnt[i][j] = make([]int64, 2)
			sum[i][j] = make([]int64, 2)
		}
	}

	// pos from right to left
	for pos := m; pos >= 0; pos-- {
		for tight := 0; tight <= 1; tight++ {
			for started := 0; started <= 1; started++ {
				if pos == m {
					if started == 1 {
						cnt[pos][tight][started] = 1
					}
					continue
				}
				limit := 9
				if tight == 1 {
					limit = digits[pos]
				}
				for d := 0; d <= limit; d++ {
					nt := tight
					if tight == 1 && d < limit {
						nt = 0
					}
					if started == 0 && d == 0 {
						cnt[pos][tight][started] += cnt[pos+1][nt][0]
						sum[pos][tight][started] += sum[pos+1][nt][0]
					} else {
						ns := 1
						cnt[pos][tight][started] += cnt[pos+1][nt][ns]
						sum[pos][tight][started] += sum[pos+1][nt][ns]
						// Add waviness contributed by current digit
						// This needs to know the previous digit
					}
				}
			}
		}
	}

	// For accurate waviness, we need to track previous digit.
	// Use DP[pos][tight][started][lastDigit] to track both
	// count and sum simultaneously.
	type pair struct {
		cnt int64
		sum int64
	}
  // Membuat matriks/slice 2D untuk DP
	dp := make([][][][]pair, m+1)
  // Range loop: iterasi dengan indeks + nilai
	for i := range dp {
		dp[i] = make([][][]pair, 2)
		for j := range dp[i] {
			dp[i][j] = make([][]pair, 2)
			for k := range dp[i][j] {
				dp[i][j][k] = make([]pair, 11) // 10 = no prev digit
			}
		}
	}

	for pos := m; pos >= 0; pos-- {
		for tight := 0; tight <= 1; tight++ {
			for started := 0; started <= 1; started++ {
				for last := 0; last <= 10; last++ {
					if pos == m {
						if started == 1 {
							dp[pos][tight][started][last] = pair{cnt: 1, sum: 0}
						}
						continue
					}
					limit := 9
					if tight == 1 {
						limit = digits[pos]
					}
					var cp, sp int64
					for d := 0; d <= limit; d++ {
						nt := tight
						if tight == 1 && d < limit {
							nt = 0
						}
						if started == 0 && d == 0 {
							p := dp[pos+1][nt][0][10]
							cp += p.cnt
							sp += p.sum
						} else {
							ns := 1
							add := int64(0)
							if last != 10 {
								diff := d - (last % 10)
								if diff < 0 {
									diff = -diff
								}
								add = int64(diff)
							}
							p := dp[pos+1][nt][ns][d]
							cp += p.cnt
							sp += p.sum + add*p.cnt
						}
					}
					dp[pos][tight][started][last] = pair{cnt: cp, sum: sp}
				}
			}
		}
	}

	return dp[0][1][0][10].sum
}

func getDigits(n int64) []int {
  // Edge case: input kosong — langsung return
	if n == 0 {
		return []int{0}
	}
	var d []int
	for n > 0 {
		d = append(d, int(n%10))
		n /= 10
	}
	for i, j := 0, len(d)-1; i < j; i, j = i+1, j-1 {
		d[i], d[j] = d[j], d[i]
	}
	return d
}
```
