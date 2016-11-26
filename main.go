package main

import (
	"fmt"
	"github.com/docopt/docopt-go"

	"github.com/zeusproject/zeus-server/account"
	"github.com/zeusproject/zeus-server/char"
	"github.com/zeusproject/zeus-server/zone"
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

		if mode == "account" {
			account.Run(args)
		} else if mode == "char" {
			char.Run(args)
		} else if mode == "zone" {
			zone.Run(args)
		} else {
			fmt.Printf("Invalid server mode: %s\n", mode)
			return
		}
	}
}
