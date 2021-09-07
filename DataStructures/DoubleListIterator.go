package collection

type node struct {
	v interface{}
	next *node
	prev *node
}
func newNode(v interface{}) *node {
	return &node{
		v: v,
	}
}
type doubleList struct {
	head node
	tail node
	iterNode *node // support iterator interface
	len int
}

func DoubleList(values ...interface{}) *doubleList {
	l := &doubleList{}
	l.head = node{}
	l.tail = node{}
	l.head.prev = &l.tail
	l.tail.next = &l.head
	l.iterNode = &l.head
	if len(values) > 0 {
		node := &l.head
		for _, v := range values {
			e := newNode(v)
			node.next = e
			e.prev = node
			node = e
			l.len++
		}
		node.next = &l.tail
		l.tail.prev = node
		return l
	}
	l.head.next = &l.tail
	l.tail.prev = &l.head
	return l
}

func (d *doubleList) HasNext() bool {
	return d.iterNode.next != &d.tail
}

func (d *doubleList) Next() interface{} {
	node := d.iterNode.next
	d.iterNode = node
	return node.v
}

func (d *doubleList) PopBack() *node {
	// pop from tail
	if d.len > 0 {
		nodePop := d.tail.prev
		prevNode := nodePop.prev
		prevNode.next = &d.tail
		d.tail.prev = prevNode
		d.len--
		return nodePop
	}
	return nil
}

func (d *doubleList) Pop() *node {
	if d.len > 0 {
		nodePop := d.head.next
		nextNode := nodePop.next
		d.head.next = nextNode
		nextNode.prev = &d.head
		d.len--
		return nodePop
	}
	return nil
}

func (d *doubleList) Push(v interface{}) {
	node := newNode(v)
	nextNode := d.head.next
	
	d.head.next = node

	node.prev = &d.head
	node.next = nextNode

	nextNode.prev = node
} 

func (d *doubleList) PushBack(v interface{}) {
	node := newNode(v)
	prevNode := d.tail.prev

	d.tail.prev = node

	node.next = &d.tail
	node.prev = prevNode

	prevNode.next = node
}