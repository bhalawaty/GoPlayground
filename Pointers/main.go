package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	age       int
}

func main() {

	person := person{firstName: "bilal", lastName: "elhalawaty", age: 28}
	person.age = 29
	//salma := bilal.updateNameWithReturn("salma")
	//fmt.Println(salma)
	bilal := &person
	fmt.Println(&bilal)

	bilal.updateNameWithPassByValue("Hend")
	fmt.Println(bilal)

}
func (bilalObject *person) updateNameWithPassByValue(newName string) {
	(*bilalObject).firstName = newName
}

func (person person) updateNameWithReturn(newName string) person {
	person.firstName = newName
	return person
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
