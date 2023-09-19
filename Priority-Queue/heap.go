package main

import (
	"container/heap"
	"fmt"
)

type Item struct {
	Value string
	P     int
}
type P []*Item

func (p P) Len() int           { return len(p) }
func (p P) Less(i, j int) bool { return p[i].P > p[j].P }
func (p P) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}
func (p *P) Push(x any) {
	item := x.(*Item)
	*p = append(*p, item)
}
func (p *P) Pop() any {
	old := *p
	x := old[len(old)-1]
	*p = old[0 : len(old)-1]
	return x
}
func main() {
	m := P{}

	item1 := &Item{Value: "red", P: 1}
	item2 := &Item{Value: "blue", P: 2}
	item3 := &Item{Value: "black", P: 3}
	heap.Init(&m)
	heap.Push(&m, item1)
	heap.Push(&m, item2)
	heap.Push(&m, item3)
	for m.Len() > 0 {
		fmt.Println(m.Pop().(*Item).Value)
	}
}
