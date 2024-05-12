package main

import (
	"fmt"
)

type Tree struct {
	v int
	l *Tree
	r *Tree
}

func main() {
	tree := Tree{
		v: 10,
	}
	fmt.Println(tree)
	tree.Insert(13)
	tree.Insert(8)
	tree.Insert(6)
	fmt.Println(tree)
	fmt.Println(tree.Sum())
}

func (t *Tree) Insert(n int) {
	if t == nil {
		return
	}
	if n < t.v {
		if t.l == nil {
			t.l = &Tree{
				v: n,
			}
			return
		}
		t.l.Insert(n)
		return
	}
	if n > t.v {
		if t.r == nil {
			t.r = &Tree{
				v: n,
			}
			return
		}
		t.r.Insert(n)
		return
	}
	t.v = n
}

func (t *Tree) Sum() int {
	if t == nil {
		return 0
	}
	return t.v + t.l.Sum() + t.r.Sum()
}
