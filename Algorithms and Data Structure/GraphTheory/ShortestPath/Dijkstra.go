package ShortestPath

import "container/heap"

type edge struct {
	to, wt int
}

//堆优化的dijkstra
//TODO: 修改并测试
func dijkstra(g [][]edge, start int) []int {
	dis := make([]int, len(g))
	for i := range dis {
		dis[i] = 1e18
	}
	dis[start] = 0
	h := hp{{start, 0}}
	for len(h) > 0 {
		p := heap.Pop(&h).(pair)
		v := p.v
		if dis[v] < p.dis {
			continue
		}
		for _, e := range g[v] {
			w := e.to
			newDis := dis[v] + e.wt
			if newDis < dis[w] {
				dis[w] = newDis
				heap.Push(&h, pair{w, newDis})
			}
		}
	}
	return dis
}

type pair struct{ v, dis int }
type hp []pair

func (h hp) Len() int              { return len(h) }
func (h hp) Less(i, j int) bool    { return h[i].dis < h[j].dis }
func (h hp) Swap(i, j int)         { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v interface{})   { *h = append(*h, v.(pair)) }
func (h *hp) Pop() (v interface{}) { a := *h; *h, v = a[:len(a)-1], a[len(a)-1]; return }

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
