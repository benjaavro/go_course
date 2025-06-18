package main

import "fmt"

// Person Struct definition
type Person struct {
	firstName string
	lastName  string
	age       int
	address   Address
	Phone     // Anonymous field to surface fields
}

type Phone struct {
	home   string
	mobile string
}

// Define methods for struct
func (p *Person) fullName() string {
	return p.firstName + " " + p.lastName
}

func (p *Person) ageIncrease() {
	p.age++
}

type Address struct {
	city, country string
}

func main() {

	// Struct init
	p := Person{
		firstName: "James",
		lastName:  "Bond",
		age:       32,
		address:   Address{
			// city:    "New York",
			// country: "US",
		},
		Phone: Phone{
			home:   "01800-HOME",
			mobile: "01800-MOBILE",
		},
	}

	p2 := Person{
		firstName: "James",
		lastName:  "Bond",
		age:       32,
		address:   Address{
			// city:    "London",
			// country: "UK",
		},
		Phone: Phone{
			home:   "01800-HOME",
			mobile: "01800-MOBILE",
		},
	}

	// Struct access
	fmt.Println(p.firstName, p.lastName, p.age)
	fmt.Println(p2.firstName, p2.lastName, p2.age)

	// Anonymous Struct
	a := struct {
		firstName string
		lastName  string
	}{
		firstName: "Benjamin",
		lastName:  "Avila Rosas",
	}

	fmt.Println(a.firstName, a.lastName)

	// Struct comparison
	fmt.Println("Are they equal? ", p == p2)

	// Struct method access
	fmt.Println(p.fullName())
	fmt.Println(p2.firstName, p2.lastName, p2.age)
	p2.ageIncrease()
	fmt.Println(p2.firstName, p2.lastName, p2.age)
	fmt.Println(p2.address.city, p2.address.country)
	fmt.Println(p2.home, p2.mobile)

}
