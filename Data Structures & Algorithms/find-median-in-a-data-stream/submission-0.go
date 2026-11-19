type MedianFinder struct {
    small IntHeap
	large IntHeap
}

type IntHeap []int

func (h IntHeap) Len() int            { return len(h) }
func (h IntHeap) Less(i, j int) bool   { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)        { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(n any) {
	*h = append(*h, n.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func Constructor() MedianFinder {
    return MedianFinder{}
}

func (this *MedianFinder) AddNum(num int)  {
	if len(this.large) > 0 && num > this.large[0] {
		heap.Push(&this.large, num)
	} else {
		heap.Push(&this.small, -num)
	}

	if len(this.small) > len(this.large) + 1 {
		heap.Push(&this.large, -heap.Pop(&this.small).(int))
	}
	if len(this.large) > len(this.small) + 1 {
		heap.Push(&this.small, -heap.Pop(&this.large).(int))
	}
}


func (this *MedianFinder) FindMedian() float64 {
    if len(this.small) > len(this.large) {
		return float64(-this.small[0])
	} else if len(this.large) > len(this.small) {
		return float64(this.large[0])
	}
	return float64(-this.small[0] + this.large[0])/2
}
