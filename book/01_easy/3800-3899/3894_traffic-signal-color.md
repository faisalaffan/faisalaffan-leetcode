# 3894 — Traffic Signal Color

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func TrafficSignalColor(timer int) string`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(1)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #3894: Traffic Signal Color
// https://leetcode.com/problems/traffic-signal-color/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(TrafficSignalColor(60))
	fmt.Println(TrafficSignalColor(5))
	fmt.Println(TrafficSignalColor(0))
	fmt.Println(TrafficSignalColor(30))
	fmt.Println(TrafficSignalColor(90))
}

// Time: O(1)
// Space: O(1)
func TrafficSignalColor(timer int) string {
	if timer == 0 {
		return "Green"
	}
	if timer == 30 {
		return "Orange"
	}
	if timer > 30 && timer <= 90 {
		return "Red"
	}
	return "Invalid"
}
```
