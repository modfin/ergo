package main

import (
	"aproject/service0"
	"aproject/service0/internal/config"
	"aproject/service0/internal/dao"
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {


	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {


		c := config.Get()

		fmt.Println(" - Got request")

		model := dao.GetModel(c.ValA, c.ValB)
		arr := []service0.ModelService0{model}

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
