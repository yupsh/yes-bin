package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"

	yup "github.com/gloo-foo/framework"
	. "github.com/yupsh/yes"
)

const (
	flagCount = "count"
)

func main() {
	app := &cli.App{
		Name:  "yes",
		Usage: "output a string repeatedly until killed",
		UsageText: `yes [STRING]...

   Repeatedly output a line with all specified STRING(s), or 'y'.`,
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:    flagCount,
				Aliases: []string{"n"},
				Usage:   "output COUNT lines instead of repeating forever",
			},
		},
		Action: action,
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "yes: %v\n", err)
		os.Exit(1)
	}
}

func action(c *cli.Context) error {
	var params []any

	// Add all arguments
	for i := 0; i < c.NArg(); i++ {
		params = append(params, c.Args().Get(i))
	}

	// Add flags based on CLI options
	if c.IsSet(flagCount) {
		params = append(params, Count(c.Int(flagCount)))
	}

	// Create and execute the yes command
	cmd := Yes(params...)
	return yup.Run(cmd)
}
