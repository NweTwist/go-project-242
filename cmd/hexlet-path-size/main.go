package main

import (
	"fmt"

	"github.com/urfave/cli/v3"

	"context"
	goproject242 "go-project-242"
	"os"
)

var humanReadable bool
var includeHidden bool
var recursive bool

func main() {
	cmd := &cli.Command{
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:        "human",
				Aliases:     []string{"H"},
				Usage:       "Size in human-readable format",
				Destination: &humanReadable,
			},

			&cli.BoolFlag{
				Name:        "all",
				Aliases:     []string{"a"},
				Usage:       "Include hidden files and directories",
				Destination: &includeHidden,
			},

			&cli.BoolFlag{
				Name:        "recursive",
				Aliases:     []string{"r"},
				Usage:       "Recursive size of directories",
				Destination: &recursive,
			},
		},
		Name:  "hexlet-path-size",
		Usage: "Calculate the size of a directory",
		Action: func(ctx context.Context, c *cli.Command) error {
			if c.NArg() == 0 {
				fmt.Println("Error: No path provided")
				return nil
			}

			path := c.Args().Get(0)
			size := goproject242.GetPathSize(path, includeHidden, recursive)

			if humanReadable {
				fmt.Printf("%s, %s\n", goproject242.FormatSize(size), path)
			} else {
				fmt.Printf("%d, %s\n", size, path)
			}
			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
