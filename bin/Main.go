package main

import (
	"fmt"
	"os/exec"
)

func main() {
	cmd := exec.Command("notepad")
	err := cmd.Run()
	if err != nil {
		fmt.Println("出错了")
	}
	fmt.Print("aaaaaaaaa")
}
