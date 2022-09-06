package dfa

import (
  "fmt"
)

// イプシロン遷移（自由移動）の遷移先も保持する
type NFANode struct {
  name       string
  acceptNode bool
  transition map[rune]NFANodeSet
  freemove   NFANodeSet
}

type NFANodeSet map[*NFANode]struct{}

func MakeNFANode(name string) *NFANode {
  return &NFANode{
    name:       name,
    acceptNode: false,
    transition: make(map[rune]NFANodeSet),
    freemove:   make(NFANodeSet),
  }
}

func (n *NFANode) IsAcceptNode() bool {
  return n.acceptNode
}

func (n *NFANode) SetAcceptNode(b bool) {
  n.acceptNode = b
}

func (n *NFANode) GetName() string {
  if n.IsAcceptNode() {
    return n.name + "*"
  }
  return n.name + " "
}

func (n *NFANode) AddTransition(r rune, to *NFANode) {
  if _, ok := n.transition[r]; ! ok {
    n.transition[r] = make(NFANodeSet)
  }
  n.transition[r].Add(to)
}

func (n *NFANode) GetTransition(r rune) NFANodeSet {
  if _, ok := n.transition[r]; ! ok {
    return nil
  }
  return n.transition[r]
}

func (n *NFANode) AddFreemove(to *NFANode) {
  n.freemove.Add(to)
}

func (n *NFANode) GetFreemove() NFANodeSet {
  set := make(NFANodeSet)

  // 遷移先にFreemoveがある可能性を考えて、再帰で辿る
  // ただしFreemoveが循環していると終わらない
  if len(n.freemove) > 0 {
    for node, _ := range n.freemove {
      set.Add(node)
      for next, _ := range node.GetFreemove() {
        set.Add(next)
      }
    }
  }
  return set
}

// NFANodeSet
func (ns NFANodeSet) Add(node *NFANode) {
  ns[node] = struct{}{}
}

func (ns NFANodeSet) GetTransition(r rune) NFANodeSet {
  set := make(NFANodeSet)
  for n, _ := range ns {
    for node, _ := range n.GetTransition(r) {
      set.Add(node)
    }
  }
  return set
}

func (ns NFANodeSet) GetFreemove() NFANodeSet {
  set := make(NFANodeSet)
  for n, _ := range ns {
    for node, _ := range n.GetFreemove() {
      set.Add(node)
    }
  }
  return set
}

// ノードの集合が入力を取って、次のノードの集合へ遷移するときは、
// 1) 遷移元,入力の組み合わせに従った、遷移先のみの集合を作る
// 2) 上記 1) の集合に含まれるノードに、イプシロン遷移（自由移動）が
//    あればそれを 1) の集合に追加する
//
// この際 1) の手順では遷移元ノードは削除されるが、
// 2) の手順では遷移元は削除しない
// ということが分かりにくいかもしれない
func (ns NFANodeSet) GetNextNFANodeSet(r rune) NFANodeSet {
  set := ns.GetTransition(r)
  for node, _ := range set.GetFreemove() {
    set.Add(node)
  }
  return set
}

func (ns NFANodeSet) HasAcceptNode() bool {
  for n, _ := range ns {
    if n.IsAcceptNode() {
      return true
    }
  }
  return false
}


// 開始ノードでこのメソッドを呼ぶことで
// 文字列を読んでNFANodeを辿り、受理状態になるかどうかを返す
// 途中で遷移先のNFANodeがなくなった場合はfalseかつerrorを返す
func (n *NFANode) Accept(s string) (bool, error) {

  set := make(NFANodeSet)
  set.Add(n)
  for node, _ := range n.GetFreemove() {
    set.Add(node)
  }

  for _, r := range s {
    set = set.GetNextNFANodeSet(r)

    if len(set) == 0 {
      // 該当する遷移先がない
      return false, fmt.Errorf("there is no destination NFANode")
    }
  }

  return set.HasAcceptNode(), nil
}


// NFANodeの遷移図を作る
func (n *NFANode) Dump() string {

  text    := ""
  visited := make(map[*NFANode]struct{})
  queue   := []*NFANode { n }

  for len(queue) > 0 {

    node := queue[0]
    queue = queue[1:]

    if _, ok := visited[node]; ok {
      continue
    }
    visited[node] = struct{}{}

    for r, set := range node.transition {
      for next, _ := range set {
        queue = append(queue, next)
        text += fmt.Sprintf("%s - %s -> %s\n", node.GetName(), string(r), next.GetName() )
      }
    }

    for next, _ := range node.freemove {
      queue = append(queue, next)
      text += fmt.Sprintf("%s -----> %s\n",  node.GetName(), next.GetName() )
    }
  }

  return text
}


