type Node struct {
	key  int
	val  int
	next *Node
	prev *Node
}

type LRUCache struct {
    capacity 	int
	hashMap		map[int]*Node
	tail		*Node
	head		*Node
}

func Constructor(capacity int) LRUCache {
	hashMap := make(map[int]*Node)
    cache := LRUCache{
		capacity: capacity,
		hashMap: hashMap,
		head: &Node{},
		tail: &Node{},
		}

	cache.head.next = cache.tail
	cache.tail.prev = cache.head
	return cache
}

func (l *LRUCache) Get(key int) int {
	curr, exists := l.hashMap[key]
	if !exists {
		return -1
	}
	
	curr.prev.next = curr.next
	curr.next.prev = curr.prev

	l.tail.prev.next = curr
	curr.next = l.tail
	curr.prev = l.tail.prev
	l.tail.prev = curr

	return curr.val
}

func (l *LRUCache) Put(key int, value int) {
	curr, exists := l.hashMap[key]
	if exists {
		curr.val = value

		curr.prev.next = curr.next
		curr.next.prev = curr.prev

		l.tail.prev.next = curr
		curr.next = l.tail
		curr.prev = l.tail.prev
		l.tail.prev = curr
		return
	}

	l.hashMap[key] = &Node{
		key: key,
		val: value,
	}
	curr, _ = l.hashMap[key]
	l.tail.prev.next = curr
	curr.next = l.tail
	curr.prev = l.tail.prev
	l.tail.prev = curr

	if len(l.hashMap) > l.capacity {
		lru := l.head.next
		delete(l.hashMap, lru.key)
		l.head.next = lru.next
		lru.next.prev = l.head
	}
	return
}
