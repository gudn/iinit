package iinit

import (
	"sync"
)

type runnable struct {
	wait *sync.WaitGroup
	next []*sync.WaitGroup
	op *Operator
}

func newRunnable(op *Operator) runnable {
	return runnable{new(sync.WaitGroup), nil, op}
}

type runner struct {
	runnables map[*Operator]runnable
}

func (r runner) buildRevDeps(v *Operator) *sync.WaitGroup {
	if t, ok := r.runnables[v]; ok {
		return t.wait
	}
	t := newRunnable(v)
	for _, dep := range g.rdeps[v] {
		wg := r.buildRevDeps(dep)
		wg.Add(1)
		t.next = append(t.next, wg)
	}
	r.runnables[v] = t
	return t.wait
}

func (r runner) run() {
	all := new(sync.WaitGroup)
	for _, t := range r.runnables {
		all.Add(1)
		go func(t runnable) {
			t.wait.Wait()
			t.op.f()
			for _, n := range t.next {
				n.Done()
			}
			all.Done()
		}(t)
	}
	all.Wait()
}
