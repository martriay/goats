package main

import "fmt"

func main() {
	t := new(trie)
	keys := []string{"casa", "casanova", "casuela", "dsa", "monkey", "salomon"}

	for k, v := range keys {
		t.Define(v, k)
	}

	t.String()

  fmt.Println("---")

	t.Delete("casuela")

  t.String()
}
