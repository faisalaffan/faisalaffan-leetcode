# 3087 — Find Trending Hashtags

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan data terstruktur dan diminta untuk mencari elemen atau pola tertentu. Tugasmu adalah menemukan posisi, jumlah, atau keberadaan elemen dengan efisien.

Seperti mencari kata di kamus — kamu tidak membaca dari halaman 1, tapi langsung ke tengah (binary search), lalu maju/mundur. Teknik pencarian yang efisien sangat penting untuk interview.

**Konsep kunci:** linear search O(n), binary search O(log n), HashMap lookup O(1).

**Fungsi yang perlu kamu implementasikan:**
```go
func findTrendingHashtags(tweets []string) []string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Prefix Sum

**Kompleksitas Waktu:** O(n * m)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **Prefix Sum** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3087: Find Trending Hashtags
// https://leetcode.com/problems/find-trending-hashtags/
// Difficulty: Medium [Paid]
// Time: O(n * m) | Space: O(n)

import (
	"fmt"
	"sort"
	"strings"
)

func findTrendingHashtags(tweets []string) []string {
  // Membuat map (HashMap) — pencarian O(1)
	freq := make(map[string]int)
	for _, tweet := range tweets {
		words := strings.Fields(tweet)
		for _, w := range words {
			if strings.HasPrefix(w, "#") {
				hashtag := strings.ToLower(w)
				freq[hashtag]++
			}
		}
	}

	type ht struct {
		tag string
		cnt int
	}
	var list []ht
	for tag, cnt := range freq {
		list = append(list, ht{tag, cnt})
	}
  // Custom sort dengan comparator
	sort.Slice(list, func(i, j int) bool {
		if list[i].cnt != list[j].cnt {
			return list[i].cnt > list[j].cnt
		}
		return list[i].tag < list[j].tag
	})

	ans := make([]string, 0, min(3, len(list)))
	for i := 0; i < min(3, len(list)); i++ {
		ans = append(ans, list[i].tag)
	}
	return ans
}

func main() {
	fmt.Println(findTrendingHashtags([]string{
		"Good morning #tech #coding",
		"Loving #coding today #go",
		"#Tech is amazing #golang",
		"#coding #coding #coding",
	}))
	fmt.Println(findTrendingHashtags([]string{
		"#a #b #c",
		"#a #b",
		"#a",
	}))
}
```
