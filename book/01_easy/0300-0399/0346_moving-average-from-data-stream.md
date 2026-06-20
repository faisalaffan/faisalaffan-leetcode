# 0346 — Moving Average From Data Stream

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func Constructor(size int) MovingAverage`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Sliding Window, BFS

**Waktu:** O(1), Space: O(n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **Sliding Window** — sering muncul di interview!

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
