# 0197 — Rising Temperature

## Deskripsi

**Soal:** [0197. Rising Temperature](https://leetcode.com/problems/rising-temperature/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** —

**Fungsi Solusi:** `func RisingTemperature() string`

## Solusi Go

```go
package main

// LeetCode #197: Rising Temperature
// https://leetcode.com/problems/rising-temperature/
// Difficulty: Easy

import "fmt"

func RisingTemperature() string {
	return "SELECT w1.id FROM Weather w1 JOIN Weather w2 ON DATEDIFF(w1.recordDate, w2.recordDate) = 1 WHERE w1.temperature > w2.temperature"
}

func main() {
	fmt.Println(RisingTemperature())
}
```
