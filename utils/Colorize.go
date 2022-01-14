package utils

import "fmt"

func Colorize(color string, message string) {
	fmt.Println(string(color), message, string(Reset()))
}
