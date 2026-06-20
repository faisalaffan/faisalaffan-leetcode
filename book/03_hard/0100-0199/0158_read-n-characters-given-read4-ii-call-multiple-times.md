# 0158 — Read N Characters Given Read4 Ii Call Multiple Times

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func read4(buf []byte) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

import "fmt"

// LeetCode #158: Read N Characters Given Read4 II - Call Multiple Times
// https://leetcode.com/problems/read-n-characters-given-read4-ii-call-multiple-times/
// Difficulty: Hard [Paid]
//
// The key challenge: read() can be called multiple times. read4() reads up to 4 chars
// from the underlying file. Any chars read by read4() but not consumed must be buffered
// for subsequent read() calls.

// --- Simulated API -----------------------------------------------------------

var fileContent string
var filePos int

// read4 reads up to 4 characters from the simulated file into buf.
// Returns the actual number of characters read (0 when EOF).
func read4(buf []byte) int {
	n := 4
	if filePos+n > len(fileContent) {
		n = len(fileContent) - filePos
	}
	for i := 0; i < n; i++ {
		buf[i] = fileContent[filePos+i]
	}
	filePos += n
	return n
}

// --- Solution ----------------------------------------------------------------

// Solution wraps read4 with an internal buffer so that characters left over
// from a previous read4() call are used first on the next read() call.
type Solution struct {
	buf    [4]byte // internal buffer from read4
	bufPtr int     // next unconsumed position in buf
	bufCnt int     // number of valid bytes in buf (from last read4)
}

// Read reads up to n characters into buf. Returns the number of chars read.
func (s *Solution) Read(buf []byte, n int) int {
	total := 0
	for total < n {
		// Refill internal buffer when exhausted.
		if s.bufPtr >= s.bufCnt {
			s.bufCnt = read4(s.buf[:])
			s.bufPtr = 0
			if s.bufCnt == 0 {
				break // EOF
			}
		}
		// Copy from internal buffer to caller buffer.
		toCopy := min(n-total, s.bufCnt-s.bufPtr)
		for i := 0; i < toCopy; i++ {
			buf[total+i] = s.buf[s.bufPtr+i]
		}
		s.bufPtr += toCopy
		total += toCopy
	}
	return total
}

// --- Helpers -----------------------------------------------------------------

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// --- Tests -------------------------------------------------------------------

func main() {
	fmt.Println("=== 0158 Read N Characters Given Read4 II - Call Multiple Times ===")

	// Test 1: Single read consuming all content.
	fileContent = "abcde"
	filePos = 0
	sol := &Solution{}
	buf := make([]byte, 5)
	n := sol.Read(buf, 5)
	fmt.Printf("Test 1 - Single read(%d) = %d, got %q (expected 5 abcde)\n", 5, n, string(buf[:n]))

	// Test 2: Multiple sequential reads.
	fileContent = "abcdefghij"
	filePos = 0
	sol = &Solution{}
	buf1 := make([]byte, 4)
	n1 := sol.Read(buf1, 4)
	buf2 := make([]byte, 4)
	n2 := sol.Read(buf2, 4)
	buf3 := make([]byte, 4)
	n3 := sol.Read(buf3, 4)
	fmt.Printf("Test 2a - Read(4) = %d, got %q (expected 4 abcd)\n", n1, string(buf1[:n1]))
	fmt.Printf("Test 2b - Read(4) = %d, got %q (expected 4 efgh)\n", n2, string(buf2[:n2]))
	fmt.Printf("Test 2c - Read(4) = %d, got %q (expected 2 ij)\n", n3, string(buf3[:n3]))

	// Test 3: Read fewer bytes than available, then more.
	fileContent = "abcdef"
	filePos = 0
	sol = &Solution{}
	buf4 := make([]byte, 2)
	n4 := sol.Read(buf4, 2)
	fmt.Printf("Test 3a - Read(2) = %d, got %q (expected 2 ab)\n", n4, string(buf4[:n4]))
	buf5 := make([]byte, 5)
	n5 := sol.Read(buf5, 5)
	fmt.Printf("Test 3b - Read(5) = %d, got %q (expected 4 cdef)\n", n5, string(buf5[:n5]))

	// Test 4: Read zero bytes.
	fileContent = "abc"
	filePos = 0
	sol = &Solution{}
	buf6 := make([]byte, 0)
	n6 := sol.Read(buf6, 0)
	fmt.Printf("Test 4 - Read(0) = %d (expected 0)\n", n6)

	// Test 5: Empty file.
	fileContent = ""
	filePos = 0
	sol = &Solution{}
	buf7 := make([]byte, 3)
	n7 := sol.Read(buf7, 3)
	fmt.Printf("Test 5 - Read(3) on empty = %d (expected 0)\n", n7)
}
```
