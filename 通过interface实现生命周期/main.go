package main

import (
	"fmt"
	"leetcode/通过interface实现生命周期/Iphone"
)

type iPhone6 struct {
	iphone phone.Iphone
}

func (nokiaPhone iPhone6) Start() {
	nokiaPhone.iphone.Start()
	fmt.Println("I am iphone6, I am starting!")
}

func (nokiaPhone iPhone6) Running() {
	nokiaPhone.iphone.Running()
	fmt.Println("I am iphone6, I am running!")
}

func (nokiaPhone iPhone6) Stop() {
	nokiaPhone.iphone.Stop()
	fmt.Println("I am iphone6, I am stop!")
}


func main() {
	//var phone phone.Iphone
	phone := new(iPhone6)
	phone.Start()
	phone.Running()
	phone.Stop()
}
