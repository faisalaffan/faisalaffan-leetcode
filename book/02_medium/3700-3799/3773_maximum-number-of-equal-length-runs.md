# 3773 — Maximum Number Of Equal Length Runs

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func maximumNumberOfEqualLengthRuns(s string) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3773: Maximum Number of Equal Length Runs
// https://leetcode.com/problems/maximum-number-of-equal-length-runs/
// Difficulty: Medium [Paid]
// Time: O(n) | Space: O(n)

import "fmt"

func maximumNumberOfEqualLengthRuns(s string) int {
  // Membuat map (HashMap) — pencarian O(1)
	cnt := make(map[int]int)
	maxCount := 0
	n := len(s)
	i := 0
	for i < n {
		j := i
		for j < n && s[j] == s[i] {
			j++
		}
		runLen := j - i
		cnt[runLen]++
		if cnt[runLen] > maxCount {
			maxCount = cnt[runLen]
		}
		i = j
	}
	return maxCount
}

func main() {
	fmt.Println(maximumNumberOfEqualLengthRuns("hello"))
	fmt.Println(maximumNumberOfEqualLengthRuns("aaabaaa"))
	fmt.Println(maximumNumberOfEqualLengthRuns("aabbcc"))
}
```
