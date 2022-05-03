package iinit

import "reflect"

// Holder of init target
//
// All operators is just `func()` function. If error occured, you should degrade
// (for example, replace with default value) or call panic.
type Operator struct {
	f func()
}

// Create a new Operator from given functions
//
// This create a new Operator for every call. Use it for closures
func New(f func()) *Operator {
	op := &Operator{f}
	g.Lock()
	defer g.Unlock()
	g.all[op] = struct{}{}
	return op
}

// Create a static Operator from static function (not closure)
//
// You should use this functions when it possible because every Operator'll be
// called only once
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

// Add to `v` some dependencies.
// This means `v` will be runned after all it dependencies
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

// Create ephemeral operator and add all provided as deps
func Parallel(ops ...*Operator) *Operator {
	op := New(func() {})
	op.Deps(ops...)
	return op
}

// Define a sequence of operators
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
