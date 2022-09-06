package avltree

type Key   int
type Value string

const Nothing Value = ""

func (k Key) LessThan(a Key) bool {
  return k < a
}

func (k Key) Equal(a Key) bool {
  return k == a
}

type node struct {
  key    Key
  value  Value
  left   *node
  right  *node
  height int
}

type Tree struct {
  root *node
}

func NewTree() *Tree {
  return &Tree{}
}

func (t *Tree) Insert(k Key, v Value) {
  if t.root == nil {
    t.root = &node{ key: k, value: v }
  } else {
    t.root = insert(t.root, k, v)
  }
}

func insert(n *node, k Key, v Value) *node {

  if k.LessThan(n.key) {
    if n.left == nil {
      n.left = &node{ key: k, value: v, height: 0 }
    } else {
      n.left = insert(n.left, k, v) // recursion
    }
  } else {
    if n.right == nil {
      n.right = &node{ key: k, value: v, height: 0 }
    } else {
      n.right = insert(n.right, k, v) // recursion
    }
  }

  switch n.calcDifference() {
  case 2:
    if k.LessThan(n.left.key) {
      //     [n]        [a] ... is new n
      //   [a]    =>  [b] [n]
      // [b]
      n = rotateRight(n)
    } else {
      //    [n]          [n]        [b] ... is new n
      //  [a]    =>    [b]    =>  [a] [n]
      //    [b]      [a]
      n = rotateLeftRight(n)
    }
  case -2:
    if k.LessThan(n.right.key) {
      //  [n]        [n]            [b] ... is new n
      //    [a]  =>    [b]    =>  [n] [a]
      //  [b]            [a]
      n = rotateRightLeft(n)
    } else {
      // [n]            [a] ... is new n
      //   [a]    =>  [n] [b]
      //     [b]
      n = rotateLeft(n)
    }
  }

  n.updateHeight()
  return n
}

func (n *node) calcDifference() int {

  left  := 0
  right := 0

  if n.left != nil {
    left = n.left.height + 1
  }
  if n.right != nil {
    right = n.right.height + 1
  }

  return left - right
}

func (n *node) updateHeight() {

  height := -1

  if n.left != nil {
    height = n.left.height
  }
  if n.right != nil && height < n.right.height {
    height = n.right.height
  }

  n.height = height + 1
}

func rotateRight(n *node) *node {
  newroot := n.left
  n.left = newroot.right
  newroot.right = n

  n.updateHeight()
  // newroot.updateHeight() in insert function
  return newroot
}

func rotateLeftRight(n *node) *node {
  newroot := n.left.right
  n.left.right = newroot.left
  newroot.left = n.left
  n.left = newroot.right
  newroot.right = n

  newroot.left.updateHeight()
  n.updateHeight()
  // newroot.updateHeight() in insert function
  return newroot
}

func rotateLeft(n *node) *node {
  newroot := n.right
  n.right = newroot.left
  newroot.left = n

  n.updateHeight()
  // newroot.updateHeight() in insert function
  return newroot
}

func rotateRightLeft(n *node) *node {
  newroot := n.right.left
  n.right.left = newroot.right
  newroot.right = n.right
  n.right = newroot.left
  newroot.left = n

  newroot.right.updateHeight()
  n.updateHeight()
  // newroot.updateHeight() in insert function
  return newroot
}


func (t *Tree) Get(k Key) (Value, bool) {
  n := t.root
  for {
    if n == nil {
      return Nothing, false
    }
    if k.Equal(n.key) {
      return n.value, true
    }
    if k.LessThan(n.key) {
      n = n.left
    } else {
      n = n.right
    }
  }
}

func (t *Tree) Contain(k Key) bool {
  _, ok := t.Get(k)
  return ok
}


func (t *Tree) Remove(k Key) (ok bool) {

  if t.root == nil {
    return false
  }

  ok, t.root = remove(t.root, k)
  return ok
}

func remove(n *node, k Key) (bool, *node) {

  var ok bool

  if n == nil {
    return false, n
  }

  if k.Equal(n.key) {

    if n.left == nil {
      return true, n.right

    } else if n.right == nil {
      return true, n.left

    } else if n.left != nil && n.left.right == nil {
      n.left.right = n.right
      n = n.left
      ok = true

    } else {
      target := rightmost(n.left)
      _, n.left = remove(n.left, target.key) //another remove recursion
      target.left = n.left
      target.right = n.right
      n = target
      ok = true
    }

  } else {

    if k.LessThan(n.key) {
      ok, n.left = remove(n.left, k)
    } else {
      ok, n.right = remove(n.right, k)
    }
  }

  if ok {

    switch n.calcDifference() {
    case 2:
      if n.left.calcDifference() >= 0 {
        n = rotateRight(n)
      } else {
        n = rotateLeftRight(n)
      }
    case -2:
      if n.right.calcDifference() >= 0 {
        n = rotateRightLeft(n)
      } else {
        n = rotateLeft(n)
      }
    }

    n.updateHeight()
  }

  return ok, n
}

func rightmost(n *node) *node {
  if n.right == nil {
    return n
  }
  return rightmost(n.right)
}


type Item struct {
  Key   Key
  Value Value
}

func (t *Tree) ToList() []Item {
  list := []Item{}

  if t.root != nil {
    tolist(t.root, &list)
  }
  return list
}

func tolist(n *node, list *[]Item) {
  if n.left != nil {
    tolist(n.left, list)
  }

  *list = append(*list, Item{ Key: n.key, Value: n.value })

  if n.right != nil {
    tolist(n.right, list)
  }
}

