# 2534 — Time Taken To Cross The Door

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func timeTakenCrossDoor(arrival []int, state []int) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** BFS

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **BFS** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2534: Time Taken to Cross the Door
// https://leetcode.com/problems/time-taken-to-cross-the-door/
// Difficulty: Hard [Paid]

import "fmt"

// timeTakenCrossDoor returns the time each person crosses the door.
// arrival[i] = time person i arrives, state[i] = 1 for enter, 0 for leave.
//
// Simulation with two queues (enterers, leavers). At each time step:
// 1. Add all people who arrived at/before current time to appropriate queue.
// 2. If both queues empty, jump to next arrival time.
// 3. Decide who crosses:
//    - One queue empty -> the other goes.
//    - Both non-empty -> prefer same direction as previous crossing.
//    - First crossing (no previous) -> enterers have priority (state=1).
// 4. Person crosses at current time, advance time by 1.
//
// Complexity: O(T + n) where T = max(arrival) + n, O(n) space
func timeTakenCrossDoor(arrival []int, state []int) []int {
	n := len(arrival)
  // Alokasi slice integer
	result := make([]int, n)

  // Alokasi slice integer
	enterQ := make([]int, 0) // indices of people waiting to enter
  // Alokasi slice integer
	leaveQ := make([]int, 0) // indices of people waiting to leave

	nextIdx := 0 // next person index to consider
	time := 0
	prevState := 1 // 1 = enter, 0 = leave; start assuming enter
	processed := 0

	for processed < n {
		// Add all people who have arrived up to current time
		for nextIdx < n && arrival[nextIdx] <= time {
			if state[nextIdx] == 1 {
				enterQ = append(enterQ, nextIdx)
			} else {
				leaveQ = append(leaveQ, nextIdx)
			}
			nextIdx++
		}

		// If no one is waiting, jump to next arrival
		if len(enterQ) == 0 && len(leaveQ) == 0 {
			if nextIdx < n {
				time = arrival[nextIdx]
				continue
			}
			break
		}

		// Decide who crosses
		var chosen int
		if len(enterQ) > 0 && len(leaveQ) > 0 {
			// Both non-empty: prefer same direction as previous crossing
			if prevState == 1 {
				chosen = enterQ[0]
				enterQ = enterQ[1:]
			} else {
				chosen = leaveQ[0]
				leaveQ = leaveQ[1:]
			}
		} else if len(enterQ) > 0 {
			chosen = enterQ[0]
			enterQ = enterQ[1:]
		} else {
			chosen = leaveQ[0]
			leaveQ = leaveQ[1:]
		}

		result[chosen] = time
		prevState = state[chosen]
		time++
		processed++
	}

	return result
}

func main() {
	// Test cases
	fmt.Println("Test 1: arrival=[0,0,0], state=[1,0,0] ->", timeTakenCrossDoor([]int{0, 0, 0}, []int{1, 0, 0}))
	fmt.Println("Test 2: arrival=[0,1,1,2,4], state=[0,1,0,0,1] ->", timeTakenCrossDoor([]int{0, 1, 1, 2, 4}, []int{0, 1, 0, 0, 1}))
	fmt.Println("Test 3: arrival=[0,0,1], state=[1,1,0] ->", timeTakenCrossDoor([]int{0, 0, 1}, []int{1, 1, 0}))
	fmt.Println("Test 4: arrival=[0], state=[0] ->", timeTakenCrossDoor([]int{0}, []int{0}))
	fmt.Println("Test 5: arrival=[0,2,3], state=[1,0,1] ->", timeTakenCrossDoor([]int{0, 2, 3}, []int{1, 0, 1}))
	fmt.Println("Test 6: arrival=[1,1,1,1], state=[0,0,1,1] ->", timeTakenCrossDoor([]int{1, 1, 1, 1}, []int{0, 0, 1, 1}))
}
```
