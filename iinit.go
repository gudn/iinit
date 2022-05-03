// Some initialization system
//
// It manages things called Operator. This is just `func()`. User defined graph
// of dependencies and run it. Iinit starts all functions parallel (in many
// goorutines) and defines syncronization between them.
//
// This is not optimal but process (usually) call only once in program life
package iinit

// Run current graph
func Iinit() {
	g.Lock()
	defer g.Unlock()
	g.panicIfLoop()
	r := runner{
		make(map[*Operator]runnable),
	}
	for k := range g.all {
		if len(g.deps[k]) == 0 {
			r.buildRevDeps(k)
		}
	}
	r.run()
}

// Rerun specified operator and all operators which depends on it (recursive)
func Rerun(v *Operator) {
	g.Lock()
	defer g.Unlock()
	g.panicIfLoop()
	r := runner{
		make(map[*Operator]runnable),
	}
	r.buildRevDeps(v)
	r.run()
}
