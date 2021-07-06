package requestProcessor

import (
	"encoding/json"
	"fmt"
	gq "github.com/graphql-go/graphql"
	"io/ioutil"
	"net/http"
)

func Process(request *http.Request) (interface{},error){

	fields := gq.Fields{
		"hello" : &gq.Field{
			Type: gq.String,
			Resolve: func(p gq.ResolveParams) (interface{}, error) {
				return "world",nil
			},
		},
	}
	reqBody, _ := ioutil.ReadAll(request.Body)

	body:= map[string]interface{}{}

	err :=json.Unmarshal(reqBody, &body)
 	if err != nil {
 		fmt.Print(err)
	}
    bodyText := body["query"].(string)

	fmt.Println("body sting:", bodyText)
	schema,err := gq.NewSchema(gq.SchemaConfig{Query: gq.NewObject(gq.ObjectConfig{Name:"Query",Fields: fields})})
	if err != nil {
		fmt.Println(err)
	}

	params := gq.Params {
		Schema: schema,
		RequestString: bodyText,
	}

	result := gq.Do(params)
	fmt.Println(result)
	return result , nil
}
