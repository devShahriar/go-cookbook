package handler

import (
	"encoding/json"
	"fmt"
	"github.com/devShahriar/go-cookbook/graphql-go/requestProcessor"
	"net/http"
)

func GraphqlHandler() func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {

		res,err := requestProcessor.Process(request)
		if err!= nil {
			fmt.Println(err)
		}
		fmt.Println(res)
		err = json.NewEncoder(writer).Encode(res)

		if err!=nil {
			fmt.Println(err)
		}

	}
}