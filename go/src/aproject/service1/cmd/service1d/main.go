package main

import (
	"aproject/lib/fac"
	"aproject/service0"
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {


	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {

		i := fac.Factorial(3)
		ii := fac.Factorial(i)

		j, err := json.Marshal(service0.ModelService0{A: int(i), B: int(ii) })

		if err != nil {
			writer.WriteHeader(500)
			return
		}
		_, _ = writer.Write(j)
	})

	fmt.Println("Starting server")
	_ = http.ListenAndServe(":8080", nil)

}
