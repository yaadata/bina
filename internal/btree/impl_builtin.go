package btree

import (
	"cmp"
	"iter"
	"slices"

	. "codeberg.org/yaadata/opt"

	"codeberg.org/yaadata/bina/core/collection"
	"codeberg.org/yaadata/bina/core/kv"
	"codeberg.org/yaadata/bina/core/predicate"
	"codeberg.org/yaadata/bina/core/where"
	"codeberg.org/yaadata/bina/internal/slice"
)

type builtinImpl[K cmp.Ordered, V any] struct {
	branchingFactor int
	height          int
	len             int
	min             Option[kv.Pair[K, V]]
	max             Option[kv.Pair[K, V]]
	root            Option[Node[K, V]]
}

var _ collection.BTree[int, int] = (*builtinImpl[int, int])(nil)

func (b *builtinImpl[K, V]) Len() int {
	return b.len
}

func (b *builtinImpl[K, V]) Contains(element K) bool {
	return get(b.root, element).IsSome()
}

func (b *builtinImpl[K, V]) IsEmpty() bool {
	return b.root.IsNone()
}

func (b *builtinImpl[K, V]) Clear() {
	b.root = None[Node[K, V]]()
	b.len = 0
	b.height = 0
}

func (b *builtinImpl[K, V]) Any(pred predicate.Predicate[kv.Pair[K, V]]) bool {
	// TODO: optimize to short circuit
	for key, value := range b.All() {
		if pred(kv.New(key, value)) {
			return true
		}
	}
	return false
}

func (b *builtinImpl[K, V]) Every(pred predicate.Predicate[kv.Pair[K, V]]) bool {
	for key, value := range b.All() {
		if !pred(kv.New(key, value)) {
			return false
		}
	}
	return true
}

func (b *builtinImpl[K, V]) Count(pred predicate.Predicate[kv.Pair[K, V]]) int {
	var count int
	for key, value := range b.All() {
		if pred(kv.New(key, value)) {
			count++
		}
	}
	return count
}

func (b *builtinImpl[K, V]) ForEach(fn func(pair kv.Pair[K, V])) {
	for key, value := range b.All() {
		fn(kv.New(key, value))
	}
}

func (b *builtinImpl[K, V]) Delete(key K) Option[V] {
	// TODO: implement me
	return None[V]()
}

func (b *builtinImpl[K, V]) Get(key K) Option[V] {
	node := b.GetNode(key)
	if node.IsNone() {
		return None[V]()
	}
	for _, element := range node.Unwrap().Values() {
		if element.Key() == key {
			return Some(element.Value())
		}
	}
	return None[V]()
}

func (b *builtinImpl[K, V]) Height() int {
	return b.height
}

func (b *builtinImpl[K, V]) Max() Option[kv.Pair[K, V]] {
	return b.max
}

func (b *builtinImpl[K, V]) Min() Option[kv.Pair[K, V]] {
	return b.min
}

func (b *builtinImpl[K, V]) Put(key K, value V) {
	pair := kv.New(key, value)

	// Case 1: Empty tree - create root
	if b.root.IsNone() {
		b.root = Some(Node[K, V]{
			elements: []kv.Pair[K, V]{pair},
			parent:   None[Node[K, V]](),
			children: slice.SliceFromComparableInterface[Node[K, V]](),
		})
		b.len = 1
		b.height = 1
		b.min = Some(pair)
		b.max = Some(pair)
		return
	}

	// Case 2: Insert recursively
	root := b.root.Unwrap()
	newRoot, promoted, wasInserted := insert(&root, key, value, b.branchingFactor)

	// Handle root split (promoted element needs new root)
	if promoted.IsSome() {
		p := promoted.Unwrap()
		newRootNode := Node[K, V]{
			elements: []kv.Pair[K, V]{p.key},
			parent:   None[Node[K, V]](),
			children: slice.SliceFromComparableInterface[Node[K, V]](),
		}
		newRootNode.children.Append(p.left)
		newRootNode.children.Append(p.right)
		b.root = Some(newRootNode)
		b.height++
	} else {
		b.root.Replace(newRoot)
	}

	if wasInserted {
		b.len++
		updateMinMax(b, key, value)
	}
}

func (b *builtinImpl[K, V]) Floor(key K) Option[kv.Pair[K, V]] {
	return floor(b.root, key, None[kv.Pair[K, V]]())
}

