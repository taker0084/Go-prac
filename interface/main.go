package main

import "fmt"

//interfaceを作っておくとテストが簡単になる
type Animal interface{
	Says() string
	NumberOfLegs() int
}

type Dog struct{
	Name string
	Breed string
}

type Gorilla struct{
	Name string
	Color string
	NumberOfTeeth int
}

func main(){
	dog := Dog{
		Name: "Samson",
		Breed: "German Shepherd",
	}
	printInfo(&dog)

	gorilla := Gorilla{
		Name: "Jock",
		Color: "Grey",
		NumberOfTeeth: 38,
	}
	printInfo(&gorilla)
}

func printInfo(a Animal){
	fmt.Println("This animal says", a.Says(), "and has", a.NumberOfLegs(), "legs")
}

func (d *Dog) Says() string{
	return "Woof"
}
func (d *Dog) NumberOfLegs() int{
	return 4
}
func (g *Gorilla) Says() string{
	return "Ugh"
}
func (g *Gorilla) NumberOfLegs() int{
	return 2
}