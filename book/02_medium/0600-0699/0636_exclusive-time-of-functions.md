# 0636 — Exclusive Time Of Functions

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func exclusiveTime(n int, logs []string) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Stack

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **Stack** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #636: Exclusive Time of Functions
// https://leetcode.com/problems/exclusive-time-of-functions/
// Difficulty: Medium

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(exclusiveTime(2, []string{"0:start:0", "1:start:2", "1:end:5", "0:end:6"}))
	fmt.Println(exclusiveTime(1, []string{"0:start:0", "0:start:2", "0:end:5", "0:start:6", "0:end:6", "0:end:7"}))
	fmt.Println(exclusiveTime(2, []string{"0:start:0", "0:start:2", "0:end:5", "1:start:6", "1:end:6", "0:end:7"}))
}

func exclusiveTime(n int, logs []string) []int {
  // Alokasi slice integer
	result := make([]int, n)
  // Alokasi slice integer
	stack := make([]int, 0)
	prevTime := 0

	for _, log := range logs {
		parts := strings.Split(log, ":")
		id, _ := strconv.Atoi(parts[0])
		typ := parts[1]
		timestamp, _ := strconv.Atoi(parts[2])

		if typ == "start" {
			if len(stack) > 0 {
				result[stack[len(stack)-1]] += timestamp - prevTime
			}
			stack = append(stack, id)
			prevTime = timestamp
		} else {
			result[id] += timestamp - prevTime + 1
			stack = stack[:len(stack)-1]
			prevTime = timestamp + 1
		}
	}

	return result
}
```
