package main

import "fmt"

type Human struct {
	Name   string
	Gender string
	Age    int
}

func (h *Human) SayHi() {
	fmt.Printf("Hi, I am %s. My age %d\n", h.Name, h.Age)
}

func (h *Human) SetAge(age int) {
	h.Age = age
}
func (h *Human) SetName(name string) {
	h.Name = name
}

type Action struct {
	Human
}

func main() {
	action := Action{}

	action.SetAge(18)
	action.SetName("Bob")
	action.SayHi()

}
