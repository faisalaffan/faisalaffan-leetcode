# 1152 — Analyze User Website Visit Pattern

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func mostVisitedPattern(username []string, timestamp []int, website []string) []string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n^3) worst, but constrained by problem input size  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

import (
	"fmt"
	"sort"
	"strings"
)

// LeetCode #1152: Analyze User Website Visit Pattern
// https://leetcode.com/problems/analyze-user-website-visit-pattern/
// Difficulty: Medium [Paid]

// Given usernames, timestamps, and websites, find the most visited
// 3-sequence pattern (3 websites in order). Ties broken by lexicographic order.

// Time: O(n^3) worst, but constrained by problem input size
// Space: O(n)

func mostVisitedPattern(username []string, timestamp []int, website []string) []string {
	n := len(username)

	// Group visits by user
  // Membuat map (HashMap) — pencarian O(1)
	userVisits := make(map[string][]visit)
	for i := 0; i < n; i++ {
		userVisits[username[i]] = append(userVisits[username[i]], visit{timestamp[i], website[i]})
	}

	// Sort each user's visits by timestamp
	for user := range userVisits {
  // Custom sort dengan comparator
		sort.Slice(userVisits[user], func(i, j int) bool {
			return userVisits[user][i].time < userVisits[user][j].time
		})
	}

	// Count patterns across users
  // Membuat map (HashMap) — pencarian O(1)
	patternCount := make(map[string]int)

	for _, visits := range userVisits {
		if len(visits) < 3 {
			continue
		}
  // Membuat map (HashMap) — pencarian O(1)
		userPatterns := make(map[string]bool)
		// Generate all 3-sequences for this user
  // Loop linear O(n): iterasi setiap elemen
		for i := 0; i < len(visits); i++ {
			for j := i + 1; j < len(visits); j++ {
				for k := j + 1; k < len(visits); k++ {
					pat := visits[i].site + "," + visits[j].site + "," + visits[k].site
					userPatterns[pat] = true
				}
			}
		}
		for pat := range userPatterns {
			patternCount[pat]++
		}
	}

	// Find best pattern
	bestPat := ""
	bestCount := 0
	for pat, count := range patternCount {
		if count > bestCount || (count == bestCount && (bestPat == "" || pat < bestPat)) {
			bestCount = count
			bestPat = pat
		}
	}

	return strings.Split(bestPat, ",")
}

type visit struct {
	time int
	site string
}

func main() {
	username := []string{"joe", "joe", "joe", "james", "james", "james", "james", "mary", "mary", "mary"}
	timestamp := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	website := []string{"home", "about", "career", "home", "cart", "maps", "home", "home", "about", "career"}
	fmt.Printf("%v (expected: [\"home\" \"about\" \"career\"])\n", mostVisitedPattern(username, timestamp, website))

	username2 := []string{"u1", "u1", "u1", "u2", "u2", "u2"}
	timestamp2 := []int{1, 2, 3, 4, 5, 6}
	website2 := []string{"a", "b", "c", "a", "b", "c"}
	fmt.Printf("%v (expected: [\"a\" \"b\" \"c\"])\n", mostVisitedPattern(username2, timestamp2, website2))
}
```
