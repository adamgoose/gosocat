package main

import (
	"fmt"

	"github.com/adamgoose/gosocat/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
