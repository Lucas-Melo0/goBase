package main

import "fmt"

func main() {
	var employees = map[string]int{"Benjamin": 20, "Manuel": 26, "Brenda": 19, "Dario": 44, "Pedro": 30}
	employeeWithAgeGreaterThen20 := 0
	for _, age := range employees {
		if age > 20 {
			employeeWithAgeGreaterThen20++
		}
	}
	println(employeeWithAgeGreaterThen20)
	employees["Lucas"] = 25
	fmt.Println(employees)
	delete(employees, "Pedro")
	fmt.Println(employees)
}
