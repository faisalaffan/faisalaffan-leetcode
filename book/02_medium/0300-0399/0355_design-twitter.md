# 0355 — Design Twitter

## Deskripsi

**Soal:** [0355. Design Twitter](https://leetcode.com/problems/design-twitter/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n log n) for newsFeed  
**Kompleksitas Ruang:** O(n)

**Algoritma:** Heap (priority queue)

**Fungsi Solusi:** `func Constructor() Twitter`

## Solusi Go

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
			heap.Push(h, &tweetHeapItem{tweet: followee.tweets})
		}
	}

  // Membuat slice untuk menyimpan hasil
	result := make([]int, 0, 10)
	for h.Len() > 0 && len(result) < 10 {
		item := heap.Pop(h).(*tweetHeapItem)
		result = append(result, item.tweet.id)
		if item.tweet.next != nil {
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
