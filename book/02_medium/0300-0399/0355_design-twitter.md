# 0355 — Design Twitter

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diminta untuk mendesain (merancang) sebuah struktur data kustom dengan operasi tertentu (insert, delete, search, update). Tugasmu adalah memilih representasi data yang tepat agar setiap operasi berjalan efisien — biasanya O(1) atau O(log n).

Ini adalah soal yang paling sering muncul di interview sistem desain. Kamu perlu memilih kombinasi struktur data yang tepat (HashMap + Heap + LinkedList) untuk mencapai kompleksitas yang diminta.

**Konsep kunci:** HashMap (O(1) lookup), Heap (priority), Doubly Linked List (O(1) remove), TreeMap (ordered keys).

**Fungsi yang perlu kamu implementasikan:**
```go
func Constructor() Twitter
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap, Heap / Priority Queue, Stack

**Kompleksitas Waktu:** O(n log n) for newsFeed  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #355: Design Twitter
// https://leetcode.com/problems/design-twitter/
// Difficulty: Medium
// Time: O(n log n) for newsFeed | Space: O(n)

import (
	"container/heap"
	"fmt"
)

type Tweet struct {
	id   int
	time int
	next *Tweet
}

type Twitter struct {
	users map[int]*User
	time  int
}

type User struct {
	id       int
	tweets   *Tweet
	followee map[int]bool
}

type tweetHeapItem struct {
	tweet *Tweet
}

type tweetMaxHeap []*tweetHeapItem

func (h tweetMaxHeap) Len() int           { return len(h) }
func (h tweetMaxHeap) Less(i, j int) bool { return h[i].tweet.time > h[j].tweet.time }
func (h tweetMaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *tweetMaxHeap) Push(x interface{}) {
	*h = append(*h, x.(*tweetHeapItem))
}

func (h *tweetMaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func Constructor() Twitter {
	return Twitter{users: make(map[int]*User)}
}

func (tw *Twitter) getUser(id int) *User {
	if _, ok := tw.users[id]; !ok {
		tw.users[id] = &User{id: id, followee: make(map[int]bool)}
		tw.users[id].followee[id] = true
	}
	return tw.users[id]
}

func (tw *Twitter) PostTweet(userId int, tweetId int) {
	user := tw.getUser(userId)
	tw.time++
	tweet := &Tweet{id: tweetId, time: tw.time, next: user.tweets}
	user.tweets = tweet
}

func (tw *Twitter) GetNewsFeed(userId int) []int {
	user := tw.getUser(userId)

	h := &tweetMaxHeap{}
	heap.Init(h)

	for followeeId := range user.followee {
		if followee, ok := tw.users[followeeId]; ok && followee.tweets != nil {
  // Masukkan elemen ke priority queue
			heap.Push(h, &tweetHeapItem{tweet: followee.tweets})
		}
	}

  // Alokasi slice integer
	result := make([]int, 0, 10)
	for h.Len() > 0 && len(result) < 10 {
  // Ambil elemen terkecil/terbesar dari heap
		item := heap.Pop(h).(*tweetHeapItem)
		result = append(result, item.tweet.id)
		if item.tweet.next != nil {
  // Masukkan elemen ke priority queue
			heap.Push(h, &tweetHeapItem{tweet: item.tweet.next})
		}
	}
	return result
}

func (tw *Twitter) Follow(followerId int, followeeId int) {
	follower := tw.getUser(followerId)
	follower.followee[followeeId] = true
	tw.getUser(followeeId)
}

func (tw *Twitter) Unfollow(followerId int, followeeId int) {
	follower := tw.getUser(followerId)
	if followerId != followeeId {
		delete(follower.followee, followeeId)
	}
}

func main() {
	tw := Constructor()
	tw.PostTweet(1, 5)
	fmt.Println("NewsFeed 1:", tw.GetNewsFeed(1))
	// Expected: [5]

	tw.Follow(1, 2)
	tw.PostTweet(2, 6)
	fmt.Println("NewsFeed 1 after follow:", tw.GetNewsFeed(1))
	// Expected: [6, 5]

	tw.Unfollow(1, 2)
	fmt.Println("NewsFeed 1 after unfollow:", tw.GetNewsFeed(1))
	// Expected: [5]
}
```
