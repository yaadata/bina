package btree

import (
	"cmp"
	"iter"

	"codeberg.org/yaadata/bina/core/collection"
	"codeberg.org/yaadata/bina/core/kv"
	. "codeberg.org/yaadata/opt"
)

type builtinImpl[K cmp.Ordered, V any] struct {
	branchingFactor int
	height          int
	len             int
	min             Option[kv.Pair[K, V]]
	max             Option[kv.Pair[K, V]]
	root            Option[Node[K, V]]
}

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
