# Bina

A comprehensive data structures library for Go, providing type-safe generic
implementations of classical and modern collection types.

## Overview

Bina separates interface definitions from implementations, allowing you to
program against abstract collection types while choosing concrete
implementations based on your performance requirements. All collections use Go
generics for compile-time type safety.

### Design Principles

- **Interface-first**: Core interfaces in `core/collection/` define contracts;
  implementations in `internal/` fulfill them
- **Composable**: Collections embed common interfaces like `Collection[T]` and
  `Aggregate[T]` for consistent behavior
- **Zero dependencies**: Only relies on the Go standard library and the `opt`
  package for optional types

## v0.1 Roadmap

### Implemented

| Category    | Structure             | Interface | Implementation |
| ----------- | --------------------- | --------- | -------------- |
| Sequential  | Array                 | ✓         | ✓              |
| Sequential  | Slice                 | ✓         | ✓              |
| Sequential  | LinkedList (singly)   | ✓         | ✓              |
| Sequential  | LinkedList (doubly)   | ✓         | ✓              |
| Sequential  | LinkedList (circular) | —         | ✓              |
| Sequential  | Deque                 | ✓         | ✓              |
| Sequential  | Stack                 | ✓         | ✓              |
| Sequential  | Queue                 | ✓         | ✓              |
| Associative | HashMap               | ✓         | ✓              |
| Associative | OrderedMap            | ✓         | ✓              |
| Sets        | HashSet               | ✓         | ✓              |
| Sets        | OrderedSet            | ✓         | ✓              |
| Trees       | BTree                 | ✓         | ✓              |

### Remaining

| Category    | Structure           | Notes                        | Interface | Implemented |
| ----------- | ------------------- | ---------------------------- | --------- | ----------- |
| Sequential  | PriorityQueue       | Heap-based                   |           |             |
| Associative | MultiMap            | Multiple values per key      |           |             |
| Associative | Trie                | Prefix tree                  |           |             |
| Associative | LRU Cache           | Least recently used eviction |           |             |
| Associative | TreeMap             | Tree-based ordered map       |           |             |
| Associative | BiMap               | Bidirectional mapping        |           |             |
| Associative | Ternary Search Tree | Hybrid trie/BST              |           |             |

## v0.2 Roadmap

| Category   | Structure        | Notes                            | Interface | Implemented |
| ---------- | ---------------- | -------------------------------- | --------- | ----------- |
| Sequential | RingBuffer       | Circular buffer                  |           |             |
| Sequential | Gap Buffer       | Efficient insertions at cursor   |           |             |
| Sequential | Rope             | For large strings                |           |             |
| Sets       | TreeSet          | Tree-based ordered set           |           |             |
| Sets       | BitSet           | Compact boolean array            |           |             |
| Trees      | Red-Black Tree   | Relaxed balanced BST             |           |             |
| Trees      | B+ Tree          | Leaf-linked B-Tree               |           |             |
| Trees      | Splay Tree       | Self-adjusting BST               |           |             |
| Trees      | Segment Tree     | Range queries                    |           |             |
| Trees      | Interval Tree    | Overlapping interval queries     |           |             |
| Trees      | K-D Tree         | Multi-dimensional search         |           |             |
| Trees      | R-Tree           | Spatial indexing                 |           |             |
| Trees      | Suffix Tree      | Substring queries                |           |             |
| Trees      | Suffix Array     | Space-efficient suffix structure |           |             |
| Trees      | Merkle Tree      | Hash-based verification          |           |             |
| Heaps      | Binary Heap      | Standard heap                    |           |             |
| Heaps      | Fibonacci Heap   | Amortized O(1) decrease-key      |           |             |
| Heaps      | Binomial Heap    | Mergeable heap                   |           |             |
| Graphs     | Adjacency List   | Sparse graph representation      |           |             |
| Graphs     | Adjacency Matrix | Dense graph representation       |           |             |
| Graphs     | Edge List        | Simple edge collection           |           |             |
| Graphs     | Incidence Matrix | Edge-vertex relationships        |           |             |
| Graphs     | Disjoint Set     | Union-Find                       |           |             |
| Trees      | Quad Tree        | 2D spatial partitioning          |           |             |
| Trees      | Octree           | 3D spatial partitioning          |           |             |

## v0.3 Roadmap

| Category      | Structure     | Notes                                | Interface | Implemented |
| ------------- | ------------- | ------------------------------------ | --------- | ----------- |
| Sequential    | Skip List     | Probabilistic balanced structure     |           |             |
| Associative   | Radix Tree    | Compressed trie                      |           |             |
| Associative   | Patricia Trie | Space-optimized trie                 |           |             |
| Sets          | Cuckoo Filter | Space-efficient alternative to Bloom |           |             |
| Sets          | HyperLogLog   | Cardinality estimation               |           |             |
| Probabilistic | Bloom Filter  | Probabilistic membership             |           |             |
| Trees         | AVL Tree      | Strictly balanced BST                |           |             |
| Trees         | Treap         | Randomized BST                       |           |             |
| Trees         | Fenwick Tree  | Binary indexed tree                  |           |             |
| Trees         | Quad Tree     | 2D spatial partitioning              |           |             |
| Trees         | Octree        | 3D spatial partitioning              |           |             |
| Heaps         | Pairing Heap  | Simplified Fibonacci heap            |           |             |
| Heaps         | D-ary Heap    | Generalized binary heap              |           |             |
| Heaps         | Min-Max Heap  | Double-ended priority queue          |           |             |
| Heaps         | Leftist Heap  | Mergeable heap                       |           |             |
| Heaps         | Skew Heap     | Self-adjusting leftist heap          |           |             |

## v0.4 Roadmap

Concurrent safe implementations of v0.1

## v0.5 Roadmap

Concurrently safe implementations of version v0.1 and v0.2

| Category   | Structure           | Notes                         | Interface | Implemented |
| ---------- | ------------------- | ----------------------------- | --------- | ----------- |
| Concurrent | Lock-Free Queue     | Non-blocking FIFO             |           |             |
| Concurrent | Lock-Free Stack     | Non-blocking LIFO             |           |             |
| Concurrent | Concurrent HashMap  | Thread-safe map               |           |             |
| Concurrent | MPMC Queue          | Multi-producer multi-consumer |           |             |
| Concurrent | Work-Stealing Deque | For task scheduling           |           |             |
