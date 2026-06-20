# 0470 — Implement Rand10 Using Rand7

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diminta untuk mendesain (merancang) sebuah struktur data kustom dengan operasi tertentu (insert, delete, search, update). Tugasmu adalah memilih representasi data yang tepat agar setiap operasi berjalan efisien — biasanya O(1) atau O(log n).

Ini adalah soal yang paling sering muncul di interview sistem desain. Kamu perlu memilih kombinasi struktur data yang tepat (HashMap + Heap + LinkedList) untuk mencapai kompleksitas yang diminta.

**Konsep kunci:** HashMap (O(1) lookup), Heap (priority), Doubly Linked List (O(1) remove), TreeMap (ordered keys).

**Fungsi yang perlu kamu implementasikan:**
```go
func ImplementRandOneZeroUsingRandSeven() int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(1) expected  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #470: Implement Rand10() Using Rand7()
// https://leetcode.com/problems/implement-rand10-using-rand7/
// Difficulty: Medium
// Time: O(1) expected
// Space: O(1)

import (
	"fmt"
	"math/rand"
)

func main() {
	// Test by generating a few random values
	for i := 0; i < 5; i++ {
		fmt.Println(ImplementRandOneZeroUsingRandSeven())
	}
}

func ImplementRandOneZeroUsingRandSeven() int {
	// Rejection sampling: (rand7()-1)*7 + rand7() gives 1-49 uniformly
	// Accept values 1-40 and map to 1-10
	for {
		val := (rand7()-1)*7 + rand7()
		if val <= 40 {
			return (val-1)%10 + 1
		}
	}
}

func rand7() int {
	return rand.Intn(7) + 1
}
```
