# 1986 — Minimum Number Of Work Sessions To Finish The Tasks

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func MinimumNumberOfWorkSessionsToFinishTheTasks(tasks []int, sessionTime int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Dynamic Programming, Bitmask

**Kompleksitas Waktu:** O(2^n * n), Space: O(2^n)  
**Kompleksitas Ruang:** O(2^n)

> **Untuk fresh graduate:** Kuasai dulu teknik **Dynamic Programming** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1986: Minimum Number of Work Sessions to Finish the Tasks
// https://leetcode.com/problems/minimum-number-of-work-sessions-to-finish-the-tasks/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(MinimumNumberOfWorkSessionsToFinishTheTasks([]int{1, 2, 3}, 3))
	fmt.Println(MinimumNumberOfWorkSessionsToFinishTheTasks([]int{3, 1, 3, 1, 1}, 8))
	fmt.Println(MinimumNumberOfWorkSessionsToFinishTheTasks([]int{1, 2, 3, 4, 5}, 15))
}

// Time: O(2^n * n), Space: O(2^n)
func MinimumNumberOfWorkSessionsToFinishTheTasks(tasks []int, sessionTime int) int {
	n := len(tasks)
	m := 1 << n
  // Alokasi slice integer
	dp := make([]int, m)
  // Alokasi slice integer
	sessions := make([]int, m)

  // Range loop: iterasi dengan indeks + nilai
	for i := range dp {
		dp[i] = n + 1
		sessions[i] = sessionTime + 1
	}
	dp[0] = 0
	sessions[0] = 0

	for mask := 0; mask < m; mask++ {
		for i := 0; i < n; i++ {
			if mask&(1<<i) == 0 {
				newMask := mask | (1 << i)
				if sessions[mask]+tasks[i] <= sessionTime {
					if dp[newMask] > dp[mask] || (dp[newMask] == dp[mask] && sessions[newMask] > sessions[mask]+tasks[i]) {
						dp[newMask] = dp[mask]
						sessions[newMask] = sessions[mask] + tasks[i]
					}
				} else {
					if dp[newMask] > dp[mask]+1 || (dp[newMask] == dp[mask]+1 && sessions[newMask] > tasks[i]) {
						dp[newMask] = dp[mask] + 1
						sessions[newMask] = tasks[i]
					}
				}
			}
		}
	}

	return dp[m-1] + 1
}
```
