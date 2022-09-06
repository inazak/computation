package heap

import (
  "testing"
)

func TestPushAndPop1(t *testing.T) {

  heap := NewHeap[int,string]()

  heap.Push(1, "A")
  heap.Push(9, "I")
  heap.Push(2, "B")
  heap.Push(8, "H")
  heap.Push(3, "C")
  heap.Push(7, "G")
  heap.Push(4, "D")
  heap.Push(6, "F")
  heap.Push(5, "E")


  expect := []struct{
    key int
    val string
  }{
    { key: 9, val: "I", },
    { key: 8, val: "H", },
    { key: 7, val: "G", },
    { key: 6, val: "F", },
    { key: 5, val: "E", },
    { key: 4, val: "D", },
    { key: 3, val: "C", },
    { key: 2, val: "B", },
    { key: 1, val: "A", },
  }

  for i, e := range expect {
    k, v := heap.Pop()
    if k != e.key || v != e.val {
      t.Errorf("no.%d got=%v, %v", i, k, v)
    }
  }

  if ! heap.Empty() {
    t.Errorf("heap is not empty")
  }
}

func TestPushAndPop2(t *testing.T) {

  heap := NewHeap[int,string]()

  //Push
  heap.Push(1, "A")
  heap.Push(2, "B")
  heap.Push(9, "I")

  expect1 := []struct{
    key int
    val string
  }{
    { key: 9, val: "I", },
    { key: 2, val: "B", },
    { key: 1, val: "A", },
  }

  //Pop
  for i, e := range expect1 {
    k, v := heap.Pop()
    if k != e.key || v != e.val {
      t.Errorf("expect1 no.%d got=%v, %v", i, k, v)
    }
  }

  //Push
  heap.Push(8, "H")
  heap.Push(7, "G")
  heap.Push(3, "C")

  expect2 := []struct{
    key int
    val string
  }{
    { key: 8, val: "H", },
    { key: 7, val: "G", },
    { key: 3, val: "C", },
  }

  //Pop
  for i, e := range expect2 {
    k, v := heap.Pop()
    if k != e.key || v != e.val {
      t.Errorf("expect2 no.%d got=%v, %v", i, k, v)
    }
  }

  //Push
  heap.Push(4, "D")
  heap.Push(6, "F")
  heap.Push(5, "E")

  expect3 := []struct{
    key int
    val string
  }{
    { key: 6, val: "F", },
    { key: 5, val: "E", },
    { key: 4, val: "D", },
  }

  //Pop
  for i, e := range expect3 {
    k, v := heap.Pop()
    if k != e.key || v != e.val {
      t.Errorf("expect3 no.%d got=%v, %v", i, k, v)
    }
  }

  if ! heap.Empty() {
    t.Errorf("heap is not empty")
  }
}


