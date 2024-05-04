package profile

import (
	"github.com/FascodeNet/alterlinux/alteriso5/config"
	"github.com/FascodeNet/alterlinux/alteriso5/utils"
	"github.com/Hayao0819/nahi/cobrautils"
	"github.com/spf13/cobra"
)

var subCmds = cobrautils.Registory{}
var profile *config.Profile

func Cmd() *cobra.Command {
	cmd := cobra.Command{
		Use:   "profile",
		Short: "Debug your profile",
		PersistentPreRunE: utils.WithParentPersistentPreRunE(func(cmd *cobra.Command, args []string) error {
			p, err := config.ReadProfile(args[0])
			if err != nil {
				return err
			}

			profile = p

			return nil
		}),
	}

	subCmds.BindSubCmds(&cmd)

	return &cmd
}
