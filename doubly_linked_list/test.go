package main

import "fmt"

func main() {
	l := new(list)
	s := new(list)

	s.Append("asd")
	fmt.Println(s.first.data, " ... ", s.last.prev.data, ":", s.last.data, ":", s.last.next.data)
	s.Append("dsa")
	fmt.Println(s.first.data, " ... ", s.last.prev.data, ":", s.last.data, ":", s.last.next.data)
	s.Append("ee")
	fmt.Println(s.first.data, " ... ", s.last.prev.data, ":", s.last.data, ":", s.last.next.data)
	s.Append("gf")
	fmt.Println(s.first.data, " ... ", s.last.prev.data, ":", s.last.data, ":", s.last.next.data)
	s.Append("d")
	fmt.Println(s.first.data, " ... ", s.last.prev.data, ":", s.last.data, ":", s.last.next.data)

	fmt.Println("Len:", s.Len())

	for i := 0; i < 10; i++ {
		l.Append(i)
		fmt.Println(l.first.data, " ... ", l.last.prev.data, ":", l.last.data, ":", l.last.next.data)
	}

	fmt.Println("Len:", l.Len())
}
