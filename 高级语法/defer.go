package main

import (
	"fmt"
	"log"
)

func foo1(i *int) int {
	*i += 100
	defer func() { *i += 200 }()
	log.Printf("i=%d", *i)
	return *i
}

func foo2(i *int) (r int) {
	*i += 100
	defer func() { r += 200 }()
	log.Printf("i=%d", *i)
	return *i
}

//func main() {
//
//	var i, r int
//
//	i,r = 0,0
//	r = foo1(&i)
//	log.Printf("i=%d, r=%d\n", i, r)
//
//	i,r = 0,0
//	r = foo2(&i)
//	log.Printf("i=%d, r=%d\n", i, r)
//}

type Student struct {
	name string
}

func main() {
	m := map[string]Student{"people": {"li"},}
	m["people"].name = "liu"
	m["annimal"].name = "niu"

	for _, v := range  m {
		defer func() {
			fmt.Println(v.name)
		}()
	}
}


