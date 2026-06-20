# 0981 — Time Based Key Value Store

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func Constructor() TimeMap
```

> **💡 Hint:** Hash map of key -> sorted list of (timestamp, value) pairs + binary search

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Binary Search

**Kompleksitas Waktu:** O(1) for set, O(log n) for get  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **Binary Search** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #981: Time Based Key-Value Store
// https://leetcode.com/problems/time-based-key-value-store/
// Difficulty: Medium
//
// Approach: Hash map of key -> sorted list of (timestamp, value) pairs + binary search
// Time: O(1) for set, O(log n) for get
// Space: O(n)

import (
	"fmt"
	"sort"
)

func main() {
	tv := Constructor()
	tv.Set("foo", "bar", 1)
	fmt.Println(tv.Get("foo", 1))  // "bar"
	fmt.Println(tv.Get("foo", 3))  // "bar"
	tv.Set("foo", "bar2", 4)
	fmt.Println(tv.Get("foo", 4))  // "bar2"
	fmt.Println(tv.Get("foo", 5))  // "bar2"
	fmt.Println(tv.Get("foo", 0))  // ""
}

type pair struct {
	timestamp int
	value     string
}

type TimeMap struct {
	store map[string][]pair
}

func Constructor() TimeMap {
	return TimeMap{store: make(map[string][]pair)}
}

func (tm *TimeMap) Set(key string, value string, timestamp int) {
	tm.store[key] = append(tm.store[key], pair{timestamp, value})
}

func (tm *TimeMap) Get(key string, timestamp int) string {
	pairs, ok := tm.store[key]
	if !ok || len(pairs) == 0 {
		return ""
	}

	idx := sort.Search(len(pairs), func(i int) bool {
		return pairs[i].timestamp > timestamp
	})

	if idx == 0 {
		return ""
	}
	return pairs[idx-1].value
}
```
