package main

import "fmt"

type person struct {
	name string
	age  int
}

type greeter interface {
	Greet() string
}
type sleeper interface {
	Sleep() string
}

func main() {
	p1 := person{"Jon", 41}

	personV(p1)
	fmt.Println("after personV, p1 Age:", p1.age)
	personP(&p1)
	fmt.Println("after personP, p1 Age:", p1.age)

	doGreet(&p1)
	/*
		doGreet(p1)

		The above won't compile.  You'll get:
		cannot use p1 (type person) as type greeter in argument to doGreet: person does not implement greeter (Greet method has pointer receiver)

		If a method is defined on a pointer receiver, it can only be
		used in an interface if a pointer is passed into the interface.

		if pointer, you might be changing it.
		then you'd hvae to take the copy and pass a pointer to the copy
		those changes would be lost.

		it's not a techincal constraint, but it would be confusing
		pointer receivers expect to make changes and the caller be ablet to see them.
	*/

	goSleep(p1)
	goSleep(&p1)

	/*
		Sleep() is defined on a value receiver, so you can pass in
		values or pointers to functions
		that require the Sleeper interface
	*/

}

func (p *person) Greet() string {
	return fmt.Sprintf("%v says, \"I'm %v and I'm %v years old.\"", p.name, p.name, p.age)
}

func (p person) Sleep() string {
	return fmt.Sprintf("%v says, \"zzzzz\"", p.name)
}

func personV(p person) {
	//Greet() is defined on a pointer receiver,
	//but we can still call it from a value.
	fmt.Println("--Person Value--")
	fmt.Printf("Name %v\nAge:%v\n", p.name, p.age)
	//this works because if you pass a pointer, compiler adds an & (implicltly)
	//here you can see the call, but an interface you can't see it.
	fmt.Println(p.Greet())
	fmt.Printf("Type: %T\n", p)
	fmt.Printf("Changing age to :%v\n", 50)
	p.age = 50
}

func personP(p *person) {
	//Sleep() is defined on a value receiver,
	//but we can still call it from a pointer
	fmt.Println("--Person Pointer--")
	fmt.Printf("Name %v\nAge:%v\n", p.name, p.age)
	fmt.Println(p.Sleep())
	fmt.Printf("Type: %T\n", p)
	fmt.Printf("Changing age to :%v\n", 50)
	p.age = 50

}

func doGreet(g greeter) {
	fmt.Println("--doGreet--")
	fmt.Println("Greeting:", g.Greet())
	fmt.Printf("Type: %T\n", g)
	fmt.Printf("\n\n")
}

func goSleep(s sleeper) {
	fmt.Println("--gosleep--")
	fmt.Println("Sleeping:", s.Sleep())
	fmt.Printf("Type: %T\n", s)
	fmt.Printf("\n\n")
}
