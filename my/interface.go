package main

import (
	"fmt"
	"math"
)

/* interface = collection of method signatures
a type "implements" an interface if it has all the methods defined in that interface
->go structs implicitly implements an interface if all requirements are met
*/

type shape interface {
	area() float64
	perimeter() float64
}

type rectangle struct {
	width, height float64
}

func (r rectangle) area() float64 {
	return r.width * r.height
}

func (r rectangle) perimeter() float64 {
	return 2*r.width + 2*r.height
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

//MULTIPLE INTERFACES

type dog interface {
	bark() string
}
type cat interface {
	meow() string
}

type supernaturalAnimal struct {
	Name string
}

func (s supernaturalAnimal) bark() string {
	msg := "Woof Woof"
	return fmt.Sprintf("%s", msg)
}
func (s supernaturalAnimal) meow() string {
	msg := "Meow Meow"
	return fmt.Sprintf("%s", msg)
}
