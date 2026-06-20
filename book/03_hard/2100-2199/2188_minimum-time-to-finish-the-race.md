# 2188 — Minimum Time To Finish The Race

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func minimumFinishTime(tires [][]int, changeTime int, numLaps int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Dynamic Programming

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **Dynamic Programming** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2188: Minimum Time to Finish the Race
// https://leetcode.com/problems/minimum-time-to-finish-the-race/
// Difficulty: Hard
//
// DP: best[l] = min time for l consecutive laps with ONE tire (no change).
// dp[k] = min over j of dp[k-j] + changeTime + best[j].

import "fmt"

func main() {
	fmt.Println(minimumFinishTime([][]int{{2, 3}, {3, 4}}, 5, 4))            // 21
	fmt.Println(minimumFinishTime([][]int{{1, 10}, {2, 2}, {3, 4}}, 2, 5))   // 13
	fmt.Println(minimumFinishTime([][]int{{3, 4}}, 2, 3))                    // 13
}

func minimumFinishTime(tires [][]int, changeTime int, numLaps int) int {
	const INF = 1 << 60

  // Alokasi slice integer
	best := make([]int, numLaps+1)
  // Range loop: iterasi dengan indeks + nilai
	for i := range best {
		best[i] = INF
	}

	for _, tire := range tires {
		f, r := tire[0], tire[1]
		total := 0
		lapTime := f
		for l := 1; l <= numLaps; l++ {
			if total+lapTime >= INF {
				break
			}
			total += lapTime
			if total < best[l] {
				best[l] = total
					}
			if int64(lapTime)*int64(r) >= INF {
				break
			}
			lapTime *= r
		}
	}

  // Alokasi slice integer
	dp := make([]int, numLaps+1)
  // Range loop: iterasi dengan indeks + nilai
	for i := range dp {
		dp[i] = INF
	}
	dp[0] = -changeTime

	for i := 1; i <= numLaps; i++ {
		for j := 1; j <= i; j++ {
			if best[j] >= INF {
				continue
			}
			cand := dp[i-j] + changeTime + best[j]
			if cand < dp[i] {
				dp[i] = cand
			}
		}
	}
	return dp[numLaps]
}

func MinimumTimeToFinishTheRace() any {
	return minimumFinishTime([][]int{{2, 3}, {3, 4}}, 5, 4)
}
```
