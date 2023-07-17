package commands

import (
	"github.com/boreq/guinea"
	"github.com/boreq/velo/cmd/velo/commands/users"
)

var MainCmd = guinea.Command{
	Run: runMain,
	Subcommands: map[string]*guinea.Command{
		"run":            &runCmd,
		"users":          &users.UsersCmd,
		"default_config": &defaultConfigCmd,
	},
	ShortDescription: "an activity tracking service",
	Description: `
Velo is a self-hosted activity tracking service.
`,
}

func runMain(c guinea.Context) error {
	return guinea.ErrInvalidParms
}
