package main

import (
	"fmt"

	"github.com/lemmego/collect"
)

type Person struct {
	Name string
	Age  int
}

func main() {

	// persons := []Person{
	// 	{
	// 		Name: "John",
	// 		Age:  30,
	// 	},
	// 	{
	// 		Name: "Jane",
	// 		Age:  25,
	// 	},
	// 	{
	// 		Name: "Bob",
	// 		Age:  40,
	// 	},
	// }

	// people := map[string]Person{
	// 	"John": {
	// 		Name: "John",
	// 		Age:  30,
	// 	},
	// 	"Jane": {
	// 		Name: "Jane",
	// 		Age:  25,
	// 	},
	// 	"Bob": {
	// 		Name: "Bob",
	// 		Age:  40,
	// 	},
	// }

	// below35Slice, _ := collect.NewSlice(persons).Filter(func(p Person, _ int) bool {
	// 	return p.Age < 35
	// }).Find(func(p Person) bool {
	// 	return p.Age == 30
	// })

	// below35Map, _ := collect.NewMap(people).Filter(func(p Person, _ string) bool {
	// 	return p.Age < 35
	// }).Keys().Find(func(s string) bool {
	// 	return s == "Bob"
	// })

	// fmt.Println(below35Slice)
	// fmt.Println(below35Map)

	// fm := collect.NewSlice([]int{2, 4, 6})

	// redueced := fm.Reduce(func(carry, value int, _ int) int {
	// 	return carry + value
	// }, 0)

	// fmt.Println(persons)
	// fmt.Println(redueced)
	// fmt.Println(fm.All())
	// fmt.Println(fm.Items())
	// fmt.Println(fm.Get())
	//
	m := map[int]int{
		1: 2,
		2: 4,
		3: 6,
	}

	mc := collect.NewMap(m)
	// Reduce function to sum the values
	result := mc.Reduce(func(carry, value int, key int) int {
		fmt.Println(key)
		return carry + value
	}, 0)

	fmt.Println(result) // Output: 6

}
