# 2398 — Maximum Number Of Robots Within Budget

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func maximumRobots(chargeTimes []int, runningCosts []int, budget int64) int
```

> **💡 Hint:** Sliding window with monotonic deque. The budget constraint for

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Two Pointer, Sliding Window, Monotonic Stack/Queue

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **Two Pointer** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2398: Maximum Number of Robots Within Budget
// https://leetcode.com/problems/maximum-number-of-robots-within-budget/
// Difficulty: Hard
//
// Approach: Sliding window with monotonic deque. The budget constraint for
// subarray [l..r] is: max(chargeTimes[l..r]) + (r-l+1) * sum(runningCosts[l..r]) <= budget.
// Maintain a deque for max chargeTime in current window. Expand right pointer,
// shrink left while over budget. Track max window length.

import "fmt"

func main() {
	// Example 1: chargeTimes=[3,6,1,3,4], runningCosts=[2,1,3,4,5], budget=25 => 3
	fmt.Println(maximumRobots([]int{3, 6, 1, 3, 4}, []int{2, 1, 3, 4, 5}, 25))
	// Example 2: chargeTimes=[11,12,19], runningCosts=[10,8,7], budget=19 => 0
	fmt.Println(maximumRobots([]int{11, 12, 19}, []int{10, 8, 7}, 19))
	// Edge: single robot within budget
	fmt.Println(maximumRobots([]int{1}, []int{1}, 2))
	// Edge: single robot over budget
	fmt.Println(maximumRobots([]int{10}, []int{10}, 1))
	// Edge: all within budget
	fmt.Println(maximumRobots([]int{1, 2, 3}, []int{1, 1, 1}, 100))
}

func maximumRobots(chargeTimes []int, runningCosts []int, budget int64) int {
	n := len(chargeTimes)
  // Alokasi slice integer
	deque := make([]int, 0) // stores indices, front = max chargeTime
	left := 0
	var sum int64
	maxLen := 0

	for right := 0; right < n; right++ {
		sum += int64(runningCosts[right])

		// Maintain deque: remove smaller elements from back
		for len(deque) > 0 && chargeTimes[deque[len(deque)-1]] <= chargeTimes[right] {
			deque = deque[:len(deque)-1]
		}
		deque = append(deque, right)

		// Check budget constraint
		for left <= right {
			maxCharge := chargeTimes[deque[0]]
			cost := int64(maxCharge) + int64(right-left+1)*sum
			if cost <= budget {
				break
			}
			// Shrink left
			if deque[0] == left {
				deque = deque[1:]
			}
			sum -= int64(runningCosts[left])
			left++
		}

		if right-left+1 > maxLen {
			maxLen = right - left + 1
		}
	}
	return maxLen
}
```
