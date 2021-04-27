package version

import (
	"encoding/json"
	"fmt"
	"os"
	"runtime/debug"
	"strings"

	"github.com/olekukonko/tablewriter"
	"github.com/pubgo/lug/version"
	"github.com/pubgo/xerror"
	"github.com/spf13/cobra"
)

var trim = strings.TrimSpace
var Cmd = &cobra.Command{
	Use:   "version",
	Short: "Print the dependency package information",
	Example: trim(`
lug version
lug version table`),
	Run: func(cmd *cobra.Command, args []string) {
		defer xerror.RespExit()

		info, ok := debug.ReadBuildInfo()
		if !ok {
			return
		}

		var typ string

		if len(args) > 0 {
			typ = args[0]
		}

		if typ == "" {
			dt, err := json.MarshalIndent(info, "", "\t")
			xerror.Panic(err)
			fmt.Println(string(dt))
			return
		}

		if typ == "table" {
			table := tablewriter.NewWriter(os.Stdout)
			table.SetHeader([]string{"Path", "Version", "Replace"})
			table.Append([]string{info.Main.Path, version.Version, replace(info.Main.Replace)})

			for _, dep := range info.Deps {
				table.Append([]string{dep.Path, dep.Version, replace(dep.Replace)})
			}
			table.Render()
		}
	},
}

func shortSum(sum string) string {
	sum = strings.Trim(sum, "h1:")
	if len(sum) > 10 {
		return sum[:10]
	}
	return sum
}

func replace(dep *debug.Module) string {
	if dep == nil {
		return ""
	}

	return fmt.Sprintf("%s:%s", dep.Path, dep.Version)
}
