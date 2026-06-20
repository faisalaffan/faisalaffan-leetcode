# 3433 — Count Mentions Per User

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func countMentions(numberOfUsers int, events [][]string) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n log n + n*u) Space: O(n + u)  
**Kompleksitas Ruang:** O(n + u)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3433: Count Mentions Per User
// https://leetcode.com/problems/count-mentions-per-user/
// Difficulty: Medium
// Time: O(n log n + n*u) Space: O(n + u)

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func countMentions(numberOfUsers int, events [][]string) []int {
  // Custom sort dengan comparator
	sort.Slice(events, func(i, j int) bool {
		x, _ := strconv.Atoi(events[i][1])
		y, _ := strconv.Atoi(events[j][1])
		if x == y {
			return events[i][0][2] < events[j][0][2]
		}
		return x < y
	})

  // Alokasi slice integer
	ans := make([]int, numberOfUsers)
  // Alokasi slice integer
	onlineT := make([]int, numberOfUsers)
	lazy := 0

	for _, e := range events {
		etype := e[0]
		cur, _ := strconv.Atoi(e[1])
		s := e[2]

		if etype[0] == 'O' {
			userID, _ := strconv.Atoi(s)
			onlineT[userID] = cur + 60
		} else if s[0] == 'A' {
			lazy++
		} else if s[0] == 'H' {
			for i := 0; i < numberOfUsers; i++ {
				if onlineT[i] <= cur {
					ans[i]++
				}
			}
		} else {
			mentions := strings.Split(s, " ")
			for _, m := range mentions {
				userID, _ := strconv.Atoi(m[2:])
				ans[userID]++
			}
		}
	}

	if lazy > 0 {
		for i := 0; i < numberOfUsers; i++ {
			ans[i] += lazy
		}
	}
	return ans
}

func main() {
	fmt.Println(countMentions(3, [][]string{{"MESSAGE", "1", "ALL"}, {"OFFLINE", "2", "1"}, {"MESSAGE", "3", "HERE"}})) // [2, 1, 2]
	fmt.Println(countMentions(2, [][]string{{"MESSAGE", "0", "id0"}, {"MESSAGE", "1", "id1"}})) // [1, 1]
}
```
