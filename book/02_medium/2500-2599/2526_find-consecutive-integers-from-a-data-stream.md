# 2526 — Find Consecutive Integers From A Data Stream

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func Constructor(value int, k int) DataStream`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(1) per call  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #2526: Find Consecutive Integers from a Data Stream
// https://leetcode.com/problems/find-consecutive-integers-from-a-data-stream/
// Difficulty: Medium
// Time: O(1) per call | Space: O(1)
// Track count of consecutive `value` seen.

import "fmt"

type DataStream struct {
	value, k, count int
}

func main() {
	ds := Constructor(4, 3)
	fmt.Println(ds.Consec(4)) // false
	fmt.Println(ds.Consec(4)) // false
	fmt.Println(ds.Consec(4)) // true
	fmt.Println(ds.Consec(3)) // false
}

func Constructor(value int, k int) DataStream {
	return DataStream{value: value, k: k}
}

func (ds *DataStream) Consec(num int) bool {
	if num == ds.value {
		ds.count++
	} else {
		ds.count = 0
	}
	return ds.count >= ds.k
}
```
