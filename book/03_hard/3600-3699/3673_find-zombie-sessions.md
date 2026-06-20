# 3673 — Find Zombie Sessions

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan data terstruktur dan diminta untuk mencari elemen atau pola tertentu. Tugasmu adalah menemukan posisi, jumlah, atau keberadaan elemen dengan efisien.

Seperti mencari kata di kamus — kamu tidak membaca dari halaman 1, tapi langsung ke tengah (binary search), lalu maju/mundur. Teknik pencarian yang efisien sangat penting untuk interview.

**Konsep kunci:** linear search O(n), binary search O(log n), HashMap lookup O(1).

**Fungsi yang perlu kamu implementasikan:**
```go
func findZombieSessions(sessions []Session) []Session
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3673: Find Zombie Sessions
// https://leetcode.com/problems/find-zombie-sessions/
// Difficulty: Hard (SQL / Database)
//
// Find sessions that are still active (zombie) — sessions with no
// logout event recorded after the login. Port to Go with data
// structures.

import "fmt"

type Session struct {
	SessionID int
	UserID    int
	LoginTime int
	Logout    bool // true if logout recorded
}

func main() {
	// Example 1
	sessions := []Session{
		{1, 1, 100, false},
		{2, 1, 200, true},
		{3, 2, 150, false},
	}
	fmt.Println(findZombieSessions(sessions))

	// Example 2: all logged out
	sessions2 := []Session{
		{1, 1, 100, true},
		{2, 2, 200, true},
	}
	fmt.Println(findZombieSessions(sessions2))
}

func findZombieSessions(sessions []Session) []Session {
	var zombie []Session
	for _, s := range sessions {
		if !s.Logout {
			zombie = append(zombie, s)
		}
	}

	// Sort by SessionID for determinism
  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(zombie); i++ {
		for j := i + 1; j < len(zombie); j++ {
			if zombie[i].SessionID > zombie[j].SessionID {
				zombie[i], zombie[j] = zombie[j], zombie[i]
			}
		}
	}
	return zombie
}
```
