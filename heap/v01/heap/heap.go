package heap

import (
  "golang.org/x/exp/constraints"
)

type Heap[K constraints.Integer, V any] struct {
  key []K
  val []V
  cur int
}

func NewHeap[K constraints.Integer, V any] () *Heap[K,V] {
  return &Heap[K,V]{
    key: make([]K, 1),
    val: make([]V, 1),
    cur: -1,
  }
}

func (h *Heap[K,V]) Push(k K, v V) {

  h.cur = h.cur + 1
  if len(h.key) > h.cur {
    h.key[h.cur] = k
    h.val[h.cur] = v
  } else {
    h.key = append(h.key, k)
    h.val = append(h.val, v)
  }

  i := h.cur
  for i > 0 && h.key[i] > h.key[parent(i)] {
    swap(h, i, parent(i))
    i = parent(i)
  }
}

func (h *Heap[K,V]) Pop() (K, V) {
  //call Empty() to confirm
  k := h.key[0]
  v := h.val[0]
  if h.cur == 0 {
    h.cur = -1
  } else {
    h.key[0] = h.key[h.cur]
    h.val[0] = h.val[h.cur]
    h.cur = h.cur - 1
    heapify(h, 0)
  }
  return k, v
}

func (h *Heap[K,V]) Empty() bool {
  return h.cur == -1
}

//max heapify
func heapify[K constraints.Integer, V any](h *Heap[K,V], i int) {
  max := i
  if left(i)  <= h.cur && h.key[left(i)]  > h.key[i] {
    max = left(i)
  }
  if right(i) <= h.cur && h.key[right(i)] > h.key[max] {
    max = right(i)
  }
  if max != i {
    swap(h, i, max)
    heapify(h, max)
  }
}

func swap[K constraints.Integer, V any](h *Heap[K,V], i, j int) {
  h.key[i], h.key[j] = h.key[j], h.key[i]
  h.val[i], h.val[j] = h.val[j], h.val[i]
}

func left(index int) int {
  return  index * 2 + 1
}

func right(index int) int {
  return  index * 2 + 2
}

func parent(index int) int {
  return  (index - 1) / 2
}

