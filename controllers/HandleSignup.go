package controllers

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
)

func HandleSignup(signupCmd *flag.FlagSet) {
	username, password := GetUserDetails()
	request_map := map[string]string{
		"username": username,
		"password": password,
	}

	request_json, err := json.Marshal(request_map)

	if err != nil {
		log.Fatal(err)
	}
	resp, err := http.Post("https://secret-vault.herokuapp.com"+"/signup", "application/json",
		bytes.NewBuffer(request_json))

	if err != nil {
		log.Fatal(err)
	}

	var res map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&res)
	fmt.Println(res["status"])

}
