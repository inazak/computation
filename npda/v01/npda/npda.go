package npda

import (
  "fmt"
  "strings"
  "reflect" //DeepEqual
)

// 遷移条件のルールは二通りある
// a) 入力文字（Symbol）とスタックの先頭文字（StackTop）が両方合致した場合
// b) スタックの先頭文字（StackTop）が合致した場合
//
// つねに a) をまず検索し、その遷移処理後の状態を対象に b) も検索する
// その際、b) の遷移先がなくなるまで繰り返し検索する必要がある
//
// a) のルールをNPDAの集合に適用する場合は、
// ルールに合致しなかった状態のインスタンスは削除される
// 入力文字に合致しなかったということになるので
// 
// b) のルールをNPDAの集合に適用する場合は、
// ルールに合致しなかった状態のインスタンスも保持される
// 現在のスタックの状態に対するFreemoveという考え方なので
//
// 検索条件に合致した場合は、状態を次のノード（NextNode）に切り替え、
// かつ スタックから一つ取り出し（Pop）、指定されたシンボルをスタックに積む（PushSymbol）

type TranRule struct {
  Symbol     rune
  StackTop   rune
  NextNode   *NPDANode
  PushSymbol []rune
}

// NPDAノードは二つのルールを持っている、先ほどの
// a) に当たるものが tranRuleOfSymbolAndStack
// b) に当たるものが tranRuldOfStack
type NPDANode struct {
  acceptNode bool
  tranRuleOfSymbolAndStack []TranRule
  tranRuldOfStack          []TranRule
}

func MakeNPDANode() *NPDANode {
  return &NPDANode{}
}

func (n *NPDANode) IsAcceptNode() bool {
  return n.acceptNode
}

func (n *NPDANode) SetAcceptNode(b bool) {
  n.acceptNode = b
}

func (n *NPDANode) AddTranRuleOfSymbolAndStack(t TranRule) {
  n.tranRuleOfSymbolAndStack = append(n.tranRuleOfSymbolAndStack, t)
}

func (n *NPDANode) GetTranRuleOfSymbolAndStack() []TranRule {
  return n.tranRuleOfSymbolAndStack
}

func (n *NPDANode) AddTranRuleOfStack(t TranRule) {
  n.tranRuldOfStack = append(n.tranRuldOfStack, t)
}

func (n *NPDANode) GetTranRuleOfStack() []TranRule {
  return n.tranRuldOfStack
}


// 実際にNPDAの動作をシミュレートする際には、それぞれの状態に
// 遷移した際のスタックの情報が必要になる
//
// NPDAでは複数の状態を管理することになるので、一つの共有のスタックを
// 使うということはできず、遷移するたびに遷移元のスタックを引き継いで
// 新しいインスタンスを作成することになる
//

type NPDAInstance struct {
  node  *NPDANode
  stack []rune
}

func MakeNPDAInstance(n *NPDANode, stack []rune) *NPDAInstance {
  return &NPDAInstance{
    node: n,
    stack: stack,
  }
}

func (ins *NPDAInstance) StackTop() rune {
  return ins.stack[len(ins.stack)-1]
}

func (ins *NPDAInstance) StackPop() rune {
  p := ins.stack[len(ins.stack)-1]
  ins.stack = ins.stack[:len(ins.stack)-1]
  return p
}

func (ins *NPDAInstance) StackPush(r rune) {
  ins.stack = append(ins.stack, r)
}

// インスタンスに対して、
// a) 入力文字（Symbol）とスタックの先頭文字（StackTop）が両方合致した場合
// に合致するルールがないか検索する
func (ins *NPDAInstance) GetTranRuleOfSymbolAndStack(symbol rune) []TranRule {
  result := []TranRule{}

  for _, rule := range ins.node.GetTranRuleOfSymbolAndStack() {
    if rule.Symbol == symbol && rule.StackTop == ins.StackTop() {
      result = append(result, rule)
    }
  }
  return result
}

// インスタンスに対して、
// b) スタックの先頭文字（StackTop）が合致した場合
// に合致するルールがないか検索する
func (ins *NPDAInstance) GetTranRuleOfStack() []TranRule {
  result := []TranRule{}

  for _, rule := range ins.node.GetTranRuleOfStack() {
    if rule.StackTop == ins.StackTop() {
      result = append(result, rule)
    }
  }
  return result
}

func (ins *NPDAInstance) HasTranRuleOfStack() bool {
  return len(ins.GetTranRuleOfStack()) > 0
}

// 合致したルールを引数にして、新しい遷移先のインスタンスを作成する
// スタックを引き継ぐためにコピーしている
// 
// 状態を次のノード（NextNode）に切り替え、
// かつ スタックから一つ取り出し（Pop）、指定されたシンボルをスタックに積む（PushSymbol）
func (ins *NPDAInstance) MakeNextInstance(t TranRule) *NPDAInstance {

  newstack := []rune{}
  for _, r := range ins.stack {
    newstack = append(newstack, r)
  }

  newins := MakeNPDAInstance(t.NextNode, newstack)

  _ = newins.StackPop()
  for _, s := range t.PushSymbol {
    newins.StackPush(s)
  }
  return newins
}


