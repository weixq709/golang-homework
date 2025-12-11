package task2

import "fmt"

type Person struct {
	Name string
	Age  int
}

type Employee struct {
	Person
	EmployeeId int
}

func (e Employee) PrintInfo() {
	fmt.Printf("id : %d, Name: %s, Age: %d\n", e.EmployeeId, e.Name, e.Age)
}
