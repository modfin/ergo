package main

import (

	"aproject/service1/internal/config"
	"aproject/service1/internal/dao"
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {


	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {



		j, err := json.Marshal(arr)

		if err != nil {
			writer.WriteHeader(500)
			return
		}
		_, _ = writer.Write(j)
	})

	fmt.Println("Starting server")
	_ = http.ListenAndServe(":8080", nil)

}
