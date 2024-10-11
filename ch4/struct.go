package main

import "fmt"

type Employee struct {
	id            int
	Name, Address string
	Salary        float64
}

type Employee2 struct {
	id            int
	Address, Name string
	Salary        float64
}

type Circle struct {
	X int
	Y int
}

//type Wheel struct {
//	Circle Circle
//	Spokes int
//}

type Wheel struct {
	Circle //embedded, anonymous field
	Spokes int
}

func main() {
	//mgmg := Employee{1, "mgmg", "Yangon", 30000}
	koko := Employee{Address: "Yangon"}
	fmt.Println(koko)
	var wheel Wheel
	//wheel.Circle.X = 2
	//wheel.Circle.Y = 1
	wheel.X = 2
}
