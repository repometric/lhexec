package main

import (
	"github.com/urfave/cli"
	"fmt"
	"os"
	"log"
	"bufio"
	"io"
	"github.com/repometric/lhexec/analyze"
)
const appVersion = "0.0.1"

func main() {
	app := cli.NewApp()
	app.Version = appVersion
	app.Usage = "Linterhub Engine Executor"
	app.Commands = []cli.Command{
		{
			Name:    "analyze",
			Aliases:  []string{"a"},
			Usage:   "Runs code analysis.",
			Action: func(c *cli.Context) error {
				var context = analyze.Context{
					Project:     c.String("project"),
					File:        c.String("file"),
					Folder:      c.String("folder"),
					Environment: c.String("environment"),
				}
				enginesArg := c.StringSlice("engine")
				if len(enginesArg) == 0 {
					cli.ShowCommandHelp(c, "analyze")
				} else {
					var out string
					if c.Bool("stdin") == true {
						reader := bufio.NewReader(os.Stdin)
						var output []rune

						for {
							input, _, err := reader.ReadRune()
							if err != nil && err == io.EOF {
								break
							}
							output = append(output, input)
						}

						for _, engine := range enginesArg {
							out, _ = analyze.AnalyzeStdin(engine, output)
							fmt.Println(engine)
							fmt.Printf("%s\n", out)
						}
					} else {
						for _, engine := range enginesArg {
							out, _ = analyze.Analyze(engine, context)
							fmt.Println(engine)
							fmt.Printf("%s\n", out)

						}
					}
				}
				return nil
			},
			Flags: []cli.Flag{
				cli.StringSliceFlag{
					Name:  "engine,e",
					Usage: "Engine name for analyze.",
				},
				cli.StringFlag{
					Name:  "project,p",
					Usage: "Project path for analyze.",
				},
				cli.StringFlag{
					Name:  "file,f",
					Usage: "File path for analyze.",
				},
				cli.StringFlag{
					Name:  "folder,F",
					Usage: "Folder path for analyze.",
				},
				cli.StringFlag{
					Name:  "environment,env",
					Usage: "The way how to analyze.",
				},
				cli.BoolFlag{
					Name:  "stdin",
					Usage: "Standard input.",
				},
			},

		},
		{
			Name:    "version",
			Aliases: []string{"v"},
			Usage:   "Returns current component version.",
			Action: func(c *cli.Context) error {
				fmt.Println(c.App.Version)
				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
