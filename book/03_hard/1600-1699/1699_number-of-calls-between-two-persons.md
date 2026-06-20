# 1699 — Number Of Calls Between Two Persons

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func numberOfCalls(calls []Call) []CallPair`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1699: Number of Calls Between Two Persons
// https://leetcode.com/problems/number-of-calls-between-two-persons/
// Difficulty: Hard [Premium]

import "fmt"

type Call struct {
	FromId   int
	ToId     int
	Duration int
}

type CallPair struct {
	Person1       int
	Person2       int
	CallCount     int
	TotalDuration int
}

func main() {
	calls1 := []Call{
		{1, 2, 59},
		{2, 1, 11},
		{1, 3, 20},
		{3, 1, 10},
	}
	result1 := numberOfCalls(calls1)
	fmt.Println("Test 1 - Call pairs:")
	for _, r := range result1 {
		fmt.Printf("  Person1: %d, Person2: %d, CallCount: %d, TotalDuration: %d\n",
			r.Person1, r.Person2, r.CallCount, r.TotalDuration)
	}
	fmt.Println("Expected: (1,2,2,70) and (1,3,2,30)\n")

	calls2 := []Call{
		{10, 20, 30},
		{20, 10, 40},
	}
	result2 := numberOfCalls(calls2)
	fmt.Println("Test 2 - Call pairs:")
	for _, r := range result2 {
		fmt.Printf("  Person1: %d, Person2: %d, CallCount: %d, TotalDuration: %d\n",
			r.Person1, r.Person2, r.CallCount, r.TotalDuration)
	}
	fmt.Println("Expected: (10,20,2,70)")
}

func numberOfCalls(calls []Call) []CallPair {
  // HashMap: O(1) lookup
	pairMap := make(map[[2]int]*CallPair)

	for _, c := range calls {
		p1, p2 := c.FromId, c.ToId
		if p1 > p2 {
			p1, p2 = p2, p1
		}
		key := [2]int{p1, p2}
		if pair, ok := pairMap[key]; ok {
			pair.CallCount++
			pair.TotalDuration += c.Duration
		} else {
			pairMap[key] = &CallPair{
				Person1:       p1,
				Person2:       p2,
				CallCount:     1,
				TotalDuration: c.Duration,
			}
		}
	}

	result := make([]CallPair, 0, len(pairMap))
	for _, pair := range pairMap {
		result = append(result, *pair)
	}
	return result
}
```
