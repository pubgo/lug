package golug_app

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/pubgo/dix/dix_run"
	"github.com/pubgo/golug/golug_config"
	"github.com/pubgo/golug/golug_entry"
	"github.com/pubgo/golug/golug_plugin"
	"github.com/pubgo/golug/golug_version"
	"github.com/pubgo/xerror"
	"github.com/spf13/cobra"
)

func Start(ent golug_entry.Entry) (err error) {
	defer xerror.RespErr(&err)

	xerror.Panic(ent.Init())

	// 启动配置, 初始化组件, 初始化插件
	plugins := golug_plugin.List(golug_plugin.Module(ent.Options().Name))
	for _, pg := range append(golug_plugin.List(), plugins...) {
		key := pg.String()

		xerror.PanicF(err, "plugin [%s] load error", key)
		xerror.PanicF(pg.Init(golug_config.GetCfg().GetStringMap(key)), "plugin [%s] init error", key)
		ent.Use(pg.Handler())
	}

	xerror.Panic(dix_run.BeforeStart())
	xerror.Panic(ent.Start())
	xerror.Panic(dix_run.AfterStart())

	return
}

func Stop(ent golug_entry.Entry) (err error) {
	defer xerror.RespErr(&err)

	xerror.Panic(dix_run.BeforeStop())
	xerror.Panic(ent.Stop())
	xerror.Panic(dix_run.AfterStop())

	return nil
}

func Run(entries ...golug_entry.Entry) (err error) {
	defer xerror.RespErr(&err)

	if len(entries) == 0 {
		return xerror.New("[entries] should not be zero")
	}

	for _, ent := range entries {
		if ent == nil {
			return xerror.New("[ent] should not be nil")
		}

		opt := ent.Options()
		if opt.Name == "" || opt.Version == "" {
			return xerror.New("neither [name] nor [version] can be empty")
		}
	}

	var rootCmd = &cobra.Command{Use: golug_config.Domain, Version: golug_version.Version}
	rootCmd.PersistentFlags().AddFlagSet(golug_config.DefaultFlags())
	rootCmd.RunE = func(cmd *cobra.Command, args []string) error { return xerror.Wrap(cmd.Help()) }

	for _, ent := range entries {
		ent := ent
		cmd := ent.Options().Command

		// 检查Command是否注册
		for _, c := range rootCmd.Commands() {
			if c.Name() == cmd.Name() {
				return xerror.Fmt("command(%s) already exists", cmd.Name())
			}
		}

		// 注册plugin的command和flags
		entPlugins := golug_plugin.List(golug_plugin.Module(ent.Options().Name))
		for _, pl := range append(golug_plugin.List(), entPlugins...) {
			cmd.PersistentFlags().AddFlagSet(pl.Flags())
			xerror.Panic(ent.Commands(pl.Commands()))
		}

		runCmd := ent.Options().RunCommand
		runCmd.RunE = func(cmd *cobra.Command, args []string) (err error) {
			defer xerror.RespErr(&err)

			xerror.Panic(Start(ent))

			if golug_config.IsBlock {
				ch := make(chan os.Signal, 1)
				signal.Notify(ch, append(Shutdown(), syscall.SIGHUP)...)
				<-ch
			}

			xerror.Panic(Stop(ent))
			return nil
		}
		rootCmd.AddCommand(cmd)
	}
	return xerror.Wrap(rootCmd.Execute())
}

// ShutDownSignals returns all the signals that are being watched for to shut down services.
func Shutdown() []os.Signal {
	return []os.Signal{
		syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGKILL,
	}
}
