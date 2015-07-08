package main

import "fmt"

type person struct {
	name string
	age  int
}

type animal struct {
	name string
	age  int
}

type greeter interface {
	Greet() string
}

type server interface {
	TakeOrder([]string)
}

func main() {
	p1 := person{"Jon", 41}
	a1 := animal{"Tumnus", 8}
	personstuffV(p1)
	personstuffP(&p1)

	doGreet(&a1)
	doGreet(&p1)
	/*
		doGreet(p1)

		The above won't compile.  You'll get:
		cannot use p1 (type person) as type greeter in argument to doGreet: person does not implement greeter (Greet method has pointer receiver)

	*/

	doServe(&p1)
	//doServe(p1) doesn't work, explain why
}

func (a *animal) Greet() string {
	return fmt.Sprintf("ArfArfArfBowWow!!")
}

func (p *person) Greet() string {
	return fmt.Sprintf("I'm %v and I'm %v years old.", p.name, p.age)

}

func (p *person) TakeOrder(items []string) {
	fmt.Println("I already forgot the order.")
}

func personstuffV(p person) {
	fmt.Println(p.Greet())
}

func personstuffP(p *person) {
	fmt.Println(p.Greet())
}

func doGreet(g greeter) {
	fmt.Println(g.Greet())
}

func doServe(s server) {
	items := []string{"Fries", "Icecream"}
	s.TakeOrder(items)
}