// NPDAでシミュレーションをする場合は、一つの状態から「複数の状態の集合」への遷移がありうる
// 現在の状態をまとめて管理するため、インスタンスの集合（Set）を用意する
type NPDAInstanceSet []*NPDAInstance

func (nis NPDAInstanceSet) HasAcceptNode() bool {
  for _, ins := range nis {
    if ins.node.IsAcceptNode() {
      return true
    }
  }
  return false
}

func (nis NPDAInstanceSet) HasTranRuleOfStack() bool {
  for _, ins := range nis {
    if ins.HasTranRuleOfStack() {
      return true
    }
  }
  return false
}

// 追加の際に、既に存在するインスタンスと同値であれば無視する
func (nis NPDAInstanceSet) Add(ins *NPDAInstance) NPDAInstanceSet {
  for _, i := range nis {
    if reflect.DeepEqual(i, ins) {
      return nis
    }
  }
  return append(nis, ins)
}

// インスタンスの集合（Set）に対して、検索条件である
// b) スタックの先頭文字（StackTop）が合致した場合
// が適合するものがあれば、そのルールに従い新しいインスタンスを生成し
// 現在の集合を更新したものを返す
func (nis NPDAInstanceSet) MakeNextInstanceSetForRuleOfStack() NPDAInstanceSet  {

  next := NPDAInstanceSet{}

  for _, ins := range nis {
    if ins.HasTranRuleOfStack() {
      for _, rule := range ins.GetTranRuleOfStack() {
        newins := ins.MakeNextInstance(rule)
        next = next.Add(newins)
      }
    } else {
      next = next.Add(ins)
    }
  }

  return next
}



// NPDAによるルール適用のシミュレーション
//
// 開始ノードでこのメソッドを呼ぶことで
// 文字列を読んでNPDAInstanceを生成しながら
// NPDANodeを辿り、受理状態になるかどうかを返す
//
// 途中で遷移先のNPDANodeがなくなった場合はfalseかつerrorを返す
//
func (n *NPDANode) Accept(s string) (bool, error) {
  return n.AcceptWithDebugPrint(s, false)
}

func (n *NPDANode) AcceptWithDebugPrint(s string, debug bool) (bool, error) {

  // 初期状態を作成する
  // スタックの末尾を示す '$' を設定する
  // 現在の集合（Set）を作成する
  start := MakeNPDAInstance(n, []rune{'$'})
  set := NPDAInstanceSet{}
  set = set.Add(start)

  // スタックの先頭文字に対する遷移先がある限り、
  // 新しいインスタンスを作成し、現在の集合（Set）を更新する
  for set.HasTranRuleOfStack() {
    set = set.MakeNextInstanceSetForRuleOfStack()
  }

  if debug {
    fmt.Printf("--------------\n")
    fmt.Printf("Start:\n")
    fmt.Printf(set.ToString())
  }

  // 文字列を一文字ずつ読み込む
  var err error
  for _, r := range s {

    // 現在のSetに対して、適用できるルールを検索し
    // 合致した場合は新しいインスタンスを作成して、集合（Set）を更新する
    set, err = set.ReadSymbol(r)
    if err != nil {
      return false, err
    }

    if debug {
      fmt.Printf("--------------\n")
      fmt.Printf("Read [%v]:\n", string(r))
      fmt.Printf(set.ToString())
    }
  }

  return set.HasAcceptNode(), nil
}



func (nis NPDAInstanceSet) ReadSymbol(r rune) (NPDAInstanceSet, error)  {

  // 遷移先がない場合
  if len(nis) == 0 {
    return nil, fmt.Errorf("there is no destination NPDANode")
  }

  next := NPDAInstanceSet{}
  for _, ins := range nis {

    // a) 入力文字（Symbol）とスタックの先頭文字（StackTop）が両方合致した場合
    for _, rule := range ins.GetTranRuleOfSymbolAndStack(r) {
      newins := ins.MakeNextInstance(rule)
      next = next.Add(newins)
    }

    // b) スタックの先頭文字（StackTop）が合致した場合
    for next.HasTranRuleOfStack() {
      next = next.MakeNextInstanceSetForRuleOfStack()
    }
  }

  return next, nil
}



// For debug
func (ins *NPDAInstance) ToString() string {
  node := fmt.Sprintf("%p", ins.node)
  node  = node[len(node)-4:]
  stk  := []string{}
  for i, _ := range ins.stack {
    stk = append(stk, string(ins.stack[len(ins.stack)-1-i]))
  }
  return fmt.Sprintf("Node(%v) Stack[%v]", node, strings.Join(stk,","))
}

// For debug
func (nis NPDAInstanceSet) ToString() string {
  s := ""
  for _, ins := range nis {
    s += ins.ToString() + "\n"
  }
  return s
}

