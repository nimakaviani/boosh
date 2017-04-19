package main

import (
	"fmt"
	"io"
	"os"

	"github.com/codegangsta/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "boosh"
	app.Usage = "BOOSH Outer Outer Shell"
	app.Version = "0.0.1.1"

	app.Commands = []cli.Command{
		{
			Name:      "generate",
			ShortName: "g",
			Usage:     "generate a CloudFormation template",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "manifest",
					Usage: "manifest to use to generate a template (default: stdin)",
				},
			},
			Action: func(c *cli.Context) {
				var source io.Reader

				if c.String("manifest") != "" {
					file, err := os.Open(c.String("manifest"))
					if err != nil {
						fatal(err)
					}

					source = file
				} else {
					source = os.Stdin
				}

				generate(source)
			},
		},
		{
			Name:      "deploy",
			ShortName: "d",
			Usage:     "deploy a CloudFormation template",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "name",
					Usage: "name of stack to deploy",
				},
				cli.StringFlag{
					Name:  "template",
					Usage: "template to deploy (default: stdin)",
				},
			},
			Action: func(c *cli.Context) {
				name := c.String("name")
				if name == "" {
					cli.ShowCommandHelp(c, "deploy")
					os.Exit(1)
				}

				var source io.Reader

				if c.String("template") != "" {
					file, err := os.Open(c.String("template"))
					if err != nil {
						fatal(err)
					}

					source = file
				} else {
					source = os.Stdin
				}

				deploy(name, source)
			},
		},
		{
			Name:      "watch",
			ShortName: "w",
			Usage:     "watch CloudFormation stack's events until completion",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "name",
					Usage: "name of stack to watch",
				},
			},
			Action: func(c *cli.Context) {
				name := c.String("name")
				if name == "" {
					cli.ShowCommandHelp(c, "watch")
					os.Exit(1)
				}

				watch(name)
			},
		},
		{
			Name:      "resources",
			ShortName: "r",
			Usage:     "create a stub from a stack's resources",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "name",
					Usage: "name of stack",
				},
			},
			Action: func(c *cli.Context) {
				name := c.String("name")
				if name == "" {
					cli.ShowCommandHelp(c, "resources")
					os.Exit(1)
				}

				resources(name)
			},
		},
	}

	app.Run(os.Args)
}

func fatal(err interface{}) {
	fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}
