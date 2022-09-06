package dfa

import (
  "fmt"
  "reflect" //DeepEqual
)

// NFAをDFAに変換する考え方は単純。
// NFAは同一のシンボルでも、複数の遷移先がある、というもの。
// では複数のNFAをまとめた集合で考えた場合はどうなるか。
// NFAの集合はシンボルを取ったら、やっぱり複数の遷移先があるが、
// シンボルでグルーピングすることはできるようになる。

// たとえば、
// 「'a' を取ったら 1 に遷移するNFA」と、
// 「'a' を取ったら 2 に遷移するNFA」を一つの集合に入れた場合、
// 「'a' を取ったら [1, 2] のNFAの集合に遷移する」というように集約できる。
// シンボルが集約されるということは、シンボルと遷移先が1対1になるので、つまりDFAになる。


// NFANodeの集合からそれらが取りうるruneのセットを作成する
type RuneSet map[rune]struct{}

func (ns NFANodeSet) GetRuneSet() RuneSet {
  set := make(RuneSet)
  for node, _ := range ns {
    for key, _ := range node.transition {
      set[key] = struct{}{}
    }
  }
  return set
}

// NFANodeの集合から、「シンボル（Rune）と次の遷移先の組み」の集合を返す
// つまり、あるNFANodeの集合が取りうるシンボルが何かと、
// それを取った場合、どのNFANodeの集合に遷移するかを網羅する
func (ns NFANodeSet) GetTransitionPattern() map[rune]NFANodeSet {
  pattern := make(map[rune]NFANodeSet)

  for r, _ := range ns.GetRuneSet() {
    set := ns.GetTransition(r)
    for node, _ := range set.GetFreemove() {
      set.Add(node)
    }
    pattern[r] = set
  }
  return pattern
}


// NFAをDFAに変換する過程で、NFANodeの集合に対する
// 遷移を順番に辿っていく必要がある
// 同じNFANodeSetには同じDFANodeを割り当てるため
// その組み合わせを記録しておく必要がある
// また循環が発生していると終わらなくなるので、
// 処理済みのNodeを参照することになる
type Visited map[*NFANodeSet]*DFANode

func (v Visited) Exist(ns NFANodeSet) (*DFANode, bool) {
  for nfaset, dfa := range v {
    if reflect.DeepEqual(ns, *nfaset) {
      return dfa, true
    }
  }
  return nil, false
}

// 自動でDFANodeに名前をつけるため
var nodeNumber int
func initNodeNumber() {
  nodeNumber = 0
}

// helper function
// ToDFAメソッド内でDFANodeを自動生成するために使う
func MakeDFANodeWithNameAndAcceptable(ns NFANodeSet) *DFANode {
  name := fmt.Sprintf("s%d", nodeNumber)
  nodeNumber += 1
  node := MakeDFANode(name)
  if ns.HasAcceptNode() {
    node.SetAcceptNode(true)
  }
  return node
}


func (n *NFANode) ToDFA() *DFANode {

  // 遷移元となるNFANodeSetと、
  // 遷移先となるNFANodeSetの組みをまとめておく
  type trans struct {
    from   NFANodeSet
    to     NFANodeSet
    symbol rune
  }

  // 開始NFANodeSetを作成する
  startNFANodeSet := make(NFANodeSet)
  startNFANodeSet.Add(n)
  for node, _ := range n.GetFreemove() {
    startNFANodeSet.Add(node)
  }

  // 対応するDFANodeを作成する
  // NFANodeSetが受理状態であれば、DFANodeも受理状態にする
  initNodeNumber()
  startDFANode := MakeDFANodeWithNameAndAcceptable(startNFANodeSet)

  // 開始ノードの遷移先をまずキューに入れる
  queue := []trans{}
  for r, ns := range startNFANodeSet.GetTransitionPattern() {
    queue = append(queue, trans{ from: startNFANodeSet, to: ns, symbol: r })
  }

  // 処理済みの遷移先として記録しておく
  visited := make(Visited)
  visited[&startNFANodeSet] = startDFANode

  // ここからキューが無くなるまで遷移先の構築を繰り返す
  for {
    if len(queue) == 0 {
      break
    }

    // 先頭を取り出し
    tr   := queue[0]
    queue = queue[1:]

    // NFANodeSetに対応するDFANodeを検索する
    from, _ := visited.Exist(tr.from)
    to,  ok := visited.Exist(tr.to)

    // 遷移先のNFANodeSet(tr.to)が、すでに記録されている場合は
    // 記録されているDFANode(to)を使って構築
    if ok {
      from.AddTransition(tr.symbol, to)

    // なかった場合はDFANodeを新しく作って記録
    } else {
      dfanode := MakeDFANodeWithNameAndAcceptable(tr.to)
      from.AddTransition(tr.symbol, dfanode)

      // 最初と同じように遷移先をキューに入れる
      for r, ns := range tr.to.GetTransitionPattern() {
        queue = append(queue, trans{ from: tr.to, to: ns, symbol: r })
      }

      // 処理済みに追加する
      visited[&tr.to] = dfanode
    }

  } // キューが空になるまで続く

  return startDFANode
}


