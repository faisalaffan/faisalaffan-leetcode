# 2671 — Frequency Tracker

## Deskripsi

**Soal:** [2671. Frequency Tracker](https://leetcode.com/problems/frequency-tracker/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(1) per operation  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

**Fungsi Solusi:** `func Constructor() FrequencyTracker`

## Solusi Go

```go
package main

// LeetCode #2671: Frequency Tracker
// https://leetcode.com/problems/frequency-tracker/
// Difficulty: Medium
// Time: O(1) per operation | Space: O(n)

import "fmt"

type FrequencyTracker struct {
	freq   map[int]int
	freqOf map[int]int
}

func Constructor() FrequencyTracker {
	return FrequencyTracker{
		freq:   make(map[int]int),
		freqOf: make(map[int]int),
	}
}

func (this *FrequencyTracker) Add(number int) {
	oldFreq := this.freq[number]
	if this.freqOf[oldFreq] > 0 {
		this.freqOf[oldFreq]--
	}
	this.freq[number]++
	this.freqOf[oldFreq+1]++
}

func (this *FrequencyTracker) DeleteOne(number int) {
	if this.freq[number] == 0 {
		return
	}
	oldFreq := this.freq[number]
	this.freqOf[oldFreq]--
	this.freq[number]--
	if this.freq[number] > 0 {
		this.freqOf[oldFreq-1]++
	}
}

func (this *FrequencyTracker) HasFrequency(frequency int) bool {
	return this.freqOf[frequency] > 0
}

func main() {
	ft := Constructor()

	// Test case 1
	ft.Add(3)
	ft.Add(3)
	fmt.Println("Test 1:", ft.HasFrequency(2))
	// Expected: true

	// Test case 2
	ft.DeleteOne(3)
	fmt.Println("Test 2:", ft.HasFrequency(1))
	// Expected: true

	// Test case 3
	fmt.Println("Test 3:", ft.HasFrequency(2))
	// Expected: false
}
```