func (b *builtinImpl[K, V]) Ceiling(key K) Option[kv.Pair[K, V]] {
	return ceiling(b.root, key, None[kv.Pair[K, V]]())
}

func (b *builtinImpl[K, V]) Range(opts ...where.WhereOption[K]) iter.Seq2[K, V] {
	wh := where.Default[K]()
	for _, opt := range opts {
		opt(wh)
	}
	if wh.From().Key().IsNone() &&
		wh.To().Key().IsNone() {
		return b.All()
	}
	pairs := rangeInorder(b.root, wh.From().Key(), wh.To().Key())
	return func(yield func(K, V) bool) {
		for _, p := range pairs {
			if !yield(p.Key(), p.Value()) {
				return
			}
		}
	}
}

func (b *builtinImpl[K, V]) GetNode(key K) Option[collection.BTreeNode[K, V]] {
	node := get(b.root, key)
	if node.IsNone() {
		return None[collection.BTreeNode[K, V]]()
	}
	n := node.Unwrap()
	var res collection.BTreeNode[K, V] = &n
	return Some(res)
}

func (b *builtinImpl[K, V]) Order() int {
	return b.branchingFactor
}

func (b *builtinImpl[K, V]) All(opts ...collection.SearchTreeTraversalOption) iter.Seq2[K, V] {
	traversalCfg := collection.DefaultSearchTreeTraversalConfiguration()
	for _, opt := range opts {
		opt(traversalCfg)
	}
	var res []kv.Pair[K, V]
	switch traversalCfg.Strategy() {
	case collection.SearchTreeStrategyInOrder:
		res = inorder(b.root)
	case collection.SearchTreeStrategyPreOrder:
		res = preorder(b.root)
	case collection.SearchTreeStrategyPostOrder:
		res = postorder(b.root)
	}
	return func(yield func(K, V) bool) {
		for _, pair := range res {
			if !yield(pair.Key(), pair.Value()) {
				return
			}
		}
	}
}

func inorder[K cmp.Ordered, V any](node Option[Node[K, V]]) []kv.Pair[K, V] {
	if node.IsNone() {
		return nil
	}
	n := node.Unwrap()
	var res []kv.Pair[K, V]
	childrenLength := n.children.Len()
	elementsLength := len(n.elements)
	for i := 0; i < elementsLength; i++ {
		// Visit child[i] (everything less than keys[i])
		if i < childrenLength {
			res = append(res, inorder(n.children.Get(i))...)
		}
		// Visit keys[i]
		res = append(res, n.elements[i])
	}
	// Visit last child (everything greater than all keys)
	if childrenLength > elementsLength {
		res = append(res, inorder(n.children.Get(elementsLength))...)
	}

	return res
}

func preorder[K cmp.Ordered, V any](node Option[Node[K, V]]) []kv.Pair[K, V] {
	if node.IsNone() {
		return nil
	}
	n := node.Unwrap()
	childrenLength := n.children.Len()
	elementsLength := len(n.elements)

	// 1. Visit ALL keys in this node first
	res := make([]kv.Pair[K, V], 0, elementsLength)
	res = append(res, n.elements...)

	// 2. Then visit all children (in preorder)
	for i := 0; i <= elementsLength; i++ {
		if i < childrenLength {
			res = append(res, preorder(n.children.Get(i))...)
		}
	}

	return res

}

func postorder[K cmp.Ordered, V any](node Option[Node[K, V]]) []kv.Pair[K, V] {
	if node.IsNone() {
		return nil
	}
	n := node.Unwrap()
	childrenLength := n.children.Len()
	elementsLength := len(n.elements)

	res := make([]kv.Pair[K, V], 0, elementsLength)
	// 1. Visit all children nodes
	for i := 0; i <= elementsLength; i++ {
		if i < childrenLength {
			res = append(res, postorder(n.children.Get(i))...)
		}
	}

	// 2. Visit ALL keys in current node
	res = append(res, n.elements...)
	return res
}

func get[K cmp.Ordered, V any](current Option[Node[K, V]], target K) Option[Node[K, V]] {
	if current.IsNone() {
		return current
	}
	node := current.Unwrap()
	elements := node.elements

	// Binary search for target
	lo, hi := 0, len(elements)
	for lo < hi {
		mid := lo + (hi-lo)/2
		switch cmp.Compare(target, elements[mid].Key()) {
		case 0:
			return current // Found
		case -1:
			hi = mid
		case 1:
			lo = mid + 1
		}
	}

	// Not found, descend into child[lo]
	if lo < node.children.Len() {
		return get(node.children.Get(lo), target)
	}
	return None[Node[K, V]]()

}

