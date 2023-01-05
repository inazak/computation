package model

const (
  HI = 1
  LO = 0
)

type Port struct {
  State   int // HI or LO
  update  func()
}

type Circuit struct {
  clock   *Port
  all     []*Port
  next    map[*Port][]*Port
  prev    map[*Port][]*Port
}

func (c *Circuit) appendPort(p *Port) {
  c.all = append(c.all, p)
}

func (c *Circuit) appendLink(from, to *Port) {
  c.next[from] = append(c.next[from], to)
  c.prev[to]   = append(c.prev[to], from)
}

func (c *Circuit) getAllPorts() []*Port {
  return c.all
}

func (c *Circuit) getNextPorts(p *Port) []*Port {
  if _, ok := c.next[p] ; ok {
    return c.next[p]
  }
  return []*Port{}
}

func (c *Circuit) getPrevPorts(p *Port) []*Port {
  if _, ok := c.prev[p] ; ok {
    return c.prev[p]
  }
  return []*Port{}
}

func MakeCircuit() *Circuit {
  c := &Circuit{}

  c.clock = &Port{
    State:  LO,
    update: func(){}, //do nothing
  }
  c.all  = []*Port{}
  c.next = make(map[*Port][]*Port)
  c.prev = make(map[*Port][]*Port)
  c.appendPort(c.clock)

  return c
}

func (c *Circuit) ClockUp() {
  c.clock.State = HI
}

func (c *Circuit) ClockDown() {
  c.clock.State = LO
}

func (c *Circuit) Cycle() {
  c.ClockDown()
  c.Update()
  c.ClockUp()
  c.Update()
}

func (c *Circuit) GetClock() *Port {
  return c.clock
}

func (c *Circuit) MakePort() *Port {
  p := &Port{
    State:  LO,
    update: func(){}, //do nothing
  }
  c.appendPort(p)
  return p
}


// Nand
//
// a ---> +------+
//        | NAND |------> out
// b +--> +------+ 
//
//  a   b   | out
//  --------|----
//  LO  LO  | HI
//  LO  HI  | HI
//  HI  LO  | HI
//  HI  HI  | LO
//
func (c *Circuit) Nand(a, b *Port) (out *Port) {
  out = &Port{
    State:  LO,
    update: func() { //closure
      if a.State == HI && b.State == HI {
        out.State = LO
      } else {
        out.State = HI
      }
    },
  }
  c.appendPort(out)
  c.appendLink(a, out)
  c.appendLink(b, out)

  return out
}

func (c *Circuit) Connect(from, to *Port) {

  if size := len(c.getPrevPorts(to)) ; size != 0 {
    panic("Connect() cant connect. line already has link.")
  }

  to.update = func() { //closure
    to.State = from.State
  }
  c.appendLink(from, to)
}

func (c *Circuit) PreUpdate() {
  c.update(true)
}

func (c *Circuit) Update() {
  c.update(false)
}

//sorry, it is actually poorly done.
func (c *Circuit) update(isPreUpdate bool) {
  queue     := []*Port{}
  shortloop := make(map[*Port]bool)
  updated   := make(map[*Port]bool)

  for _, p := range c.getAllPorts() {

    //setup first queue entries
    if c.clock == p || len(c.getPrevPorts(p)) == 0 {
      queue = append(queue, p)
    }

    shortloop[p] = c.hasShortLoop(p)
    updated[p] = ! isPreUpdate
  }

  for len(queue) > 0 {
    p := queue[0]
    queue = queue[1:] //pop

    prevstate := p.State
    p.update()

    if ! updated[p] || p.State != prevstate || len(c.getPrevPorts(p)) == 0 {

      updated[p] = true

      // propagated to next ports
      for _, nextp := range c.getNextPorts(p) {
        if shortloop[nextp] {
          queue = append(append([]*Port{}, nextp), queue...) //depth-first
        } else {
          queue = append(queue, nextp) //breadth-first
        }
      }
    }
  } //queue loop
}

func (c *Circuit) hasShortLoop(p *Port) bool {
  return c.hasShortLoopRec(p, p, 4)
}

func (c *Circuit) hasShortLoopRec(target, other *Port, depth int) bool {
  if depth == 0 {
    return false
  }

  for _, n := range c.getNextPorts(other) {
    if n == target {
      return true
    }
    if c.hasShortLoopRec(target, n, depth-1) {
      return true
    }
  }

  return false
}

