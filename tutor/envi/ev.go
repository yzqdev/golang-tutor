package main

import (
	"fmt"
	"os"
	"os/user"
)

func getUserHome() string {
	current, err := user.Current()
	if err != nil {
		return "粗我"
	}
	return current.HomeDir
}
func getEnv() string {
	home := os.Getenv("appData")
	return home
}
func main() {
	fmt.Println(getUserHome())
	fmt.Println(getEnv())
}
