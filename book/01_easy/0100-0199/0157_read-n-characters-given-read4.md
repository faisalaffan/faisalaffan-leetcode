# 0157 — Read N Characters Given Read4

## Deskripsi

**Soal:** [0157. Read N Characters Given Read4](https://leetcode.com/problems/read-n-characters-given-read4/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

**Fungsi Solusi:** `func read4(buf4 []byte) int`

## Solusi Go

```go
package main

// LeetCode #157: Read N Characters Given Read4
// https://leetcode.com/problems/read-n-characters-given-read4/
// Difficulty: Easy [Paid]

import "fmt"

var _buf []byte
var _pos int

func read4(buf4 []byte) int {
	n := 0
	for i := 0; i < 4 && _pos < len(_buf); i++ {
		buf4[i] = _buf[_pos]
		_pos++
		n++
	}
	return n
}

// Time: O(n) | Space: O(1)
func Read(buf []byte, n int) int {
  // Membuat slice untuk menyimpan hasil
	buf4 := make([]byte, 4)
	total := 0
	for total < n {
		count := read4(buf4)
		if count == 0 {
			break
		}
		for i := 0; i < count && total < n; i++ {
			buf[total] = buf4[i]
			total++
		}
	}
	return total
}

func main() {
	_buf = []byte("abc")
	_pos = 0
  // Membuat slice untuk menyimpan hasil
	buf := make([]byte, 10)
	fmt.Println(Read(buf, 4), string(buf[:3]))

	_buf = []byte("abcde")
	_pos = 0
  // Membuat slice untuk menyimpan hasil
	buf2 := make([]byte, 10)
	fmt.Println(Read(buf2, 5), string(buf2[:5]))
}
```
