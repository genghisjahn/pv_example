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

func main() {
	p1 := person{"Jon", 41}
	a1 := animal{"Tumnus", 8}
	personstuffV(p1)
	personstuffP(&p1)

	doGreet(&a1)
	doGreet(&p1)
}

func (a *animal) Greet() string {
	return fmt.Sprintf("ArfArfArfBowWow!!")
}

func (p *person) Greet() string {
	return fmt.Sprintf("I'm %v and I'm %v years old.%T\n", p.name, p.age, p)

}

func (p *person) Greet1() {
	fmt.Printf("I'm %v and I'm %v years old.%T\n", p.name, p.age, p)
}

func (p person) Greet2() {
	fmt.Printf("I'm %v and I'm still %v, my birthday is a ways off.%T\n", p.name, p.age, p)
}

func personstuffV(p person) {
	p.Greet1()
	p.Greet2()
}

func personstuffP(p *person) {
	p.Greet1()
	p.Greet2()
}

func doGreet(g greeter) {
	fmt.Println(g.Greet(), fmt.Sprintf("%T", g))

}
