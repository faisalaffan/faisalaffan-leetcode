# 2020 — Number Of Accounts That Did Not Stream

## Deskripsi

**Soal:** [2020. Number Of Accounts That Did Not Stream](https://leetcode.com/problems/number-of-accounts-that-did-not-stream/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n + m)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

**Fungsi Solusi:** `func numberOfAccountsThatDidNotStream(subscriptions [][]int, streams [][]int) int`

## Solusi Go

```go
package main

// LeetCode #2020: Number of Accounts That Did Not Stream
// https://leetcode.com/problems/number-of-accounts-that-did-not-stream/
// Difficulty: Medium [Paid]
// Time: O(n + m) | Space: O(n)

import "fmt"

func numberOfAccountsThatDidNotStream(subscriptions [][]int, streams [][]int) int {
	if len(subscriptions) == 0 {
		return 0
	}

	// subscriptions[i] = [start_i, end_i] (account i's subscription period)
	// streams[j] = [date_j, account_j] (stream event)
  // Membuat slice untuk menyimpan hasil
	hasStreamed := make([]bool, len(subscriptions))

	for _, s := range streams {
		accID := s[1]
		date := s[0]
		if accID >= 0 && accID < len(subscriptions) {
			if date >= subscriptions[accID][0] && date <= subscriptions[accID][1] {
				hasStreamed[accID] = true
			}
		}
	}

	count := 0
	for _, v := range hasStreamed {
		if !v {
			count++
		}
	}
	return count
}

func main() {
	// Test case 1: Account 1 never streamed
	subs1 := [][]int{{0, 180}, {60, 365}, {180, 365}}
	streams1 := [][]int{{15, 0}, {150, 0}, {300, 2}}
	fmt.Println("Test 1:", numberOfAccountsThatDidNotStream(subs1, streams1))
	// Expected: 1

	// Test case 2: All streamed
	subs2 := [][]int{{0, 365}}
	streams2 := [][]int{{100, 0}}
	fmt.Println("Test 2:", numberOfAccountsThatDidNotStream(subs2, streams2))
	// Expected: 0

	// Test case 3: None streamed
	subs3 := [][]int{{0, 30}, {31, 60}}
	fmt.Println("Test 3:", numberOfAccountsThatDidNotStream(subs3, nil))
	// Expected: 2
}
```
