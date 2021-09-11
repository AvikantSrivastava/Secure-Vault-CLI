package main

import (
	"flag"
	"fmt"
	"os"
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

func colorize(color Color, message string) {
	fmt.Println(string(color), message, string(ColorReset))
}

func main() {
	colorize(ColorYellow, "Welcome to Secret Vault")

	loginCmd := flag.NewFlagSet("login", flag.ExitOnError)
	signupCmd := flag.NewFlagSet("signup", flag.ExitOnError)

	if len(os.Args) < 2 {
		colorize(ColorRed, "Enter a valid command 'login' or 'signup")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "login":
		fmt.Printf("login")
		HandleLogin(loginCmd)
	case "signup":
		fmt.Printf("signup")
		HandleLogin(signupCmd)

	default: // if we don't understand the input
	}

}

func HandleLogin(loginCmd *flag.FlagSet) {
	loginCmd.Parse(os.Args[2:])

}

func HandleSignup(signupCmd *flag.FlagSet) {
	signupCmd.Parse(os.Args[2:])

}
