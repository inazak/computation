package avltree

import (
  "fmt"
  "strings"
)

func toString(n *node) string {
  if n == nil {
    return ""
  }
  return fmt.Sprintf("[%v:%v:%v]", n.key, n.value, n.height)
}

func allEntryIsNotNil(list []*node) bool {
  for _, e := range list {
    if e != nil {
      return true
    }
  }
  return false
}


func (t *Tree) Print() {

  if t.root == nil {
    return
  }

  if t.root.left == nil && t.root.right == nil {
    fmt.Printf("%v\n", toString(t.root))
    return
  }

  queue       := []*node{ t.root }
  next_queue  := []*node{}
  output      := []string{ "" }
  next_output := []string{}
  space       := 0
  next_space  := 0

  for allEntryIsNotNil(queue) {

    for _, node := range queue {

      if node == nil {
        next_queue = append(next_queue, nil, nil)
      } else {
        next_queue = append(next_queue, node.right, node.left)
      }

      s := strings.Repeat(" ", space) + toString(node)
      next_output = append(next_output, s)
      next_output = append(next_output, output[0])
      output = output[1:]

      if len(s) > next_space {
        next_space = len(s)
      }
    }

    queue  = next_queue
    output = next_output
    space  = next_space
    next_queue  = []*node{}
    next_output = []string{}
  }

  for _, o := range output {
    fmt.Printf("%s\n", o)
  }
}

