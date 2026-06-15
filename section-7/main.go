package main

import "fmt"

type floatMap map[string]float64

func (m floatMap) output() {
	fmt.Println(m)
}

func main() {
	userNames := make([]string, 2, 5)

	userNames[0] = "Ada"

	// capacity of 5 makes append more efficient
	userNames = append(userNames, "Bob")

	fmt.Println(userNames)

	courses := make(floatMap, 2)

	courses["go"] = 4.7
	courses["react"] = 4.8

	courses.output()

	for _, user := range userNames {
		fmt.Println("Value:", user)
	}
}
