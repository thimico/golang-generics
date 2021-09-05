package main 

import (
	"fmt"
	"time"
)

type myInt int

type SmallInts interface {
}

type Ints interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

func add[T Ints](a, b T) T {
	return a + b
}

func example1() {
	var a myInt = 2
	var b myInt = 3

	fmt.Printf("Hello gerneric world 2 + 3 =%d\n", add(a,b))
}

func existsInStringSlice(el string, s []string) bool {
	for _, str := range s {
		if el == str {
			return true
		}
	}

	return false
}

func example2() {
	s := []string{"a", "b", "c"}
	el := "a"
	fmt.Printf("existsInStringSlice(%q, %v): %v\n", el, s, existsInStringSlice(el, s))
	
	el = "e"

	fmt.Printf("existsInStringSlice(%q, %v): %v\n", el, s, existsInStringSlice(el, s))
}

func existsInSlice[T comparable](el T, s []T) bool {
	for _, se := range s {
		if el == se {
			return true
		}
	}

	return false
}



func example3() {
	ints := []int{1, 2, 3}
	intEl := 3
	fmt.Printf("existsInSlice(%d, %v): %v\n", intEl, ints, existsInSlice(intEl, ints))
	s := []string{"a", "b", "c"}
	el := "a"
	fmt.Printf("existsInSlice(%q, %v): %v\n", el, s, existsInSlice(el, s))
	
	el = "e"

	fmt.Printf("existsInSlice(%q, %v): %v\n", el, s, existsInSlice(el, s))

}


func unique[T comparable](in []T) []T {
	keys := make(map[T]struct{})
	var out []T

	for _, v := range in {
		if _, ok := keys[v]; ok {
			continue
		}

		keys[v] = struct{}{}
		out = append(out, v)
	}

	return out
}

type set[T comparable] map[T]struct{}

func newEmptySet[T comparable]() set[T] {
	return make(set[T])
}

func (s set[T]) Add(el T) {
	s[el] = struct{}{}
}

func (s set[T]) Exists(el T) bool {
	_, ok := s[el]
	return ok

}

func (s set[T]) Delete(el T) {
	delete(s, el)
}

func uniqueGeneric[T comparable](in []T) []T {
	s := newEmptySet[T]()
	var out []T

	for _, v := range in {
		if s.Exists(v) {
			continue
		}

		s.Add(v)
		out = append(out, v)
	}

	return out
}

func factorial(n int) int {
	var res = 1

	for i := 2; i <= n; i++ {
		res *= i
	
	}
	return res
}

func example5() {
	arg := 4

	fmt.Printf("factorial(%d): %d\n", arg, factorial(arg))
	
}


func decorator(f func(int) int) func(int) int {
	return func(a int) int {
		start := time.Now()

		res := f(a)

		fmt.Printf("f() took %s\n", time.Since(start))

		return res
	}
}

func example6() {
	arg := 4
	
	f := decorator(factorial)

	fmt.Printf("factorial(%d): %d\n", arg, f(arg))
	
}


func timingDecorator[T, V any](name string, f func(T) (V, error)) func(T) (V, error) {
	return func(req T) (V, error) {
		start := time.Now()

		res, err := f(req)

		fmt.Printf("%s() took %s\n", name, time.Since(start))

		return res, err
	}
}

type myRequest struct{}
type myResponse struct{}
func handleRequest(req *myRequest) (*myResponse, error) {
	time.Sleep(time.Millisecond)

	return &myResponse{}, nil
}

func example7() {
	f := timingDecorator("handleRequest", handleRequest)	
	
	res, err := f(&myRequest{})
	fmt.Printf("handleRequest(): %v, %v\n", res, err)
	
}


func main() {
	example1()
	example2()
	example3()
	example5()	
	example6()	
	example7()	

}
