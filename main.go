package main

import (
	"fmt"
	"github.com/docopt/docopt-go"

	"github.com/zeusproject/zeus-server/login"
)

func main() {
	usage := `Zeus Project

Usage:
	zeus-project server <mode>
	`

	args, err := docopt.Parse(usage, nil, true, "Zeus Project 1.0", false)

	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}

	if args["server"] == true {
		mode := args["<mode>"]

		if mode == "login" {
			login.Run(args)
		} else {
			fmt.Printf("Invalid server mode: %s\n", mode)
			return
		}
	}
}
