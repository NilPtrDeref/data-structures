# Data Structures in Go

A collection of basic data structures implemented in Go using generics.

## Contents

- [Binary Search Tree (BST)](#binary-search-tree-bst)
- [Linked List](#linked-list)
- [Queue](#queue)
- [Stack](#stack)

## Data Structures

### Binary Search Tree (BST)

A generic Binary Search Tree implementation. It requires a comparison function to order elements.

```go
import "github.com/woodywood117/data-structures/bst"

func compareInt(a, b int) int {
	if a < b { return -1 }
	if a > b { return 1 }
	return 0
}

// The comparison function must be passed as a pointer to a variable
var cmp = compareInt
tree := bst.New[int](&cmp)

tree.Insert(5)
tree.Insert(3)
if tree.Contains(5) {
    tree.Remove(5)
}
```

### Linked List

A generic doubly linked list.

```go
import "github.com/woodywood117/data-structures/linked_list"

list := linked_list.New[int]()
list.Add(10)
list.Add(20)

val, err := list.PopHead() // 10
length := list.Length()    // 1
```

### Queue

A generic FIFO (First-In-First-Out) queue implemented using a linked list.

```go
import "github.com/woodywood117/data-structures/queue"

q := queue.New[int]()
q.Enqueue(1)
q.Enqueue(2)

val, err := q.Dequeue() // 1
```

### Stack

A generic LIFO (Last-In-First-Out) stack with a fixed capacity. Note that this implementation works with pointers to the data type.

```go
import "github.com/woodywood117/data-structures/stack"

s, err := stack.New[int](10)
val := 42
s.Push(&val)

top, err := s.Pop() // *int (42)
```

## Running Tests

To run tests for all packages:

```bash
go test ./...
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
