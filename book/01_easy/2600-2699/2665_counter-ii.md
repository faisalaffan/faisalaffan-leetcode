# 2665 — Counter Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func CounterIi(init int) *Counter`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(1)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #2665: Counter II
// https://leetcode.com/problems/counter-ii/
// Difficulty: Easy
// Time: O(1) | Space: O(1)
// Note: JavaScript problem, adapted to Go. Returns an object with increment/decrement/reset.

import "fmt"

func main() {
	counter := CounterIi(5)
	fmt.Println(counter.increment())
	fmt.Println(counter.reset())
	fmt.Println(counter.decrement())
}

type Counter struct {
	init  int
	value int
}

func (c *Counter) increment() int {
	c.value++
	return c.value
}

func (c *Counter) decrement() int {
	c.value--
	return c.value
}

func (c *Counter) reset() int {
	c.value = c.init
	return c.value
}

func CounterIi(init int) *Counter {
	return &Counter{init: init, value: init}
}
```
