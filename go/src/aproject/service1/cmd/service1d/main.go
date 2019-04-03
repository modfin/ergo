package main

import (
	"aproject/service0"
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {


	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {



		j, err := json.Marshal(service0.ModelService0{A: 3, B: 4})

		if err != nil {
			writer.WriteHeader(500)
			return
		}
		_, _ = writer.Write(j)
	})

	fmt.Println("Starting server")
	_ = http.ListenAndServe(":8080", nil)

}
