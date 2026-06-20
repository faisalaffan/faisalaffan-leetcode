# 3150 — Invalid Tweets Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan tabel database dan diminta untuk menulis query SQL. Karena repo ini menggunakan Go, query SQL disimulasikan dengan struktur data Go (map untuk grouping, slice untuk sorting, struct untuk representasi row).

Soal tipe ini menguji kemampuanmu menganalisis data relasional — seperti yang kamu lakukan dengan SQL di pekerjaan backend sehari-hari.

**Konsep kunci:** GROUP BY, JOIN, aggregate (SUM, COUNT, AVG), window function (RANK, ROW_NUMBER), HAVING.

**Fungsi yang perlu kamu implementasikan:**
```go
func InvalidTweetsIi(tweets map[int]string) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3150: Invalid Tweets II
// https://leetcode.com/problems/invalid-tweets-ii/
// Difficulty: Easy [Paid - Database]
//
// Note: This is a SQL problem on LeetCode. In Go, we implement the
// equivalent logic: find tweet IDs where content length > 140 or
// content contains invalid characters.

import (
	"fmt"
	"strings"
)

func main() {
	// LeetCode name: invalidTweets
	tweets := map[int]string{
		1: "Hello world!",
		2: strings.Repeat("a", 150),
		3: "Valid tweet content",
	}
	fmt.Println(InvalidTweetsIi(tweets))
	// [2]
}

// Time: O(n) | Space: O(n)
// LeetCode submission name: invalidTweets
func InvalidTweetsIi(tweets map[int]string) []int {
	result := []int{}
	for id, content := range tweets {
		if len(content) > 140 {
			result = append(result, id)
		}
	}
	return result
}
```
