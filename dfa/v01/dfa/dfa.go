package dfa

import (
  "fmt"
)


type DFANode struct {
  name       string
  transition map[rune]*DFANode
  acceptNode bool
}

func MakeDFANode(name string) *DFANode {
  return &DFANode{
    name: name,
    acceptNode: false,
    transition: make(map[rune]*DFANode),
  }
}

func (n *DFANode) IsAcceptNode() bool {
  return n.acceptNode
}

func (n *DFANode) SetAcceptNode(b bool) {
  n.acceptNode = b
}

func (n *DFANode) GetName() string {
  if n.IsAcceptNode() {
    return n.name + "*"
  }
  return n.name + " "
}

func (n *DFANode) AddTransition(r rune, to *DFANode) {
  n.transition[r] = to
}

func (n *DFANode) GetTransition(r rune) *DFANode {
  return n.transition[r]
}


// 開始ノードでこのメソッドを呼ぶことで
// 文字列を読んでDFANodeを辿り、受理状態になるかどうかを返す
// 途中で遷移先のDFANodeがなくなった場合はfalseかつerrorを返す
func (n *DFANode) Accept(s string) (bool, error) {

  node := n
  for _, r := range s {
    node = node.GetTransition(r)

    if node == nil {
      // 該当する遷移先がない
      return false, fmt.Errorf("there is no destination DFANode")
    }
  }

  return node.IsAcceptNode(), nil
}


// DFANodeの遷移図を作る
func (n *DFANode) Dump() string {

  text    := ""
  visited := make(map[*DFANode]struct{})
  queue   := []*DFANode { n }

  for len(queue) > 0 {

    node := queue[0]
    queue = queue[1:]

    if _, ok := visited[node]; ok {
      continue
    }
    visited[node] = struct{}{}

    for r, next := range node.transition {
      queue = append(queue, next)
      text += fmt.Sprintf("%s - %s -> %s\n", node.GetName(), string(r), next.GetName() )
    }
  }

  return text
}


