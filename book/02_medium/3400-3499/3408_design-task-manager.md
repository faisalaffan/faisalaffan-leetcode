# 3408 — Design Task Manager

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diminta mendesain struktur data kustom dengan operasi spesifik (insert, delete, search). Target: O(1) atau O(log n) per operasi.

**Cara berpikir:** Kombinasikan HashMap + Heap + Linked List sesuai kebutuhan.

**Fungsi Solusi:** `func Constructor(tasks [][]int) *TaskManager`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap, Heap

**Waktu:** O(log n) for operations  Space: O(n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3408: Design Task Manager
// https://leetcode.com/problems/design-task-manager/
// Difficulty: Medium
// Time: O(log n) for operations  Space: O(n)

import (
	"container/heap"
	"fmt"
)

type Task struct {
	userId   int
	taskId   int
	priority int
}

type MaxHeap []Task

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i].priority > h[j].priority || (h[i].priority == h[j].priority && h[i].taskId > h[j].taskId) }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x any) { *h = append(*h, x.(Task)) }

func (h *MaxHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

type TaskManager struct {
	tasks map[int]Task // taskId -> Task
	heap  *MaxHeap
}

func Constructor(tasks [][]int) *TaskManager {
	tm := &TaskManager{
		tasks: make(map[int]Task),
		heap:  &MaxHeap{},
	}
	heap.Init(tm.heap)
	for _, t := range tasks {
		task := Task{userId: t[0], taskId: t[1], priority: t[2]}
		tm.tasks[t[1]] = task
  // Push ke priority queue
		heap.Push(tm.heap, task)
	}
	return tm
}

func (tm *TaskManager) Add(userId, taskId, priority int) {
	task := Task{userId: userId, taskId: taskId, priority: priority}
	tm.tasks[taskId] = task
  // Push ke priority queue
	heap.Push(tm.heap, task)
}

func (tm *TaskManager) Edit(taskId, newPriority int) {
	if task, ok := tm.tasks[taskId]; ok {
		task.priority = newPriority
		tm.tasks[taskId] = task
  // Push ke priority queue
		heap.Push(tm.heap, task)
	}
}

func (tm *TaskManager) Rmv(taskId int) {
	delete(tm.tasks, taskId)
}

func (tm *TaskManager) ExecTop() int {
	for tm.heap.Len() > 0 {
  // Pop dari priority queue
		task := heap.Pop(tm.heap).(Task)
		if saved, ok := tm.tasks[task.taskId]; ok && saved.userId == task.userId && saved.priority == task.priority {
			delete(tm.tasks, task.taskId)
			return task.userId
		}
	}
	return -1
}

func main() {
	tm := Constructor([][]int{{1, 101, 10}, {2, 102, 20}, {3, 103, 15}})
	fmt.Println(tm.ExecTop()) // 2
	tm.Add(4, 104, 30)
	fmt.Println(tm.ExecTop()) // 4
	tm.Edit(102, 40)
	fmt.Println(tm.ExecTop()) // 2
	fmt.Println(tm.ExecTop()) // -1
}
```
