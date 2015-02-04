package main

import "fmt"

type trie struct {
  root   trie_node
  length int
}

type trie_node struct {
  present bool
  data    interface{}
  keys    map[string]*trie_node
}

func (this *trie) Defined(key string) bool {
  res := true
  node := &this.root

  for res {
    node, res = node.keys[key[0:1]]
    key = key[1:]
  }

  return res && node.present
}

func (this *trie) Define(key string, x interface{}) {
  n1 := &this.root

  for len(key) != 0 {
    letter := key[0:1]

    if len(n1.keys) == 0 {
      n1.keys = make(map[string]*trie_node)
    }

    n2, defined := n1.keys[letter]

    if !defined {
      n2 = new(trie_node)
      n1.keys[letter] = n2
    }

    if len(key) == 1 {
      n2.present = true
      n2.data = x
    }

    key = key[1:]
    n1 = n2
  }

  this.length++
}

func (this *trie) Get(key string) interface{} {
  node := &this.root

  for len(key) > 0 {
    node = node.keys[key[0:1]]
    key = key[1:]
  }

  return node.data
}

func (this *trie) Len() int {
  return this.length
}

func (this *trie) Delete(key string) {
  this.root.Delete(key)
  this.length--
}

func (this *trie_node) Delete(key string) {
  if len(key) > 0 {
    letter := key[0:1]
    next_node := this.keys[letter]

    next_node.Delete(key[1:])

    if len(next_node.keys) == 0 && !next_node.present {
      delete(this.keys, letter)
    }
  } else {
    this.present = false
  }
}

func (this *trie) String() {
  node := &this.root
  node.String(0)
  print("\n")
}

func (this *trie_node) String(depth int) {
  var indent string

  for i := 0; i < depth; i++ {
    indent += " "
  }

  for k, v := range this.keys {
    fmt.Print(indent, k)
    if v.present {
      fmt.Print(" : ", v.data)
    }
    fmt.Print("\n")
    v.String(depth + 1)
  }
}
