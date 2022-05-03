package iinit

import (
	"sync"
)

type adjList map[*Operator][]*Operator

func (l adjList) push(k *Operator, v ...*Operator) {
	l[k] = append(l[k], v...)
}

type opGraph struct {
	sync.Mutex
	all         map[*Operator]struct{}
	deps, rdeps adjList
	staticCache map[uintptr]*Operator
}

var g opGraph = opGraph{
	deps:        make(map[*Operator][]*Operator),
	rdeps:       make(map[*Operator][]*Operator),
	all:         make(map[*Operator]struct{}),
	staticCache: make(map[uintptr]*Operator),
}

func (o *opGraph) panicIfLoop() {
	const (
		UNVISITED = 0
		VIEWED    = 1
		PROCESSED = 2
	)
	state := make(map[*Operator]int)
	q := make([]*Operator, 0)
	for k := range o.all {
		if len(o.deps[k]) == 0 {
			q = append(q, k)
			state[k] = VIEWED
		} else {
			state[k] = UNVISITED
		}
	}
	for len(q) != 0 {
		v := q[0]

		// pop first and shift
		for i := 1; i < len(q); i++ {
			q[i-1] = q[i]
		}
		q = q[:len(q)-1]

		state[v] = PROCESSED
		for _, u := range o.rdeps[v] {
			if state[u] == PROCESSED {
				panic("loop detected")
			} else if state[u] == UNVISITED {
				state[u] = VIEWED
				q = append(q, u)
			}
		}
	}
}
