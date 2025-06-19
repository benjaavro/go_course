package main

import "fmt"

type person struct {
	name string
	age  int
}

func (p person) introduce() {
	fmt.Printf("Hi! I'm %s and I'm %d years old.\n", p.name, p.age)
}

type Employee struct {
	person
	employeeId string
	salary     float64
}

func (e Employee) introduce() {
	fmt.Printf("Hi! I'm %s and I earn $%.2f\n", e.name, e.salary)
}

func main() {
	emp := Employee{
		person: person{
			name: "Benjamin Avila",
			age:  29,
		},
		employeeId: "E0088",
		salary:     1198000.0,
	}

	fmt.Println(emp.name)
	fmt.Println(emp.age)
	fmt.Println(emp.employeeId)
	fmt.Println(emp.salary)
	emp.introduce()
}
