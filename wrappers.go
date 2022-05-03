package iinit

// Call Static on every op and return Sequential
func SequentialS(ops ...func()) *Operator {
	converted := make([]*Operator, len(ops))
	for i, op := range ops {
		converted[i] = Static(op)
	}
	return Sequential(converted...)
}

// Call New on every op and return Sequential
func SequentialN(ops ...func()) *Operator {
	converted := make([]*Operator, len(ops))
	for i, op := range ops {
		converted[i] = New(op)
	}
	return Sequential(converted...)
}

// Call Static on every op and return Parallel
func ParallelS(ops ...func()) *Operator {
	converted := make([]*Operator, len(ops))
	for i, op := range ops {
		converted[i] = Static(op)
	}
	return Parallel(converted...)
}

// Call New on every op and return Parallel
func ParallelN(ops ...func()) *Operator {
	converted := make([]*Operator, len(ops))
	for i, op := range ops {
		converted[i] = New(op)
	}
	return Parallel(converted...)
}
