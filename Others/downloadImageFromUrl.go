package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	filename := "sample.pdf"
	url := "http://www.africau.edu/images/default/sample.pdf"
	req, _ := http.Get(url)
	defer req.Body.Close()
	file, err := os.Create(filename)
	_, err = io.Copy(file, req.Body)
	if err != nil {
		fmt.Print(err)
	}

}
