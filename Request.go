package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func Request() {
	resp, err := http.Get("https://juiceshop-test-1.herokuapp.com")
	if err != nil {
		log.Fatalln(err)
	}
	//We Read the response body on the line below.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	//Convert the body to type string
	sb := string(body)
	log.Printf(sb)
}
