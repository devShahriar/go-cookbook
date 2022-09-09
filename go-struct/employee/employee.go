package employee

import "fmt"

type Employee struct {
	EmployeeId string
	Name       string
	Age        int
	Gender     string
}

func (e *Employee) PrintUserDetails() {
	fmt.Printf("EmployeeId : %v\n Name: %v\n Age: %v\n Gender: %v", e.EmployeeId, e.Name, e.Age, e.Gender)
}

func (e *Employee) UpdateUserId(id string) {
	e.EmployeeId = id
}

func (e *Employee) UpdateAge(age int) {
	e.Age = age
}

func (e *Employee) UpdetAgeAndUserId(age int, id string) {
	e.UpdateAge(age)
	e.UpdateUserId(id)
}

func (e *Employee) UpdateUserName(name string) {
	e.Name = name
}
