type Node struct {
	val int
	next *Node
}

type LinkedList struct {
	head *Node
	tail *Node
	size int
}

func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

func (ll *LinkedList) Get(index int) int {
	if index < 0 || index >= ll.size {
		return -1
	}
	curr := ll.head
	for i := 0; i < index; i++ {
		curr = curr.next
	}
	return curr.val
}

func (ll *LinkedList) InsertHead(val int) {
	node := &Node{val: val, next: ll.head}
	ll.head = node

	if ll.tail == nil {
		ll.tail = node
	}

	ll.size++
}

func (ll *LinkedList) InsertTail(val int) {
	node := &Node{val: val, next: nil}
	
	if ll.tail == nil {
		ll.head = node
		ll.tail = node
	} else {
		ll.tail.next = node
		ll.tail = node
	}

	ll.size++
}

func (ll *LinkedList) Remove(index int) bool {
	if index < 0 || index >= ll.size {
		return false
	}

	curr := ll.head
	var prev *Node
	for i := 0; i < index; i++ {
		prev = curr
		curr = curr.next
	}

	if prev == nil {
		ll.head = curr.next
	} else {
		prev.next = curr.next
	}

	if curr == ll.tail {
		ll.tail = prev
	}
	
	ll.size--
	if ll.size == 0 {
		ll.head = nil
		ll.tail = nil
	}
	return true
}

func (ll *LinkedList) GetValues() []int {
	arr := make([]int, 0, ll.size)

	
	for curr := ll.head; curr != nil; curr = curr.next {
		arr = append(arr, curr.val)
	}
	return arr
}
