package avltree

import (
  "testing"
)

func TestInsert1(t *testing.T) {

  SampleData := []struct {
    key   Key
    value Value
  }{
    { key: 1, value: "A" },
    { key: 2, value: "B" },
    { key: 3, value: "C" },
    { key: 4, value: "D" },
    { key: 5, value: "E" },
    { key: 6, value: "F" },
    { key: 7, value: "G" },
    { key: 8, value: "H" },
    { key: 9, value: "I" },
  }

  tree := NewTree()
  for _, e := range SampleData {
    tree.Insert( e.key, e.value )
  }
  //tree.Print()

  // tree shape is ...
  //         4
  //     2       6
  //   1   3   5   8
  //              7 9

  node := tree.root
  if node == nil {
    t.Fatalf("expected=4,D got=nil")
  }
  if node.key != 4 || node.value != "D" {
    t.Errorf("expected=4,D got=%v,%v", node.key, node.value)
  }

  node = tree.root.left
  if node == nil {
    t.Fatalf("expected=2,B got=nil")
  }
  if node.key != 2 || node.value != "B" {
    t.Errorf("expected=2,B got=%v,%v", node.key, node.value)
  }

  node = tree.root.right
  if node == nil {
    t.Fatalf("expected=6,F got=nil")
  }
  if node.key != 6 || node.value != "F" {
    t.Errorf("expected=6,F got=%v,%v", node.key, node.value)
  }

  node = tree.root.left.left
  if node == nil {
    t.Fatalf("expected=1,A got=nil")
  }
  if node.key != 1 || node.value != "A" {
    t.Errorf("expected=1,A got=%v,%v", node.key, node.value)
  }

  node = tree.root.left.right
  if node == nil {
    t.Fatalf("expected=3,C got=nil")
  }
  if node.key != 3 || node.value != "C" {
    t.Errorf("expected=3,C got=%v,%v", node.key, node.value)
  }

  node = tree.root.right.left
  if node == nil {
    t.Fatalf("expected=5,E got=nil")
  }
  if node.key != 5 || node.value != "E" {
    t.Errorf("expected=5,E got=%v,%v", node.key, node.value)
  }

  node = tree.root.right.right
  if node == nil {
    t.Fatalf("expected=8,H got=nil")
  }
  if node.key != 8 || node.value != "H" {
    t.Errorf("expected=8,H got=%v,%v", node.key, node.value)
  }

  node = tree.root.right.right.left
  if node == nil {
    t.Fatalf("expected=7,G got=nil")
  }
  if node.key != 7 || node.value != "G" {
    t.Errorf("expected=7,G got=%v,%v", node.key, node.value)
  }

  node = tree.root.right.right.right
  if node == nil {
    t.Fatalf("expected=9,I got=nil")
  }
  if node.key != 9 || node.value != "I" {
    t.Errorf("expected=9,I got=%v,%v", node.key, node.value)
  }
}

func TestInsert2(t *testing.T) {

  SampleData := []struct {
    key   Key
    value Value
  }{
    { key: 9, value: "I" },
    { key: 8, value: "H" },
    { key: 7, value: "G" },
    { key: 6, value: "F" },
    { key: 5, value: "E" },
    { key: 4, value: "D" },
    { key: 3, value: "C" },
    { key: 2, value: "B" },
    { key: 1, value: "A" },
  }

  tree := NewTree()
  for _, e := range SampleData {
    tree.Insert( e.key, e.value )
  }
  //tree.Print()

  // tree shape is ...
  //         6
  //     4       8
  //   2   5   7   9
  //  1 3

  node := tree.root
  if node == nil {
    t.Fatalf("expected=6,F got=nil")
  }
  if node.key != 6 || node.value != "F" {
    t.Errorf("expected=6,F got=%v,%v", node.key, node.value)
  }

  node = tree.root.left
  if node == nil {
    t.Fatalf("expected=4,D got=nil")
  }
  if node.key != 4 || node.value != "D" {
    t.Errorf("expected=4,D got=%v,%v", node.key, node.value)
  }

  node = tree.root.right
  if node == nil {
    t.Fatalf("expected=8,H got=nil")
  }
  if node.key != 8 || node.value != "H" {
    t.Errorf("expected=8,H got=%v,%v", node.key, node.value)
  }

  node = tree.root.left.left
  if node == nil {
    t.Fatalf("expected=2,B got=nil")
  }
  if node.key != 2 || node.value != "B" {
    t.Errorf("expected=2,B got=%v,%v", node.key, node.value)
  }

  node = tree.root.left.right
  if node == nil {
    t.Fatalf("expected=5,E got=nil")
  }
  if node.key != 5 || node.value != "E" {
    t.Errorf("expected=5,E got=%v,%v", node.key, node.value)
  }

  node = tree.root.right.left
  if node == nil {
    t.Fatalf("expected=7,G got=nil")
  }
  if node.key != 7 || node.value != "G" {
    t.Errorf("expected=7,G got=%v,%v", node.key, node.value)
  }

  node = tree.root.right.right
  if node == nil {
    t.Fatalf("expected=9,I got=nil")
  }
  if node.key != 9 || node.value != "I" {
    t.Errorf("expected=9,I got=%v,%v", node.key, node.value)
  }

  node = tree.root.left.left.left
  if node == nil {
    t.Fatalf("expected=1,A got=nil")
  }
  if node.key != 1 || node.value != "A" {
    t.Errorf("expected=1,A got=%v,%v", node.key, node.value)
  }

  node = tree.root.left.left.right
  if node == nil {
    t.Fatalf("expected=3,C got=nil")
  }
  if node.key != 3 || node.value != "C" {
    t.Errorf("expected=3,C got=%v,%v", node.key, node.value)
  }
}

