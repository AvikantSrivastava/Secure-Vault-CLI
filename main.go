package main

import (
	"flag"
	"fmt"
	"os"

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
		Items: []string{"login", "signup"},
	}

	if len(os.Args) < 2 {
		_, result, err = prompt.Run()
		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}

	} else {
		result = os.Args[1]
	}

	switch result {
	case "login":
		HandleLogin(loginCmd)
	case "signup":
		HandleLogin(signupCmd)

	default: // if we don't understand the input
		colorize(ColorRed, "Enter a valid command")
	}

}

func HandleLogin(loginCmd *flag.FlagSet) {
	fmt.Printf("login")
	// loginCmd.Parse(os.Args[2:])

}

func HandleSignup(signupCmd *flag.FlagSet) {
	fmt.Printf("signup")
	// signupCmd.Parse(os.Args[2:])

}
