package main

import (
	"fmt"
	"os"

	"github.com/gocarina/gocsv"
)

type CSVReader struct {
	Path string
}

func (c *CSVReader) ReadFile() {
	Clients := []*Client{}
	clientFile, err := os.OpenFile(c.Path, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		fmt.Println(err)
	}
	defer clientFile.Close()
	gocsv.UnmarshalFile(clientFile, &Clients)
	for _, client := range Clients {
		fmt.Println("Hello", client.Name)
	}

}

type Client struct {
	Id   string `csv:"client_id"`
	Name string `csv:client_name`
	Age  string `csv:client_age`
}

func main() {
	csvreader := CSVReader{"people.csv"}
	csvreader.ReadFile()
}
