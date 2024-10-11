package main

import "fmt"

type car struct {
	Wheel int
}

type truck struct {
	Make string
	Year int
	car
}

type truck2 struct {
	Make string
	Year int
}
type Gear struct {
	One int
	Two int
}
type truck3 struct {
	Make string
	Year int
	Gear
}

func main() {
	t1 := truck{
		"Japan",
		2020,
		car{
			4,
		},
	}
	t2 := truck2{
		"USA",
		2021,
	}
	t3 := truck3{
		Make: "USA",
		Year: 2021,
		Gear: Gear{
			1,
			2,
		},
	}
	t4 := struct {
		A string
		B string
	}{
		"hello",
		"bro",
	}
	fmt.Println(t1, t2, t3, t4)
}
