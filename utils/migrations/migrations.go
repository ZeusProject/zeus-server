package migrations

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/mattes/migrate/file"
	"github.com/mattes/migrate/migrate"
	"github.com/mattes/migrate/migrate/direction"

	"github.com/zeusproject/zeus-server/database"
	"github.com/zeusproject/zeus-server/utils"

	_ "github.com/mattes/migrate/driver/postgres"
)

const MigrationsPath = "./migrations"

func buildUrl() (string, error) {
	config := &database.Config{}

	if err := utils.LoadConfig("database", config); err != nil {
		return "", err
	}

	return config.Postgres.ConnectionUrl(), nil
}

func RunUp(args map[string]interface{}) {
	url, err := buildUrl()

	if err != nil {
		fmt.Printf("An error ocurred reading the configuration: %v\n", err)
		os.Exit(-1)
		return
	}

	pipe := migrate.NewPipe()

	go migrate.Up(pipe, url, MigrationsPath)

	ok := writePipe(pipe)

	if !ok {
		os.Exit(-1)
	}
}

func RunDown(args map[string]interface{}) {
	url, err := buildUrl()

	if err != nil {
		fmt.Printf("An error ocurred reading the configuration: %v\n", err)
		os.Exit(-1)
		return
	}

	pipe := migrate.NewPipe()

	go migrate.Down(pipe, url, MigrationsPath)

	ok := writePipe(pipe)

	if !ok {
		os.Exit(-1)
	}
}

func writePipe(pipe chan interface{}) (ok bool) {
	okFlag := true

	if pipe != nil {
		for {
			select {
			case item, more := <-pipe:
				if !more {
					return okFlag
				}

				switch item.(type) {
				case string:
					fmt.Println(item.(string))

				case error:
					c := color.New(color.FgRed)
					c.Println(item.(error).Error(), "\n")
					okFlag = false

				case file.File:
					f := item.(file.File)
					c := color.New(color.FgBlue)

					if f.Direction == direction.Up {
						c.Print(">")
					} else if f.Direction == direction.Down {
						c.Print("<")
					}

					fmt.Printf(" %s\n", f.FileName)

				default:
					text := fmt.Sprint(item)
					fmt.Println(text)
				}
			}
		}
	}

	return okFlag
}
