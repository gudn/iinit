package iinit_test

import (
	"fmt"

	"github.com/gudn/iinit"
)

func first() {
	fmt.Println("first")
}

func second() {
	fmt.Println("second")
}

func third() {
	fmt.Println("third")
}

func printer(what string) func() {
	return func() {
		fmt.Println("printer: ", what)
	}
}

func init() {
	iinit.Sequential(
		iinit.ParallelS(third, second),
		iinit.New(printer("after parallel")),
	)
	iinit.Sequential(
		iinit.Static(first),
		iinit.Static(third),
		iinit.New(printer("after third")),
	)
	iinit.Iinit()
}

func Example() {
	iinit.Rerun(iinit.Static(second))
}
