package creational_petterns

import "fmt"

type Human interface {
	walk()
	clone() Human
}

type person struct {
	name string
}

func (f *person) walk() {
	fmt.Println(f.name + " is walking ...")
}

func (f *person) clone() Human {
	return &person{name: f.name + "_clone"}
}

// use this design pattern when you want to have work with our object but don't want to change prev one
