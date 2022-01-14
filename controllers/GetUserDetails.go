package controllers

import (
	"fmt"

	models "Secure-Vault-CLI/models"

	"github.com/manifoldco/promptui"
)

func GetUserDetails() (string, string) {

	var res models.ResponseModel

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
		fmt.Printf("Prompt failed %v\n", res.Err)
	}

	return username, password
}