func floor[K cmp.Ordered, V any](
	node Option[Node[K, V]],
	key K,
	candidate Option[kv.Pair[K, V]],
) Option[kv.Pair[K, V]] {
	if node.IsNone() {
		return candidate
	}
	n := node.Unwrap()
	elements := n.elements

	// Binary search for key position
	lo, hi := 0, len(elements)
	for lo < hi {
		mid := lo + (hi-lo)/2
		switch cmp.Compare(key, elements[mid].Key()) {
		case 0:
			return Some(elements[mid]) // Exact match
		case -1:
			hi = mid
		case 1:
			lo = mid + 1
		}
	}

	// lo is insertion point: elements[lo-1] < key < elements[lo]

	// Try to find closer floor in child subtree
	if lo < n.children.Len() {
		result := floor(n.children.Get(lo), key, candidate)
		if result.IsSome() {
			return result
		}
	}

	// Use elements[lo-1] as floor if it exists
	if lo > 0 {
		return Some(elements[lo-1])
	}

	return candidate
}

func ceiling[K cmp.Ordered, V any](
	node Option[Node[K, V]],
	key K,
	candidate Option[kv.Pair[K, V]],
) Option[kv.Pair[K, V]] {
	if node.IsNone() {
		return candidate
	}
	n := node.Unwrap()
	elements := n.elements

	// Binary search for key position
	lo, hi := 0, len(elements)
	for lo < hi {
		mid := lo + (hi-lo)/2
		switch cmp.Compare(key, elements[mid].Key()) {
		case 0:
			return Some(elements[mid]) // Exact match
		case -1:
			hi = mid
		case 1:
			lo = mid + 1
		}
	}

	// lo is insertion point: elements[lo-1] < key < elements[lo]

	// Try to find ceiling in child subtree
	if lo < n.children.Len() {
		result := ceiling(n.children.Get(lo), key, candidate)
		if result.IsSome() {
			return result
		}
	}

	// Use elements[lo] as ceiling if it exists
	if lo < len(elements) {
		return Some(elements[lo])
	}

	return candidate
}

// promotedKey holds a key promoted during split, with its left and right children
type promotedKey[K cmp.Ordered, V any] struct {
	key   kv.Pair[K, V]
	left  Node[K, V]
	right Node[K, V]
}

func insert[K cmp.Ordered, V any](
	node *Node[K, V],
	key K,
	value V,
	branchingFactor int,
) (Node[K, V], Option[promotedKey[K, V]], bool) {
	elements := node.elements

	// Binary search for key position
	lo, hi := 0, len(elements)
	for lo < hi {
		mid := lo + (hi-lo)/2
		switch cmp.Compare(key, elements[mid].Key()) {
		case 0:
			// Key exists - update value and return (no split, not inserted)
			node.elements[mid] = kv.New(key, value)
			return *node, None[promotedKey[K, V]](), false
		case -1:
			hi = mid
		case 1:
			lo = mid + 1
		}
	}

	// lo is now the insertion index

	if node.children.Len() == 0 {
		// Leaf node - insert here
		return insertIntoLeaf(node, lo, key, value, branchingFactor)
	}

	// Internal node - recurse into child
	child := node.children.Get(lo).Unwrap()
	newChild, promoted, wasInserted := insert(&child, key, value, branchingFactor)

	if promoted.IsSome() {
		// Child was split, insert promoted key into this node
		p := promoted.Unwrap()
		// Replace the child that was split with the left part
		// and insert the right part after it
		node.children.RemoveAt(lo)
		node.children.Insert(lo, p.left)
		node.children.Insert(lo+1, p.right)

		resultNode, resultPromoted := insertPromoted(node, p.key, branchingFactor)
		return resultNode, resultPromoted, wasInserted
	}

	// Update the child with the new version
	node.children.RemoveAt(lo)
	node.children.Insert(lo, newChild)

	return *node, None[promotedKey[K, V]](), wasInserted
}

