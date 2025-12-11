package task2

import "testing"

func TestEmployeeInfo(t *testing.T) {
	emp := Employee{EmployeeId: 1, Person: Person{"joyce", 20}}
	emp.PrintInfo()
}
