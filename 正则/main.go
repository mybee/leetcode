package main

import (
	"fmt"
	"regexp"
)

func main() {
	s := `^/v2/roles/(\w+)/members$`
	reg := regexp.MustCompile(s)
	a := reg.FindStringSubmatch("/v2/roles/ss/members")

	fmt.Println(a)
}
