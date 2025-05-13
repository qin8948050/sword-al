package sort

// LRU算法
type LRUCache struct {
	Size     int
	Capacity int
	Cache    map[int]*DlinkNode
	Head     *DlinkNode
	Tail     *DlinkNode
}

type DlinkNode struct {
	Pre   *DlinkNode
	Next  *DlinkNode
	Value int
	Key   int
}

func initDlinkNode(key int, value int) *DlinkNode {
	return &DlinkNode{
		Key:   key,
		Value: value,
	}
}

func Constructor(capacity int) LRUCache {
	head := initDlinkNode(0, 0)
	tail := initDlinkNode(0, 0)
	head.Next = tail
	tail.Pre = head

	return LRUCache{
		Capacity: capacity,
		Head:     head,
		Tail:     tail,
		Cache:    make(map[int]*DlinkNode),
	}
}

func (this *LRUCache) addToHead(node *DlinkNode) {
	node.Pre = this.Head
	node.Next = this.Head.Next
	this.Head.Next.Pre = node
	this.Head.Next = node
}

func (this *LRUCache) removeFromList(node *DlinkNode) {
	node.Pre.Next = node.Next
	node.Next.Pre = node.Pre
}

func (this *LRUCache) moveNode(node *DlinkNode) {
	this.removeFromList(node)
	this.addToHead(node)
}

func (this *LRUCache) removeNode() {
	if this.Size > this.Capacity {
		node := this.Tail.Pre
		this.removeFromList(node)
		delete(this.Cache, node.Key)
		this.Size--
	}
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.Cache[key]; ok {
		this.moveNode(node)
		return node.Value
	}
	return -1
}

func (this *LRUCache) Add(key int, node *DlinkNode) {
	this.Cache[key] = node
	this.addToHead(node)
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.Cache[key]; ok {
		node.Value = value
		this.moveNode(node)
	} else {
		node := initDlinkNode(key, value)
		this.Add(key, node)
		this.Size++
		this.removeNode()
	}
}