func TestInsert3(t *testing.T) {

  SampleData := []struct {
    key   Key
    value Value
  }{
    { key: 1, value: "A" },
    { key: 2, value: "B" },
    { key: 5, value: "E" }, // left rotate
    { key: 4, value: "D" },
    { key: 6, value: "F" },
    { key: 3, value: "C" }, // right left rotate
  }

  tree := NewTree()
  for _, e := range SampleData {
    tree.Insert( e.key, e.value )
  }
  //tree.Print()

  // tree shape is ...
  //     4
  //   2   5
  //  1 3   6

  node := tree.root
  if node == nil {
    t.Fatalf("expected=4,D got=nil")
  }
  if node.key != 4 || node.value != "D" {
    t.Errorf("expected=4,D got=%v,%v", node.key, node.value)
  }

  node = tree.root.left
  if node == nil {
    t.Fatalf("expected=2,B got=nil")
  }
  if node.key != 2 || node.value != "B" {
    t.Errorf("expected=2,B got=%v,%v", node.key, node.value)
  }

  node = tree.root.right
  if node == nil {
    t.Fatalf("expected=5,E got=nil")
  }
  if node.key != 5 || node.value != "E" {
    t.Errorf("expected=5,E got=%v,%v", node.key, node.value)
  }

  node = tree.root.left.left
  if node == nil {
    t.Fatalf("expected=1,A got=nil")
  }
  if node.key != 1 || node.value != "A" {
    t.Errorf("expected=1,A got=%v,%v", node.key, node.value)
  }

  node = tree.root.left.right
  if node == nil {
    t.Fatalf("expected=3,C got=nil")
  }
  if node.key != 3 || node.value != "C" {
    t.Errorf("expected=3,C got=%v,%v", node.key, node.value)
  }

  node = tree.root.right.right
  if node == nil {
    t.Fatalf("expected=6,F got=nil")
  }
  if node.key != 6 || node.value != "F" {
    t.Errorf("expected=6,F got=%v,%v", node.key, node.value)
  }
}

func TestInsert4(t *testing.T) {

  SampleData := []struct {
    key   Key
    value Value
  }{
    { key: 6, value: "F" },
    { key: 5, value: "E" },
    { key: 2, value: "B" }, // right rotate
    { key: 3, value: "C" },
    { key: 1, value: "A" },
    { key: 4, value: "D" }, // left right rotate
  }

  tree := NewTree()
  for _, e := range SampleData {
    tree.Insert( e.key, e.value )
  }
  //tree.Print()

  // tree shape is ...
  //         3
  //     2       5
  //   1       4   6

  node := tree.root
  if node == nil {
    t.Fatalf("expected=3,C got=nil")
  }
  if node.key != 3 || node.value != "C" {
    t.Errorf("expected=3,C got=%v,%v", node.key, node.value)
  }

  node = tree.root.left
  if node == nil {
    t.Fatalf("expected=2,B got=nil")
  }
  if node.key != 2 || node.value != "B" {
    t.Errorf("expected=2,B got=%v,%v", node.key, node.value)
  }

  node = tree.root.right
  if node == nil {
    t.Fatalf("expected=5,E got=nil")
  }
  if node.key != 5 || node.value != "E" {
    t.Errorf("expected=5,E got=%v,%v", node.key, node.value)
  }

  node = tree.root.left.left
  if node == nil {
    t.Fatalf("expected=1,A got=nil")
  }
  if node.key != 1 || node.value != "A" {
    t.Errorf("expected=1,A got=%v,%v", node.key, node.value)
  }

  node = tree.root.right.left
  if node == nil {
    t.Fatalf("expected=4,D got=nil")
  }
  if node.key != 4 || node.value != "D" {
    t.Errorf("expected=4,D got=%v,%v", node.key, node.value)
  }

  node = tree.root.right.right
  if node == nil {
    t.Fatalf("expected=6,F got=nil")
  }
  if node.key != 6 || node.value != "F" {
    t.Errorf("expected=6,F got=%v,%v", node.key, node.value)
  }
}


