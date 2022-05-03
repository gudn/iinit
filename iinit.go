package iinit

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
