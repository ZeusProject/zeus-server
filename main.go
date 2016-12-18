package main

import (
	"fmt"
	"github.com/docopt/docopt-go"
	"math/rand"
	"time"

	"github.com/zeusproject/zeus-server/account"
	"github.com/zeusproject/zeus-server/char"
	"github.com/zeusproject/zeus-server/inter"
	"github.com/zeusproject/zeus-server/utils/migrations"
	"github.com/zeusproject/zeus-server/zone"
)

func main() {
	rand.Seed(time.Now().Unix())

	usage := `Zeus Server

Usage:
	zeus-server db migrate [up | down]
	zeus-server server account [options]
	zeus-server server char [options]
	zeus-server server inter [options]
	zeus-server server zone [options]
	zeus-server -h | --help
	zeus-server --version

Options:
	-h --help   Show this screen.
	--version   Show version.
    --nobanner  Supress banner.
	`

	args, err := docopt.Parse(usage, nil, true, "Zeus Server 1.0", false)

	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}

	if args["server"] == true {
		if args["nobanner"] != true {
			fmt.Printf("\t                   \n")
			fmt.Printf("\t                   \n")
			fmt.Printf("\t _______ _   _ ___ \n")
			fmt.Printf("\t|_  / _ \\ | | / __|\n")
			fmt.Printf("\t / /  __/ |_| \\__ \\\n")
			fmt.Printf("\t/___\\___|\\__,_|___/\n")
			fmt.Printf("\t                   \n")
			fmt.Printf("\t                   \n")
		}

		if args["account"] == true {
			account.Run(args)
		} else if args["char"] == true {
			char.Run(args)
		} else if args["inter"] == true {
			inter.Run(args)
		} else if args["zone"] == true {
			zone.Run(args)
		}
	} else if args["db"] == true {
		if args["migrate"] == true {
			if args["up"] == true {
				migrations.RunUp(args)
			} else if args["down"] == true {
				migrations.RunDown(args)
			}
		}
	}
}
