package pkg

type LinkedList[V any] interface {
	Add(*LinkedList[V], V)
	Pop(*LinkedList[V], V)
	Peek(*LinkedList[V])
}
