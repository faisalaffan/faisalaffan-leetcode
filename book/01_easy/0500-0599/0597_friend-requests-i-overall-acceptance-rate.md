# 0597 — Friend Requests I Overall Acceptance Rate

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func FriendRequestsIOverallAcceptanceRate() string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #597: Friend Requests I: Overall Acceptance Rate
// https://leetcode.com/problems/friend-requests-i-overall-acceptance-rate/
// Difficulty: Easy [Paid]

import "fmt"

func FriendRequestsIOverallAcceptanceRate() string {
	return "SELECT ROUND(IFNULL((SELECT COUNT(DISTINCT requester_id, accepter_id) FROM RequestAccepted) / (SELECT COUNT(DISTINCT sender_id, send_to_id) FROM FriendRequest), 0), 2) AS accept_rate"
}

func main() {
	fmt.Println(FriendRequestsIOverallAcceptanceRate())
}
```