func TestGet(t *testing.T) {

  SampleData := []struct {
    key   Key
    value Value
  }{
    { key: 9, value: "I" },
    { key: 8, value: "H" },
    { key: 7, value: "G" },
    { key: 6, value: "F" },
    { key: 5, value: "E" },
    { key: 4, value: "D" },
    { key: 3, value: "C" },
    { key: 2, value: "B" },
    { key: 1, value: "A" },
  }

  tree := NewTree()
  for _, e := range SampleData {
    tree.Insert( e.key, e.value )
  }

  v1, ok := tree.Get(3)
  if !ok || v1 != "C" {
    t.Errorf("expected=%v,%v got=%v,%v", true, "C", ok, v1)
  }

  v2, ok := tree.Get(999)
  if ok || v2 != Nothing {
    t.Errorf("expected=%v,%v got=%v,%v", false, Nothing, ok, v2)
  }
}


func TestRemove1(t *testing.T) {

  SampleData := []struct {
    key   Key
    value Value
  }{
    { key: 1, value: "A" },
    { key: 2, value: "B" },
    { key: 3, value: "C" },
    { key: 4, value: "D" },
    { key: 5, value: "E" },
    { key: 6, value: "F" },
    { key: 7, value: "G" },
    { key: 8, value: "H" },
    { key: 9, value: "I" },
  }

  tree := NewTree()
  for _, e := range SampleData {
    tree.Insert( e.key, e.value )
  }
  //tree.Print()

  // tree shape is ...
  //         4
  //     2       6
  //   1   3   5   8
  //              7 9

  var ok bool

  if ok = tree.Remove(4) ; !ok {
    t.Fatal("remove fail in 4")
  }
  if ok = tree.Remove(3) ; !ok {
    t.Fatal("remove fail in 3")
  }

  // tree shape is ...
  //         5
  //     2       6
  //   1           8
  //              7 9

  if ok = tree.Remove(5) ; !ok {
    t.Fatal("remove fail in 5")
  }

  // tree shape is ...
  //         6
  //     2       8
  //   1        7  9

  if ok = tree.Remove(6) ; !ok {
    t.Fatal("remove fail in 6")
  }
  if ok = tree.Remove(2) ; !ok {
    t.Fatal("remove fail in 2")
  }

  // tree shape is ...
  //         7
  //     1       8
  //               9

  if ok = tree.Remove(7) ; !ok {
    t.Fatal("remove fail in 7")
  }
  if ok = tree.Remove(8) ; !ok {
    t.Fatal("remove fail in 8")
  }
  if ok = tree.Remove(1) ; !ok {
    t.Fatal("remove fail in 1")
  }
  if ok = tree.Remove(9) ; !ok {
    t.Fatal("remove fail in 9")
  }

  if tree.root != nil {
    t.Errorf("tree.root expected nil, but got=%v", toString(tree.root))
  }

}

func TestRemove2(t *testing.T) {

  SampleData := []struct {
    key   Key
    value Value
  }{
    { key: 9, value: "I" },
    { key: 8, value: "H" },
    { key: 7, value: "G" },
    { key: 6, value: "F" },
    { key: 5, value: "E" },
    { key: 3, value: "C" },
    { key: 2, value: "B" },
    { key: 1, value: "A" },
    { key: 4, value: "D" },
  }

  tree := NewTree()
  for _, e := range SampleData {
    tree.Insert( e.key, e.value )
  }

  // tree shape is ...
  //         6
  //     3       8
  //   2   5   7   9  ... 5 is left.rightmost()
  //  1   4

  if ok := tree.Remove(6) ; !ok {
    t.Fatal("remove fail in 6")
  }

  // tree shape is ...
  //         5
  //     3       8
  //   2   4   7   9
  //  1
}

func TestToList(t *testing.T) {

  SampleData := []struct {
    key   Key
    value Value
  }{
    { key: 9, value: "I" },
    { key: 1, value: "A" },
    { key: 8, value: "H" },
    { key: 2, value: "B" },
    { key: 7, value: "G" },
    { key: 3, value: "C" },
    { key: 6, value: "F" },
    { key: 4, value: "D" },
    { key: 5, value: "E" },
  }

  tree := NewTree()
  for _, e := range SampleData {
    tree.Insert( e.key, e.value )
  }

  expect := []struct {
    key   Key
    value Value
  }{
    { key: 1, value: "A" },
    { key: 2, value: "B" },
    { key: 3, value: "C" },
    { key: 4, value: "D" },
    { key: 5, value: "E" },
    { key: 6, value: "F" },
    { key: 7, value: "G" },
    { key: 8, value: "H" },
    { key: 9, value: "I" },
  }

  list := tree.ToList()
  for i, e := range list {
    if e.Key != expect[i].key || e.Value != expect[i].value {
      t.Errorf("no.%d got %v:%v", i, e.Key, e.Value)
    }
  }
}

