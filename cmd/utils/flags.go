package utils

import (
	"gopkg.in/urfave/cli.v1"
	"path/filepath"
	"os"
	"flag"
	"fmt"
	"runtime"
	"github.com/golang/glog"
)

type App struct {
	*cli.App
}

func (app *App)AddFlag(flag cli.Flag) {
	app.Flags = append(app.Flags, flag)
}

func (app *App)AddFlags(flag []cli.Flag) {
	app.Flags = append(app.Flags, flag...)
}

func (app *App)AddCommand(cmd cli.Command) {
	app.Commands = append(app.Commands, cmd)
}

func (app *App)AddCommands(cmds []cli.Command) {
	app.Commands = append(app.Commands, cmds...)
}

func (app *App)AddBefore(before cli.BeforeFunc) {
	b := app.Before
	if b != nil {
		app.Before = func(ctx *cli.Context) error {
			if err := b(ctx); err != nil {
				return err
			}
			if err := before(ctx); err != nil {
				return err
			}
			return nil
		}
	} else {
		app.Before = before
	}
}

func (app *App)AddAfter(after cli.AfterFunc) {

	a := app.After
	if a != nil {
		app.After = func(ctx *cli.Context) error {
			if err := a(ctx); err != nil {
				return err
			}
			if err := after(ctx); err != nil {
				return err
			}
			return nil
		}
	} else {
		app.After = after
	}
}

func NewApp(usage string) *App {
	app:=&App{
		App:cli.NewApp(),
	}
	//app := cli.NewApp()
	app.Name = filepath.Base(os.Args[0])
	app.Author = "指旺金科"
	app.Email = "shouhe.wu@fintecher.cn"
	app.Version = Version
	app.Usage = usage
	app.Commands = []cli.Command{
		glogCommand,
	}
	app.Before = func(ctx *cli.Context) error {
		runtime.GOMAXPROCS(runtime.NumCPU())
		return nil
	}
	app.After = func(ctx *cli.Context) error {
		glog.Flush()
		return nil
	}
	return app
}

//glog
var (
	glogCommand = cli.Command{
		Action:    glogAction,
		Name:      "glog",
		Usage:     "设置glog",
		Flags:[]cli.Flag{
			cli.IntFlag{
				Name: "v", Value: 0, Usage: "log level for V logs",
			},
			cli.BoolFlag{
				Name: "logtostderr", Usage: "log to standard error instead of files",
			},
			cli.IntFlag{
				Name:  "stderrthreshold",
				Usage: "logs at or above this threshold go to stderr",
			},
			cli.BoolFlag{
				Name: "alsologtostderr", Usage: "log to standard error as well as files",
			},
			cli.StringFlag{
				Name:  "vmodule",
				Usage: "comma-separated list of pattern=N settings for file-filtered logging",
			},
			cli.StringFlag{
				Name: "log_dir", Usage: "If non-empty, write log files in this directory",
			},
			cli.StringFlag{
				Name:  "log_backtrace_at",
				Usage: "when logging hits line file:N, emit a stack trace",
				Value: ":0",
			},
		},
	}
)

func glogAction(ctx *cli.Context) error {
	GlogGangstaShim(ctx)
	return nil
}

func glogFlagShim(fakeVals map[string]string) {
	flag.VisitAll(func(fl *flag.Flag) {
		if val, ok := fakeVals[fl.Name]; ok {
			fl.Value.Set(val)
		}
	})
}

func GlogGangstaShim(c *cli.Context) {
	_ = flag.CommandLine.Parse([]string{})
	glogFlagShim(map[string]string{
		"v":                fmt.Sprint(c.Int("v")),
		"logtostderr":      fmt.Sprint(c.Bool("logtostderr")),
		"stderrthreshold":  fmt.Sprint(c.Int("stderrthreshold")),
		"alsologtostderr":  fmt.Sprint(c.Bool("alsologtostderr")),
		"vmodule":          c.String("vmodule"),
		"log_dir":          c.String("log_dir"),
		"log_backtrace_at": c.String("log_backtrace_at"),
	})
}
