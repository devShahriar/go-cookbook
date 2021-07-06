package main

import (
	"fmt"
	"github.com/devShahriar/go-cookbook/graphql-go/handler"
	"net/http"
)

func main () {


	http.HandleFunc("/queries" , handler.GraphqlHandler())
	http.HandleFunc("/g", func(writer http.ResponseWriter, request *http.Request) {
			writer.Write([]byte("asdfasdf"))
	})

	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		panic(err.Error())
	}
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}

}