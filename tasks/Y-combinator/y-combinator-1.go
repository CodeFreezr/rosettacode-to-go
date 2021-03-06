package main

import "fmt"

type Func func(int) int
type FuncFunc func(Func) Func
type RecursiveFunc func(RecursiveFunc) Func

func main() {
	fac := Y(almost_fac)
	fib := Y(almost_fib)
	fmt.Println("fac(10) = ", fac(10))
	fmt.Println("fib(10) = ", fib(10))
}

func Y(f FuncFunc) Func {
	g := func(r RecursiveFunc) Func {
		return f(func(x int) int {
			return r(r)(x)
		})
	}
	return g(g)
}

func almost_fac(f Func) Func {
	return func(x int) int {
		if x <= 1 {
			return 1
		}
		return x * f(x-1)
	}
}

func almost_fib(f Func) Func {
	return func(x int) int {
		if x <= 2 {
			return 1
		}
		return f(x-1) + f(x-2)
	}
}

//\Y-combinator\y-combinator-1.go
