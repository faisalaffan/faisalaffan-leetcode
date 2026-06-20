# 0597 — Friend Requests I Overall Acceptance Rate

## Deskripsi

**Soal:** [0597. Friend Requests I Overall Acceptance Rate](https://leetcode.com/problems/friend-requests-i-overall-acceptance-rate/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** —

**Fungsi Solusi:** `func FriendRequestsIOverallAcceptanceRate() string`

## Solusi Go

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
