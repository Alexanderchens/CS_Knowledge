package PageReplacement

type PageNum int
type PageType int

const NullPage PageType = -1

type DualListNode struct {
	_PageNum     PageNum
	_PageContent PageType
	pre, next    *DualListNode
}

type LRUCache struct {
	mp         map[PageNum]*DualListNode
	size       int
	now        int
	head, tail *DualListNode
}

func Constructor(capacity int) LRUCache {
	_LRUCache := LRUCache{make(map[PageNum]*DualListNode), capacity, 0, nil, nil}
	return _LRUCache
}

func (this *LRUCache) Get(key PageNum) PageType {
	if res, ok := this.mp[key]; ok {
		return res._PageContent
	}
	return -1
}

func (this *LRUCache) Put(key PageNum, value PageType) {
	if _, ok := this.mp[key]; ok {
		this.mp[key]._PageContent = value
		this.AddToHead(this.mp[key])
		return
	}
	this.mp[key] = &DualListNode{key, value, nil, nil}
	this.now++
	this.AddToHead(this.mp[key])
	if this.now > this.size {
		this.RemoveTail()
	}
}

func (this *LRUCache) GetFromMid(target *DualListNode) {

}

func (this *LRUCache) AddToHead(target *DualListNode) {
	this.GetFromMid(target)
	target.next = this.head
	this.head.pre = target
	this.head = target
}

func (this *LRUCache) RemoveTail() {
	p := this.tail
	this.tail = this.tail.pre
	this.tail.next = nil
	p.pre = nil
	delete(this.mp, p._PageNum)
}

func main() {

}
