package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world !")
}

type Config struct {
	Ip   string `json:"ip"`
	Port string `json:"port"`
}

func main() {
	fileBytes, err := ioutil.ReadFile("./config.json")
	if err != nil {
		fmt.Println(err)
	}

	var config Config

	if err := json.Unmarshal(fileBytes, &config); err != nil {
		fmt.Println(err)
	}

	http.HandleFunc("/", handler)

	fmt.Println("Serving from port :", config.Port)
	e := http.ListenAndServe(config.Ip+":"+config.Port, nil)
	if e != nil {
		fmt.Println("err", e)
	}
}
