package commands

import "github.com/urfave/cli/v2"

var Commands = []*cli.Command{
	Test(), Migrate(), Download(), Picture(),
}
