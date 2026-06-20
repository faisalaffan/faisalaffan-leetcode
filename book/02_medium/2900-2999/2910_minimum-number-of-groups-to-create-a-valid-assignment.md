# 2910 — Minimum Number Of Groups To Create A Valid Assignment

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func minGroupsForValidAssignment(nums []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** O(n + m*minFreq)  
**Kompleksitas Ruang:** O(m)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2910: Minimum Number of Groups to Create a Valid Assignment
// https://leetcode.com/problems/minimum-number-of-groups-to-create-a-valid-assignment/
// Difficulty: Medium
// Time: O(n + m*minFreq) | Space: O(m)

import "fmt"

func main() {
	fmt.Println(minGroupsForValidAssignment([]int{3, 3, 3, 3, 3, 1, 1}))
	fmt.Println(minGroupsForValidAssignment([]int{10, 10, 10, 10, 10}))
	fmt.Println(minGroupsForValidAssignment([]int{1, 1, 1, 2, 2, 2}))
}

func minGroupsForValidAssignment(nums []int) int {
	cnt := map[int]int{}
	for _, x := range nums {
		cnt[x]++
	}
	k := len(nums)
	for _, v := range cnt {
		if v < k {
			k = v
		}
	}
	for ; ; k-- {
		ans := 0
		ok := true
		for _, v := range cnt {
			if v/k < v%k {
				ok = false
				break
			}
			ans += (v + k) / (k + 1)
		}
		if ok {
			return ans
		}
	}
}
```
