package main

import (
	"fmt"
	"jsmole/web"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func Process(url string, output string) {
	wq, err := web.CreateNewWebQuery(url)
	if err != nil {
		fmt.Println("Something wrong happened", err)
	}

	if urls, err := wq.GetMaps(); err == nil {
		// Iterating through js map files
		for _, elt := range urls {
			err = web.ProcessMap(elt, output)
			if err != nil {
				fmt.Println("err", err)
			}
		}
	}
}

func main() {
	var url string
	var output string

	app := &cli.App{
		Name:  "jsmole",
		Usage: "Google debugger but locally",

		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "url",
				Aliases:     []string{"u"},
				Value:       "",
				Usage:       "Website's url to be scanned",
				Destination: &url,
			},
			&cli.StringFlag{
				Name:        "output",
				Value:       "./output",
				Aliases:     []string{"o"},
				Usage:       "Output folder to be selected",
				Destination: &output,
			},
		},

		Action: func(cCtx *cli.Context) error {
			if url == "" {
				return cli.Exit("No url found", 1)
			}

			Process(url, output)
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
