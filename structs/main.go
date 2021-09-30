package main

import "fmt"

type ContactInfo struct {
	email   string
	zipCode int
}

type Person struct {
	firstName string
	lastName  string
	ContactInfo
}

func main() {

	jim := Person{
		firstName: "Jim",
		lastName:  "Party",
		ContactInfo: ContactInfo{
			email:   "jim@gmail.com",
			zipCode: 94000,
		},
	}
	// Turn `address` into `value` with *address
	// Turn `value` into `address` with &value
	// & operator is memory address of the value
	// the variable is pointing to
	// it's pass-by-value
	jim.updateName("jimmy")
	jim.print()
}

// needs pointer to update already created struct...
// type description = we're working with a pointer, can only receive a pointer
func (p *Person) updateName(newFirstName string) {
	// * operator is value at that memory address for pointer
	// this would be the Jim struct in memory
	(*p).firstName = newFirstName
}

func (p Person) print() {
	fmt.Printf("%+v", p)
}
