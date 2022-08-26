package running

import (
	"fmt"
	"os"
	"sort"

	"github.com/pubgo/funk/assert"
	"github.com/pubgo/funk/recovery"
	"github.com/urfave/cli/v2"

	"github.com/pubgo/lava/cmd/cmds/healthcmd"
	"github.com/pubgo/lava/cmd/cmds/versioncmd"
	"github.com/pubgo/lava/core/flags"
	"github.com/pubgo/lava/core/runmode"
	"github.com/pubgo/lava/version"
)

func Run(cmds ...*cli.Command) {
	defer recovery.Exit()

	var app = &cli.App{
		Name:                   version.Project(),
		Suggest:                true,
		UseShortOptionHandling: true,
		Usage:                  fmt.Sprintf("%s service", version.Project()),
		Version:                version.Version(),
		Flags:                  flags.GetFlags(),
		Commands:               append(cmds, versioncmd.Cmd(), healthcmd.Cmd()),
		ExtraInfo:              runmode.GetVersion,
	}

	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))
	assert.Must(app.Run(os.Args))
}
