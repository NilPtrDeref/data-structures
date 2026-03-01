package lru

type node[T comparable] struct {
	previous, next *node[T]
	value          T
}

type Lru[T comparable] struct {
	head, tail       *node[T]
	hash             map[T]*node[T]
	length, capacity int
}

func New[T comparable](capacity int) *Lru[T] {
	return &Lru[T]{
		hash:     make(map[T]*node[T]),
		capacity: capacity,
	}
}

func lpush[T comparable](lru *Lru[T], item T) *node[T] {
	n := &node[T]{value: item}
	if lru.head == nil {
		lru.head = n
		lru.tail = n
		return n
	}

	lru.head.previous = n
	n.next = lru.head
	lru.head = n
	return n
}

func lremove[T comparable](lru *Lru[T], n *node[T]) {
	if n == nil {
		return
	}

	if n == lru.head {
		lru.head = n.next
		if lru.head != nil {
			lru.head.previous = nil
		}
		return
	}

	if n == lru.tail {
		lru.tail = n.previous
		if lru.tail != nil {
			lru.tail.next = nil
		}
		return
	}

	if n.next != nil {
		n.next.previous = n.previous
	}

	if n.previous != nil {
		n.previous.next = n.next
	}
}

func Put[T comparable](lru *Lru[T], item T) {
	n, ok := lru.hash[item]
	if !ok {
		lru.hash[item] = lpush(lru, item)
		lru.length++
	} else {
		lremove(lru, n)
		lru.hash[item] = lpush(lru, n.value)
	}

	if lru.length > lru.capacity {
		lru.length--
		delete(lru.hash, lru.tail.value)
		lremove(lru, lru.tail)
	}
}

func Get[T comparable](lru *Lru[T], item T) *T {
	n, ok := lru.hash[item]
	if !ok {
		return nil
	}

	lremove(lru, n)
	lru.hash[item] = lpush(lru, n.value)
	return &n.value
}
