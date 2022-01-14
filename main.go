package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	controllers "Secure-Vault-CLI/controllers"
	models "Secure-Vault-CLI/models"
	Colors "Secure-Vault-CLI/utils"

	"github.com/manifoldco/promptui"
)

func main() {

	var res models.ResponseModel

	Colors.Colorize(Colors.Yellow(), "Welcome to Secret Vault\n")

	loginCmd := flag.NewFlagSet("login", flag.ExitOnError)
	signupCmd := flag.NewFlagSet("signup", flag.ExitOnError)

	prompt := promptui.Select{
		Label: "Select an Option",
		Items: []string{"Login", "Signup"},
	}

	if len(os.Args) < 2 {
		_, res.Result, res.Err = prompt.Run()
		if res.Err != nil {
			fmt.Printf("Prompt failed %v\n", res.Err)
			os.Exit(1)
		}

	} else {
		res.Result = os.Args[1]
	}

	switch strings.ToLower(res.Result) {
	case "login":
		controllers.HandleLogin(loginCmd)
	case "signup":
		controllers.HandleSignup(signupCmd)

	default: // if we don't understand the input
		Colors.Colorize(Colors.Red(), "Enter a valid command")
	}

}
