package main

import (
	"fmt"
	"github.com/mitchellh/go-homedir"
	"github.com/ttacon/chalk"
	"github.com/urfave/cli"
	"log"
	"os"
)

func doesFileExist(path string) bool {
	_, err := os.Stat(path)

	if os.IsNotExist(err) {
		return false
	}
	return true
}

func main() {
	app := cli.NewApp()
	app.Name = "Boiled"
	app.Version = "0.1.0"
	app.Usage = "a tool to create and manage boilerplate applications"
	app.Authors = []cli.Author{
		cli.Author{
			Name:  "Tyler Boright",
			Email: "ru.lai.development@gmail.com",
		},
	}

	app.Action = func(c *cli.Context) error {
		fmt.Println(chalk.Magenta.Color("Add a boilerplate project today!"))
		return nil
	}

	app.Commands = []cli.Command{
		{
			Name:    "egg list",
			Aliases: []string{"el"},
			Usage:   "list all of the eggs you have in your boiled config",
			Action: func(c *cli.Context) error {
				homeDir, err := homedir.Dir()
				if err != nil {
					panic(err)
				}

				if doesFileExist(homeDir+"/.boiled/eggCarton.json") == false {
					return fmt.Errorf(chalk.Red.Color("You currently do not have any eggs.  Add a boilerplate and run again!"))
				}

				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
