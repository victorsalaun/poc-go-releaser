package root

import (
	"github.com/MakeNowJust/heredoc"
	"github.com/spf13/cobra"
	versionCmd "github.com/victorsalaun/poc-go-releaser/pkg/cmd/version"
	"github.com/victorsalaun/poc-go-releaser/pkg/cmdutil"
)

func NewCmdRoot(factory *cmdutil.Factory) *cobra.Command {
	cmd := &cobra.Command{
		Use:           "qcv <command> <subcommand> [flags]",
		Short:         "qcv CLI",
		SilenceErrors: false,
		SilenceUsage:  true,
		Example: heredoc.Doc(`
			$ qcv install
			$ qcv verify
			$ qcv version
		`),
	}

	cmd.AddCommand(versionCmd.NewCmdVersion(factory))

	return cmd
}
