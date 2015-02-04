package main

type list struct {
  first *list_node
  last  *list_node
}

type list_node struct {
  prev *list_node
  next *list_node
  data interface{}
}

func (self *list) Append(x interface{}) {
  new_node := new(list_node)
  new_node.data = x

  if (*self).first == nil {
    self.first = new_node
    self.last = new_node
  }

  new_node.prev = self.last
  new_node.next = self.first
  self.last = new_node

  new_node.prev.next = new_node
  new_node.next.prev = new_node
}

func (self *list) Len() (length int) {
  if self.first != nil {
    length++
  }

  for f := self.first; f != self.last; f = f.next {
    length++
  }

  return
}
