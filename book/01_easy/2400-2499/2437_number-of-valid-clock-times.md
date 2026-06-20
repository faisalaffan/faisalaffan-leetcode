# 2437 — Number Of Valid Clock Times

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan string berisi tanda kurung: `()`, `[]`, `{}`. Tugasmu adalah memeriksa apakah string tersebut **valid** — setiap kurung buka harus ditutup oleh kurung yang sesuai dalam urutan benar.

Contoh valid: `()[]{}`, `({[]})`. Tidak valid: `(]`, `([)]`.

**Cara berpikir:** Gunakan Stack. Kurung buka → push. Kurung tutup → pop dan cek kecocokan. Di akhir, stack harus kosong.

**Fungsi Solusi:** `func NumberOfValidClockTimes(time string) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** —  |  **Ruang:** —


## 💻 Solusi Go

```go
package main

// LeetCode #2437: Number of Valid Clock Times
// https://leetcode.com/problems/number-of-valid-clock-times/
// Difficulty: Easy
// Time O(1) | Space O(1)

import "fmt"

func main() {
	fmt.Println(NumberOfValidClockTimes("?5:00")) // 2
	fmt.Println(NumberOfValidClockTimes("0?:0?")) // 100
	fmt.Println(NumberOfValidClockTimes("??:??")) // 1440
}

func NumberOfValidClockTimes(time string) int {
	ways := 1

	// Hours tens digit
	if time[0] == '?' {
		if time[1] == '?' {
			ways *= 24
		} else if time[1] <= '3' {
			ways *= 3 // 0,1,2
		} else {
			ways *= 2 // 0,1
		}
	} else if time[1] == '?' {
		if time[0] == '0' || time[0] == '1' {
			ways *= 10
		} else { // time[0] == '2'
			ways *= 4
		}
	}

	// Minutes tens digit
	if time[3] == '?' {
		ways *= 6
	}
	// Minutes ones digit
	if time[4] == '?' {
		ways *= 10
	}

	return ways
}
```
