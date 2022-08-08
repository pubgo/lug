package runmode

import (
	"io/ioutil"
	"os"
	"strings"
	"syscall"

	"github.com/pubgo/funk/assert"
	"github.com/pubgo/x/pathutil"

	"github.com/pubgo/lava/internal/pkg/utils"
	"github.com/pubgo/lava/version"
)

// 默认的全局配置
var (
	HttpPort = 8000
	GrpcPort = 50051
	Block    = true
	Project  = version.Project()
	Level    = "debug"

	IsDebug bool

	// DeviceID 主机设备ID
	DeviceID = version.DeviceID()

	// InstanceID service id
	InstanceID = version.InstanceID()

	Signal os.Signal = syscall.Signal(0)

	Version = version.Version()

	CommitID = version.CommitID()

	// Pwd 当前目录
	Pwd = assert.Exit1(os.Getwd())

	// Hostname 主机名
	Hostname = utils.FirstFnNotEmpty(
		func() string { return os.Getenv("HOSTNAME") },
		func() string { return assert.Exit1(os.Hostname()) },
	)

	// Namespace K8s命名空间
	Namespace = utils.FirstFnNotEmpty(
		func() string { return os.Getenv("NAMESPACE") },
		func() string { return os.Getenv("POD_NAMESPACE") },
		func() string {
			var file = "/var/run/secrets/kubernetes.io/serviceaccount/namespace"
			if pathutil.IsNotExist(file) {
				return ""
			}

			return strings.TrimSpace(string(assert.Exit1(ioutil.ReadFile(file))))
		},
	)
)
