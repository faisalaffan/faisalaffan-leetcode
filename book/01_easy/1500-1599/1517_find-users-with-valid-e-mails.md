# 1517 — Find Users With Valid E Mails

## Deskripsi

**Soal:** [1517. Find Users With Valid E Mails](https://leetcode.com/problems/find-users-with-valid-e-mails/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** N/A (SQL query), Space: N/A  
**Kompleksitas Ruang:** N/A

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #1517: Find Users With Valid E-Mails
// https://leetcode.com/problems/find-users-with-valid-e-mails/
// Difficulty: Easy
//
// This is a SQL problem. The solution is the SQL query below.
// Table: Users (user_id, name, mail)

import "fmt"

func main() {
	fmt.Println(FindUsersWithValidEMails())
}

// Time: N/A (SQL query), Space: N/A
func FindUsersWithValidEMails() string {
	return `SELECT user_id, name, mail
FROM Users
WHERE mail REGEXP '^[A-Za-z][A-Za-z0-9_.-]*@leetcode\\.com$';`
}
```
