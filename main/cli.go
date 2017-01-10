package main

import (
	"os"
	"fmt"
	"github.com/17golang/golang/cmd/utils"
	"gopkg.in/urfave/cli.v1"
	"github.com/17golang/golang/goutils/config"
)

var (
	app = utils.NewApp("指旺命令行接口")
	AppConstant map[string]interface{} = make(map[string]interface{})
)

func main() {
	app.Action = action
	app.AddAfter(func(*cli.Context) error {
		fmt.Println("hhhhhhh")
		return nil
	})

	app.AddFlag(cli.StringFlag{
		Name:"config",
		Usage:"设置配置文件",
	})
	if err := app.Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func action(ctx *cli.Context) error {
	config.ReadConfig(ctx, "config", &AppConstant)
	fmt.Println(AppConstant)
	return nil
}
