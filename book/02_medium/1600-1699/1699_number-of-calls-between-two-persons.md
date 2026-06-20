# 1699 — Number Of Calls Between Two Persons

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func numberOfCalls(records []CallRecord) [][3]int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** O(n), Space: O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1699: Number of Calls Between Two Persons
// https://leetcode.com/problems/number-of-calls-between-two-persons/
// Difficulty: Medium [Paid] (SQL)
// This is a SQL problem. In Go we simulate the logic.
// Time: O(n), Space: O(n)

import "fmt"

type CallRecord struct {
	FromID int
	ToID   int
	Dur    int
}

func numberOfCalls(records []CallRecord) [][3]int {
  // Membuat map (HashMap) — pencarian O(1)
	callMap := make(map[[2]int]int) // [min,max] -> total duration

	for _, r := range records {
		a, b := r.FromID, r.ToID
		if a > b {
			a, b = b, a
		}
		key := [2]int{a, b}
		callMap[key] += r.Dur
	}

  // Alokasi slice integer
	result := make([][3]int, 0, len(callMap))
	for key, dur := range callMap {
		result = append(result, [3]int{key[0], key[1], dur})
	}
	return result
}

func main() {
	records := []CallRecord{
		{1, 2, 10},
		{2, 1, 20},
		{1, 3, 30},
	}
	result := numberOfCalls(records)
	for _, r := range result {
		fmt.Printf("(%d,%d): %d\n", r[0], r[1], r[2])
	}
}
```
