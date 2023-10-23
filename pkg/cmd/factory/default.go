package factory

import (
	"github.com/victorsalaun/poc-go-releaser/pkg/cmdutil"
	"github.com/victorsalaun/poc-go-releaser/pkg/iostreams"
)

func New() *cmdutil.Factory {
	f := &cmdutil.Factory{}

	f.IOStreams = ioStreams()
	return f
}

func ioStreams() *iostreams.IOStreams {
	io := iostreams.System()
	return io
}
