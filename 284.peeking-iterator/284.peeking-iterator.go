type PeekingIterator struct {
	Iter    *Iterator
	HasNext bool
	Next    int
}

func Constructor(iter *Iterator) *PeekingIterator {
	return &PeekingIterator{iter, iter.hasNext(), iter.next()}
}

func (p *PeekingIterator) hasNext() bool {
	return p.HasNext
}

func (p *PeekingIterator) next() int {
	defer func() {
		p.HasNext = p.Iter.hasNext()
		p.Next = p.Iter.next()

	}()

	return p.Next
}

func (p *PeekingIterator) peek() int {
	return p.Next
}
