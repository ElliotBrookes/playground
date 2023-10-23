package pkg

import "fmt"

type node[T any] struct {
	value T
	next  *node[T]
}

type SingleList[T any] struct {
	Head, Tail *node[T]
}

func (*SingleList[T]) Add(T) {
	//
	fmt.Println("Added to LL")
}

func (*SingleList[T]) Pop(T) {
	//
	fmt.Println("Popped 'Next' value from LL")
}

func (*SingleList[T]) Peek() {
	//
	fmt.Println("Peeked into LL")
}
