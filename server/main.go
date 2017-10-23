package main

import (
	"fmt"
	"github.com/mitchellh/mapstructure"	
)

type Message struct {
	Name string `json:name`
	Data interface{} `json:data`
}

type Channel struct {
	Id string `json:id`
	Name string `json:name`
}

func main() {
	router := &Router{}

	router.Handle("channel add", addChannel)

	http.Handle("/", router)
	http.ListenAndServe(":4000", nil)
}
