package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

func main() {

	// var map1 map[string]interface{}
	// map1 = make(map[string]interface{})

	// data := `{
	// 	"name":"John",
	// 	"age":29,
	// 	"hobbies":[
	// 	   "martial arts",
	// 	   "breakfast foods",
	// 	   "piano"
	// 	]
	//  }`

	// j.Unmarshal([]byte(data), &map1)

	// fmt.Printf("type of hobbies %T", map1["hobbies"])

	//struct

	data := `{
		"Name":"shudip",
		"Phone":454545,
		"Address":"adfasdafsd"
	}`
	type Person struct {
		Name    string `json:"name"`
		Phone   int    `json:"phone"`
		Address string `json:"address"`
	}

	var p1 Person

	err := json.Unmarshal([]byte(data), &p1)
	if err != nil {
		fmt.Println(err)
	}
	n := fmt.Sprintf("%T", p1)
	o := ConvStr(n)

	fmt.Println(o)
}
func ConvStr(str string) string {
	str = strings.Trim(str, "*")
	fmt.Println(str)
	ind := strings.Index(str, ".")
	fmt.Println(ind)
	if ind < 0 {
		return strings.ToLower(str)
	}

	str = str[ind+1:]
	return strings.ToLower(str)
}
