# 0379 — Design Phone Directory

## Deskripsi

**Soal:** [0379. Design Phone Directory](https://leetcode.com/problems/design-phone-directory/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(1) per operation  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

**Fungsi Solusi:** `func Constructor(maxNumbers int) PhoneDirectory`

## Solusi Go

```go
package main

// LeetCode #379: Design Phone Directory
// https://leetcode.com/problems/design-phone-directory/
// Difficulty: Medium [Paid]
// Time: O(1) per operation | Space: O(n)

import "fmt"

type PhoneDirectory struct {
	available []int
	used      []bool
	idx       int
}

func Constructor(maxNumbers int) PhoneDirectory {
  // Membuat slice untuk menyimpan hasil
	available := make([]int, maxNumbers)
	for i := 0; i < maxNumbers; i++ {
		available[i] = i
	}
	return PhoneDirectory{
		available: available,
		used:      make([]bool, maxNumbers),
		idx:       0,
	}
}

func (pd *PhoneDirectory) Get() int {
	if pd.idx >= len(pd.available) {
		return -1
	}
	num := pd.available[pd.idx]
	pd.idx++
	pd.used[num] = true
	return num
}

func (pd *PhoneDirectory) Check(number int) bool {
	if number < 0 || number >= len(pd.used) {
		return false
	}
	return !pd.used[number]
}

func (pd *PhoneDirectory) Release(number int) {
	if number < 0 || number >= len(pd.used) || !pd.used[number] {
		return
	}
	pd.used[number] = false
	pd.idx--
	pd.available[pd.idx] = number
}

func main() {
	pd := Constructor(3)
	fmt.Println("Get:", pd.Get())      // 0
	fmt.Println("Get:", pd.Get())      // 1
	fmt.Println("Check 2:", pd.Check(2)) // true
	fmt.Println("Check 1:", pd.Check(1)) // false (in use)
	pd.Release(1)
	fmt.Println("Check 1 after release:", pd.Check(1)) // true
	fmt.Println("Get:", pd.Get())      // 1 (recycled)
	fmt.Println("Get:", pd.Get())      // 2
	fmt.Println("Get:", pd.Get())      // -1 (none left)
}
```
