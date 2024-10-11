package main

import "sort"

func main() {
	ages := map[string]int{"mgmg": 20, "koko": 18}
	//ages := make(map[string]int)
	//ages["mgmg"] = 20
	//ages["koko"] = 18
	names := make([]string, 0, len(ages))
	for name := range ages {
		names = append(names, name)
	}
	sort.Strings(names)
	var fruits map[string]int
	fruits["apple"] = 20 //cannot assign to nil type map //runtime error

}
