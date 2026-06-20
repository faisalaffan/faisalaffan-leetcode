# 0432 — All Oone Data Structure

## Deskripsi

**Soal:** [0432. All Oone Data Structure](https://leetcode.com/problems/all-oone-data-structure/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** LIS (Longest Increasing Subsequence)

## Solusi Go

```go
package main

import "fmt"

// LeetCode #432: All O(1) Data Structure
// https://leetcode.com/problems/all-oone-data-structure/
// Difficulty: Hard
//
// Doubly linked list of frequency buckets. Each bucket holds keys with that count.
// inc, dec, getMaxKey, getMinKey all O(1).

func main() {
	allOne := Constructor()
	allOne.Inc("hello")
	allOne.Inc("hello")
	fmt.Println("MaxKey after 2x hello:", allOne.GetMaxKey())
	fmt.Println("MinKey after 2x hello:", allOne.GetMinKey())

	allOne.Inc("world")
	allOne.Inc("world")
	allOne.Inc("world")
	fmt.Println("MaxKey after 3x world:", allOne.GetMaxKey())

	allOne.Dec("world")
	allOne.Dec("world")
	fmt.Println("MaxKey after dec world 2x:", allOne.GetMaxKey())
	fmt.Println("MinKey after dec world 2x:", allOne.GetMinKey())

	allOne.Dec("hello")
	allOne.Dec("hello")
	fmt.Println("After dec hello to 0:")
	fmt.Println("MaxKey:", allOne.GetMaxKey())
	fmt.Println("MinKey:", allOne.GetMinKey())

	// Empty
	empty := Constructor()
	fmt.Println("Empty MaxKey:", empty.GetMaxKey())
	fmt.Println("Empty MinKey:", empty.GetMinKey())
}

type bucket struct {
	count int
	keys  map[string]bool
	prev  *bucket
	next  *bucket
}

type AllOne struct {
	head    *bucket
	tail    *bucket
	key2node map[string]*bucket
}

func Constructor() AllOne {
	head := &bucket{count: 0, keys: make(map[string]bool)}
	tail := &bucket{count: 0, keys: make(map[string]bool)}
	head.next = tail
	tail.prev = head
	return AllOne{head: head, tail: tail, key2node: make(map[string]*bucket)}
}

func (a *AllOne) Inc(key string) {
	cur, ok := a.key2node[key]
	if !ok {
		cur = a.head
	}
	nc := cur.count + 1
	if cur.next.count != nc {
		a.insertAfter(cur, &bucket{count: nc, keys: make(map[string]bool)})
	}
	nb := cur.next
	nb.keys[key] = true
	a.key2node[key] = nb

	if cur != a.head {
		delete(cur.keys, key)
		if len(cur.keys) == 0 {
			a.remove(cur)
		}
	}
}

func (a *AllOne) Dec(key string) {
	cur, ok := a.key2node[key]
	if !ok {
		return
	}
	nc := cur.count - 1
	if nc == 0 {
		delete(cur.keys, key)
		delete(a.key2node, key)
		if len(cur.keys) == 0 {
			a.remove(cur)
		}
		return
	}
	if cur.prev.count != nc {
		a.insertAfter(cur.prev, &bucket{count: nc, keys: make(map[string]bool)})
	}
	pb := cur.prev
	pb.keys[key] = true
	a.key2node[key] = pb

	delete(cur.keys, key)
	if len(cur.keys) == 0 {
		a.remove(cur)
	}
}

func (a *AllOne) GetMaxKey() string {
	if a.tail.prev == a.head {
		return ""
	}
	for k := range a.tail.prev.keys {
		return k
	}
	return ""
}

func (a *AllOne) GetMinKey() string {
	if a.head.next == a.tail {
		return ""
	}
	for k := range a.head.next.keys {
		return k
	}
	return ""
}

func (a *AllOne) insertAfter(prev *bucket, b *bucket) {
	b.prev = prev
	b.next = prev.next
	prev.next.prev = b
	prev.next = b
}

func (a *AllOne) remove(b *bucket) {
	b.prev.next = b.next
	b.next.prev = b.prev
}
```
