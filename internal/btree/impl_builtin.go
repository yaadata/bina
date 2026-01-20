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

func (b *builtinImpl[K, V]) GetNode(key K) Option[collection.BTreeNode[V]] {
	node := get(b.root, key)
	if node.IsNone() {
		return None[collection.BTreeNode[V]]()
	}
	n := node.Unwrap()
	var res collection.BTreeNode[V] = &n
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
	for child := range n.lessBranch.Values() {
		res = append(res, inorder(Some(child))...)
	}
	res = append(res, n.inner)
	for child := range n.greaterBranch.Values() {
		res = append(res, inorder(Some(child))...)
	}
	return res
}

func preorder[K cmp.Ordered, V any](node Option[Node[K, V]]) []kv.Pair[K, V] {
	if node.IsNone() {
		return nil
	}
	n := node.Unwrap()
	res := []kv.Pair[K, V]{n.inner}
	for child := range n.lessBranch.Values() {
		res = append(res, inorder(Some(child))...)
	}
	for child := range n.greaterBranch.Values() {
		res = append(res, inorder(Some(child))...)
	}
	return res
}

func postorder[K cmp.Ordered, V any](node Option[Node[K, V]]) []kv.Pair[K, V] {
	if node.IsNone() {
		return nil
	}
	n := node.Unwrap()
	var res []kv.Pair[K, V]
	for child := range n.lessBranch.Values() {
		res = append(res, inorder(Some(child))...)
	}
	for child := range n.greaterBranch.Values() {
		res = append(res, inorder(Some(child))...)
	}
	return res
}

func get[K cmp.Ordered, V any](current Option[Node[K, V]], target K) Option[Node[K, V]] {
	if current.IsNone() {
		return current
	}
	node := current.Unwrap()
	switch cmp.Compare(target, node.inner.Key()) {
	case -1:
		sibling := node.lessSibling
		if sibling.IsSome() {
			if cmp.Compare(target, sibling.Unwrap().inner.Key()) <= 0 {
				return get(sibling, target)
			}
		}
		for child := range node.lessBranch.Values() {
			res := get(Some(child), target)
			if res.IsSome() {
				return res
			}
		}
	case 0:
		return current
	case 1:
		sibling := node.greaterSibling
		if sibling.IsSome() {
			if cmp.Compare(target, sibling.Unwrap().inner.Key()) <= 0 {
				return get(sibling, target)
			}
		}
		for child := range node.greaterBranch.Values() {
			res := get(Some(child), target)
			if res.IsSome() {
				return res
			}
		}
	}

	return None[Node[K, V]]()
}
