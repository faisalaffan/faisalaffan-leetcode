# 1993 — Operations On Tree

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func Constructor1993(parent []int) LockingTree
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** DFS

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **DFS** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1993: Operations on Tree
// https://leetcode.com/problems/operations-on-tree/
// Difficulty: Medium

import "fmt"

func main() {
	obj := Constructor1993([]int{-1, 0, 0, 1, 1, 2, 2})
	fmt.Println(obj.Lock(2, 2))
	fmt.Println(obj.Unlock(2, 3))
	fmt.Println(obj.Unlock(2, 2))
	fmt.Println(obj.Lock(4, 5))
	fmt.Println(obj.Upgrade(0, 1))
	fmt.Println(obj.Lock(0, 1))
}

// LockingTree struct represents the tree with lock states
type LockingTree struct {
	locked   []int
	parent   []int
	children [][]int
}

// Constructor1993 initializes the LockingTree
func Constructor1993(parent []int) LockingTree {
	n := len(parent)
  // Alokasi slice integer
	locked := make([]int, n)
  // Range loop: iterasi dengan indeks + nilai
	for i := range locked {
		locked[i] = -1
	}
  // Membuat matriks/slice 2D untuk DP
	children := make([][]int, n)
	for i := 1; i < n; i++ {
		children[parent[i]] = append(children[parent[i]], i)
	}
	return LockingTree{locked, parent, children}
}

// Lock locks the node for the given user if it is unlocked
func (this *LockingTree) Lock(num int, user int) bool {
	if this.locked[num] == -1 {
		this.locked[num] = user
		return true
	}
	return false
}

// Unlock unlocks the node if it is locked by the given user
func (this *LockingTree) Unlock(num int, user int) bool {
	if this.locked[num] == user {
		this.locked[num] = -1
		return true
	}
	return false
}

// Upgrade locks the node for the user if:
// 1. node is unlocked
// 2. at least one locked descendant exists
// 3. no locked ancestors exist
func (this *LockingTree) Upgrade(num int, user int) bool {
	if this.locked[num] != -1 {
		return false
	}
	// Check ancestors
	x := num
	for ; x != -1; x = this.parent[x] {
		if this.locked[x] != -1 {
			return false
		}
	}
	// Find and unlock locked descendants
	find := false
	var dfs func(int)
	dfs = func(x int) {
		for _, y := range this.children[x] {
			if this.locked[y] != -1 {
				find = true
				this.locked[y] = -1
			}
			dfs(y)
		}
	}
	dfs(num)
	if !find {
		return false
	}
	this.locked[num] = user
	return true
}
```
