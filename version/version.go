package version

import (
	"runtime"

	ver "github.com/hashicorp/go-version"
	"github.com/pubgo/lug/runenv"
	"github.com/pubgo/lug/vars"
	"github.com/pubgo/xerror"
)

var CommitID = ""
var BuildTime = "2021-03-20 16:52:09"
var Version = "v0.0.1"
var Data = ""
var Domain string

func GetVer() map[string]interface{} {
	return map[string]interface{}{
		"data":       Data,
		"build_time": BuildTime,
		"version":    Version,
		"go_root":    runtime.GOROOT(),
		"go_arch":    runtime.GOARCH,
		"go_os":      runtime.GOOS,
		"go_version": runtime.Version(),
		"commit_id":  CommitID,
		"project":    runenv.Project,
	}
}

func init() {
	if Domain != "" {
		runenv.Domain = Domain
	}

	xerror.ExitErr(ver.NewVersion(Version))
	vars.Watch("version", func() interface{} { return GetVer() })
}
