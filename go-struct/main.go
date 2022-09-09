package main

import "fmt"

var ListOfRegisterUser []Person = []Person{}

type Person interface {
	PrintUserDetails()
	UpdateUserId(string)
	UpdateAge(int)
	UpdetAgeAndUserId(int, string)
}

type Student struct {
	UserId string
	Age    int
}

//It varifies that &Student has implemented Person
var _ Person = &Student{}

// - struct itself (value type)
// - struct as reference (reference type)

func (s *Student) PrintUserDetails() {
	fmt.Printf("UserId : %v\nUser Age: %v", s.UserId, s.Age)
}

func (s *Student) UpdateUserId(userId string) {
	s.UserId = userId
}

func (s *Student) UpdateAge(age int) {
	s.Age = age
}

func (s *Student) UpdetAgeAndUserId(age int, userId string) {
	s.UpdateAge(age)
	s.UpdateUserId(userId)
}

func RegisterUser(p Person) {
	ListOfRegisterUser = append(ListOfRegisterUser, p)
}

func main() {
	s := &Student{"ramisa", 12}
	s1 := &Student{"chudip", 23}
	s2 := &Student{"ramisa_loves_shudip", 1111}
	RegisterUser(s)
	RegisterUser(s1)
	RegisterUser(s2)

	for _, i := range ListOfRegisterUser {
		fmt.Println(i)
	}

}
