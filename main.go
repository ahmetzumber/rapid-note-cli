package main

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
)

func main() {
	a := strings.ToUpper("hello")
	fmt.Println(a)
	cmd := exec.Command("chrome")

	err := cmd.Run()

	if err != nil {
		log.Fatal(err)
	}
}
