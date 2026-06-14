# Grokking Algorithms — Go Edition

My personal journey through *Grokking Algorithms* by Aditya Bhargava,
implementing every algorithm and exercise in idiomatic Go with tests.

## Progress

- [x] Ch 1 — Introduction to Algorithms (Binary Search)
- [x] Ch 2 — Selection Sort
- [x] Ch 3 — Recursion
- [x] Ch 4 — Quicksort
- [x] Ch 5 — Hash Tables
- [ ] Ch 6 — Breadth-First Search
- [ ] Ch 7 — Dijkstra's Algorithm
- [ ] Ch 8 — Greedy Algorithms
- [ ] Ch 9 — Dynamic Programming
- [ ] Ch 10 — K-Nearest Neighbors
- [ ] Ch 11 — Where to Go Next

## Running

```bash
go test ./...        # run every chapter's tests
go test ./ch01-...   # run a single chapter
```

## Structure

Each chapter is a self-contained Go package with the algorithm,
table-driven tests, and a short README explaining the idea + Big-O.
