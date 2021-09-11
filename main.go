package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/manifoldco/promptui"
)

type Color string

const (
	ColorBlack  Color = "\u001b[30m"
	ColorRed          = "\u001b[31m"
	ColorGreen        = "\u001b[32m"
	ColorYellow       = "\u001b[33m"
	ColorBlue         = "\u001b[34m"
	ColorReset        = "\u001b[0m"
)

var (
	err    error
	result string
)

func colorize(color Color, message string) {
	fmt.Println(string(color), message, string(ColorReset))
}

func main() {
	colorize(ColorYellow, "Welcome to Secret Vault\n")

	loginCmd := flag.NewFlagSet("login", flag.ExitOnError)
	signupCmd := flag.NewFlagSet("signup", flag.ExitOnError)

	prompt := promptui.Select{
		Label: "Select an Option",
		Items: []string{"Login", "Signup"},
	}

	if len(os.Args) < 2 {
		_, result, err = prompt.Run()
		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			os.Exit(1)
		}

	} else {
		result = os.Args[1]
	}
	switch strings.ToLower(result) {
	case "login":
		HandleLogin(loginCmd)
	case "signup":
		HandleSignup(signupCmd)

	default: // if we don't understand the input
		colorize(ColorRed, "Enter a valid command")
	}

}

func getUserDetails() (string, string) {
	username_prompt := promptui.Prompt{
		Label: "Username",
	}
	password_prompt := promptui.Prompt{
		Label: "Password",
		Mask:  '*',
	}
	username, err1 := username_prompt.Run()
	password, err2 := password_prompt.Run()

	if err1 != nil || err2 != nil {
		fmt.Printf("Prompt failed %v\n", err)
	}

	return username, password
}

func HandleLogin(loginCmd *flag.FlagSet) {
	username, password := getUserDetails()
	request_map := map[string]string{
		"username": username,
		"password": password,
	}

	request_json, err := json.Marshal(request_map)

	if err != nil {
		log.Fatal(err)
	}
	resp, err := http.Post("http://127.0.0.1:8000"+"/login", "application/json",
		bytes.NewBuffer(request_json))

	if err != nil {
		log.Fatal(err)
	}

	var res map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&res)
	fmt.Println(res["status"])

}

func HandleSignup(signupCmd *flag.FlagSet) {
	username, password := getUserDetails()
	request_map := map[string]string{
		"username": username,
		"password": password,
	}

	request_json, err := json.Marshal(request_map)

	if err != nil {
		log.Fatal(err)
	}
	resp, err := http.Post("http://127.0.0.1:8000"+"/signup", "application/json",
		bytes.NewBuffer(request_json))

	if err != nil {
		log.Fatal(err)
	}

	var res map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&res)
	fmt.Println(res["status"])

}
