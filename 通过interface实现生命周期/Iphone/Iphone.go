package phone

import "fmt"

type Iphone struct {

}

type Phone interface {
	Call()
}

func (phone Iphone) Start() {
	fmt.Println("I am phone, I can call you!")
}

func (phone Iphone) Running() {
	fmt.Println("I am phone, I can call you!")
}

func (phone Iphone) Stop() {
	fmt.Println("I am phone, I can call you!")
}