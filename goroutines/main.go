package main

type Person struct {
	Name    string
	Phone   int
	Address string
}

func main() {
	http
	p := Person{"sdfasfas", 545454, "address"}
	PrintDetails(p)
	// data, _ := json.Marshal(p)

	// fmt.Println(string(data))

}
