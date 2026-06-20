# 0346 — Moving Average From Data Stream

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func Constructor(size int) MovingAverage
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Sliding Window, BFS

**Kompleksitas Waktu:** O(1), Space: O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **Sliding Window** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #346: Moving Average from Data Stream
// https://leetcode.com/problems/moving-average-from-data-stream/
// Difficulty: Easy [Paid]

import "fmt"

// MovingAverage maintains a sliding window average of the last size values.
type MovingAverage struct {
	size  int
	queue []int
	sum   int
}

// Constructor creates a new MovingAverage with the given window size.
func Constructor(size int) MovingAverage {
	return MovingAverage{size: size}
}

// Next adds a value and returns the moving average.
// Time: O(1), Space: O(n)
func (m *MovingAverage) Next(val int) float64 {
	m.queue = append(m.queue, val)
	m.sum += val
	if len(m.queue) > m.size {
		m.sum -= m.queue[0]
		m.queue = m.queue[1:]
	}
	return float64(m.sum) / float64(len(m.queue))
}

func main() {
	obj := Constructor(3)
	fmt.Println(obj.Next(1))
	fmt.Println(obj.Next(10))
	fmt.Println(obj.Next(3))
	fmt.Println(obj.Next(5))
}
```
