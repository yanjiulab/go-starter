package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

// IsAdult is a value-receiver method.
func (p Person) IsAdult() bool {
	return p.Age >= 18
}

type Employee struct {
	Person
	Title string
}

// Promote is a pointer-receiver method that mutates state.
func (e *Employee) Promote(newTitle string) {
	e.Title = newTitle
}

func main() {
	e := Employee{
		Person: Person{Name: "Bob", Age: 30},
		Title:  "Developer",
	}

	fmt.Printf("%+v\n", e)
	fmt.Println("embedded name:", e.Name)
	fmt.Println("adult?", e.IsAdult())

	e.Promote("Senior Developer")
	fmt.Println("after promote:", e.Title)
}
