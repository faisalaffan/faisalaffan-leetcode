# 0346 — Moving Average From Data Stream

## Deskripsi

**Soal:** [0346. Moving Average From Data Stream](https://leetcode.com/problems/moving-average-from-data-stream/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(1), Space: O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** Sliding Window (jendela geser), Queue (antrian FIFO)

**Fungsi Solusi:** `func Constructor(size int) MovingAverage`

## Solusi Go

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
