package iinit

import "reflect"

type Operator struct {
	f func()
}

func New(f func()) *Operator {
	op := &Operator{f}
	g.Lock()
	defer g.Unlock()
	g.all[op] = struct{}{}
	return op
}

func Static(f func()) *Operator {
	g.Lock()
	defer g.Unlock()
	addr := reflect.ValueOf(f).Pointer()
	if val, ok := g.staticCache[addr]; ok {
		return val
	} else {
		// NOTE New(f) inlined for reuse lock
		val := &Operator{f}
		g.all[val] = struct{}{}
		g.staticCache[addr] = val
		return val
	}
}

func (v *Operator) Deps(deps ...*Operator) {
	if v == nil {
		return
	}
	g.Lock()
	defer g.Unlock()
	g.deps.push(v, deps...)
	for _, dep := range deps {
		g.rdeps.push(dep, v)
	}
}

func Parrallel(ops ...*Operator) *Operator {
	op := New(func() {})
	op.Deps(ops...)
	return op
}

func Sequential(ops ...*Operator) *Operator {
	if len(ops) == 0 {
		return nil
	}
	prev := ops[0]
	for _, op := range ops[1:] {
		if prev != nil {
			op.Deps(prev)
		}
		prev = op
	}
	return prev
}
