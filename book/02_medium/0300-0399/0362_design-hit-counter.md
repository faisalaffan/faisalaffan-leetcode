# 0362 — Design Hit Counter

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func Constructor() HitCounter
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(1) per hit  
**Kompleksitas Ruang:** O(1) (fixed 300 buckets)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #362: Design Hit Counter
// https://leetcode.com/problems/design-hit-counter/
// Difficulty: Medium [Paid]
// Time: O(1) per hit | O(s) per getHits | Space: O(1) (fixed 300 buckets)

import "fmt"

type HitCounter struct {
	timestamps [300]int
	hits       [300]int
}

func Constructor() HitCounter {
	return HitCounter{}
}

func (hc *HitCounter) Hit(timestamp int) {
	idx := timestamp % 300
	if hc.timestamps[idx] != timestamp {
		hc.timestamps[idx] = timestamp
		hc.hits[idx] = 1
	} else {
		hc.hits[idx]++
	}
}

func (hc *HitCounter) GetHits(timestamp int) int {
	total := 0
	for i := 0; i < 300; i++ {
		if timestamp-hc.timestamps[i] < 300 {
			total += hc.hits[i]
		}
	}
	return total
}

func main() {
	hc := Constructor()
	hc.Hit(1)
	hc.Hit(2)
	hc.Hit(3)
	fmt.Println("Hits at t=4:", hc.GetHits(4))
	// Expected: 3

	hc.Hit(300)
	fmt.Println("Hits at t=300:", hc.GetHits(300))
	// Expected: 4

	fmt.Println("Hits at t=301:", hc.GetHits(301))
	// Expected: 3

	fmt.Println("Hits at t=302:", hc.GetHits(302))
	// Expected: 2
}
```
