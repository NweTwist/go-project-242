package main

import (
	"fmt"

	"github.com/urfave/cli/v3"

	code "code"
	"context"
	"os"
)

func main() {
	cmd := &cli.Command{
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "human",
				Aliases: []string{"H"},
				Usage:   "human-readable sizes (auto-select unit)",
			},

			&cli.BoolFlag{
				Name:    "all",
				Aliases: []string{"a"},
				Usage:   "include hidden files and directories",
			},

			&cli.BoolFlag{
				Name:    "recursive",
				Aliases: []string{"r"},
				Usage:   "recursive size of directories",
			},
		},
		Name:      "hexlet-path-size",
		Usage:     "print size of a file or directory; supports -r (recursive), -H (human-readable), -a (include hidden)",
		ArgsUsage: "<path>",
		Action: func(ctx context.Context, c *cli.Command) error {
			if c.NArg() != 1 {
				return fmt.Errorf("expected exactly one path")
			}

			humanReadable := c.Bool("human")
			includeHidden := c.Bool("all")
			recursive := c.Bool("recursive")

			path := c.Args().Get(0)
			size, err := code.GetPathSize(path, recursive, humanReadable, includeHidden)
			if err != nil {
				return err
			}

			fmt.Printf("%s\t%s\n", size, path)
			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