func insertIntoLeaf[K cmp.Ordered, V any](
	node *Node[K, V],
	pos int,
	key K,
	value V,
	branchingFactor int,
) (Node[K, V], Option[promotedKey[K, V]], bool) {
	// Insert new pair at position
	newPair := kv.New(key, value)
	node.elements = slices.Insert(node.elements, pos, newPair)

	// Check for overflow
	maxKeys := 2*branchingFactor - 1
	if len(node.elements) <= maxKeys {
		return *node, None[promotedKey[K, V]](), true
	}

	// Split needed
	resultNode, promoted := splitNode(node, branchingFactor)
	return resultNode, promoted, true
}

func splitNode[K cmp.Ordered, V any](
	node *Node[K, V],
	branchingFactor int,
) (Node[K, V], Option[promotedKey[K, V]]) {
	t := branchingFactor
	midIdx := t - 1 // Index of median element (0-indexed)

	// Median element to promote
	median := node.elements[midIdx]

	// Left node keeps elements [0, midIdx)
	leftNode := Node[K, V]{
		elements: slices.Clone(node.elements[:midIdx]),
		parent:   node.parent,
		children: slice.SliceFromComparableInterface[Node[K, V]](),
	}

	// Right node gets elements [midIdx+1, end)
	rightNode := Node[K, V]{
		elements: slices.Clone(node.elements[midIdx+1:]),
		parent:   node.parent,
		children: slice.SliceFromComparableInterface[Node[K, V]](),
	}

	// Redistribute children if internal node
	if node.children.Len() > 0 {
		// Left gets children [0, midIdx+1)
		for i := 0; i <= midIdx; i++ {
			if child := node.children.Get(i); child.IsSome() {
				leftNode.children.Append(child.Unwrap())
			}
		}
		// Right gets children [midIdx+1, end)
		for i := midIdx + 1; i < node.children.Len(); i++ {
			if child := node.children.Get(i); child.IsSome() {
				rightNode.children.Append(child.Unwrap())
			}
		}
	}

	promoted := promotedKey[K, V]{
		key:   median,
		left:  leftNode,
		right: rightNode,
	}

	return leftNode, Some(promoted)
}

func insertPromoted[K cmp.Ordered, V any](
	node *Node[K, V],
	key kv.Pair[K, V],
	branchingFactor int,
) (Node[K, V], Option[promotedKey[K, V]]) {
	// Find position to insert promoted key
	pos := 0
	for pos < len(node.elements) && cmp.Compare(key.Key(), node.elements[pos].Key()) > 0 {
		pos++
	}

	// Insert key at position
	node.elements = slices.Insert(node.elements, pos, key)

	// Check for overflow
	maxKeys := 2*branchingFactor - 1
	if len(node.elements) <= maxKeys {
		return *node, None[promotedKey[K, V]]()
	}

	// Split this node too
	return splitNode(node, branchingFactor)
}

func updateMinMax[K cmp.Ordered, V any](tree *builtinImpl[K, V], key K, value V) {
	pair := kv.New(key, value)

	if tree.min.IsNone() || cmp.Compare(key, tree.min.Unwrap().Key()) < 0 {
		tree.min = Some(pair)
	}
	if tree.max.IsNone() || cmp.Compare(key, tree.max.Unwrap().Key()) > 0 {
		tree.max = Some(pair)
	}
}

func rangeInorder[K cmp.Ordered, V any](
	node Option[Node[K, V]],
	from Option[K],
	to Option[K],
) []kv.Pair[K, V] {
	if node.IsNone() {
		return nil
	}
	n := node.Unwrap()
	var res []kv.Pair[K, V]
	childrenLength := n.children.Len()
	elementsLength := len(n.elements)
	hitUpperBound := false

	for i := 0; i < elementsLength; i++ {
		key := n.elements[i].Key()

		// Check if we should visit the left child
		// (only if this key >= from, meaning left subtree might have valid keys)
		if i < childrenLength {
			if from.IsNone() || cmp.Compare(key, from.Unwrap()) >= 0 {
				res = append(res, rangeInorder(n.children.Get(i), from, to)...)
			}
		}

		// Check upper bound first - if key >= to, stop traversal
		if to.IsSome() && cmp.Compare(key, to.Unwrap()) >= 0 {
			hitUpperBound = true
			break
		}

		// Check lower bound - only include if key >= from
		if from.IsNone() || cmp.Compare(key, from.Unwrap()) >= 0 {
			res = append(res, n.elements[i])
		}
	}

	// Visit last child only if we didn't hit the upper bound
	if !hitUpperBound && childrenLength > elementsLength {
		res = append(res, rangeInorder(n.children.Get(elementsLength), from, to)...)
	}

	return res
}
