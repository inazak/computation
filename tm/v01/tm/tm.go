package tm

type TMNode struct {
  rule map[rune]TranRule
  name string
  acceptNode bool
}

type TranRule struct {
  Expect rune
  Write  rune
  Move   int
  Next   *TMNode
}

const (
  BLANK = '_' //undefined tape cell
)

const (
  LEFT  = 1
  RIGHT = 2
)


func MakeTMNode(name string) *TMNode {
  return &TMNode{
    rule: make(map[rune]TranRule),
    name: name,
    acceptNode: false,
  }
}

func (n *TMNode) SetAcceptNode(b bool) {
  n.acceptNode = b
}

func (n *TMNode) IsAcceptNode() bool {
  return n.acceptNode
}

func (n *TMNode) AddTranRule(t TranRule) {
  n.rule[t.Expect] = t
}

func (n *TMNode) GetTranRule(r rune) (TranRule, bool) {
  rule, ok := n.rule[r]
  return rule, ok
}

type Tape struct {
  head  rune
  left  []rune // the end of slice is the right edge
  right []rune // the end of slice is the left edge
}

func MakeBlankTape() *Tape {
  return MakeTape( nil, '_', nil )
}

func MakeTape(left []rune, head rune, right []rune) *Tape {
  for i,j := 0,len(right)-1; i<len(right)/2; i,j = i+1,j-1 {
    right[i], right[j] = right[j], right[i]
  }
  return &Tape{ head: head, left: left, right: right }
}

func (t *Tape) ReadHead() rune {
  return t.head
}

func (t *Tape) WriteHead(r rune) {
  t.head = r
}

func (t *Tape) MoveLeft() {
  t.right = append(t.right, t.head)
  if len(t.left) > 0 {
    t.head = t.left[len(t.left)-1]
    t.left = t.left[:len(t.left)-1]
  } else {
    t.head = '_'
  }
}

func (t *Tape) MoveRight() {
  t.left = append(t.left, t.head)
  if len(t.right) > 0 {
    t.head  = t.right[len(t.right)-1]
    t.right = t.right[:len(t.right)-1]
  } else {
    t.head = '_'
  }
}

func (t *Tape) Move(direction int) {
  if direction == LEFT {
    t.MoveLeft()
  } else if direction == RIGHT {
    t.MoveRight()
  } else {
    panic("Move: unknown direction")
  }
}


func (n *TMNode) HasNext(t *Tape) bool {
  head  := t.ReadHead()
  _, ok := n.GetTranRule(head)
  return ok
}

func (n *TMNode) Step(t *Tape) *TMNode {

  head := t.ReadHead()
  rule, ok := n.GetTranRule(head)

  if ! ok {
    return nil
  }

  t.WriteHead(rule.Write)
  t.Move(rule.Move)

  return rule.Next
}

func (n *TMNode) Run(t *Tape) bool {

  curr := n
  for {
    next := curr.Step(t)
    if next == nil {
      break
    }
    curr = next
  }

  return curr.IsAcceptNode()
}


// for debug
func (t *Tape) DumpLeft(size int) string {
  dump := make([]rune, size)
  for i := 0; i<size; i++ {
    if len(t.left) > i {
      dump[ size -i -1 ] = t.left[ len(t.left) -i -1 ]
    } else {
      dump[ size -i -1 ] = BLANK
    }
  }
  return string(dump)
}

func (t *Tape) DumpRight(size int) string {
  dump := make([]rune, size)
  for i := 0; i<size; i++ {
    if len(t.right) > i {
      dump[ i ] = t.right[ len(t.right) -i -1 ]
    } else {
      dump[ i ] = BLANK
    }
  }
  return string(dump)
}

func (t *Tape) DumpTape(size int) string {
  return t.DumpLeft(size) + "[" + string(t.head) + "]" + t.DumpRight(size)
}

