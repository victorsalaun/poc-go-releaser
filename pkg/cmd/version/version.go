package version

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/victorsalaun/poc-go-releaser/internal/build"
	"github.com/victorsalaun/poc-go-releaser/pkg/cmdutil"
	"github.com/victorsalaun/poc-go-releaser/pkg/iostreams"
)

type VersionOptions struct {
	IO *iostreams.IOStreams
}

func NewCmdVersion(f *cmdutil.Factory) *cobra.Command {
	opts := &VersionOptions{
		IO: f.IOStreams,
	}
	cmd := &cobra.Command{
		Use: "version",
		RunE: func(cmd *cobra.Command, args []string) error {
			return run(opts)
		},
	}

	return cmd
}

func run(opts *VersionOptions) error {
	_, err := fmt.Fprintf(opts.IO.Out, "QCV %s (%s)", build.Version, build.Date)
	return err
}
