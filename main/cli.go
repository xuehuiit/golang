package main

/**
  本文件演示如何使用go语言开发常驻内存程序时通用的开发框架，本框架演示的程序的入口代码，这些方式
  在生产系统中是可以使用的。

  本框架的命令行参数的处理采用了gopkg.in/urfave/cli.v1框架，在使用之前请自行下载

  如果有问题可以同我们联系 411321681

*/

import (
	"fmt"
	"github.com/17golang/golang/cmd/utils"
	"github.com/17golang/golang/goenv"
	"github.com/17golang/golang/goutils/config"
	gologging "github.com/golang/glog"
	"gopkg.in/urfave/cli.v1"
	"os"
	"os/signal"
)

var (
	// 初始时指明当前程序的名称
	app = utils.NewApp("指旺命令行接口")
	//定义程序的上下文环境，在这里用要给MAP来存储全局的变量，也可以引入其他的定义
	//AppConstant map[string]interface{} = make(map[string]interface{})
)

func main() {

	//定义程序入口方法，根据具体的项目自行定义
	app.Action = action

	//设置程序作者，不设置采用默认的方式
	app.Author = "zhiwang_test_monitor"

	//设置联系人邮件
	app.Email = "send_mail@test.com"

	app.AddBefore(func(*cli.Context) error {

		go signalListen()
		return nil

	})

	//入口方法直销完成之后的执行的方法
	app.AddAfter(func(*cli.Context) error {
		gologging.Info(" ===程序入口方法执行之后的方法===  ")
		return nil
	})

	//设置命令行参数
	app.AddFlag(cli.StringFlag{
		Name:  "config",
		Usage: "设置配置文件",
	})

	//设置当前程序的运行路径
	app.AddFlag(cli.StringFlag{
		Name:  "appdir",
		Usage: "设置当程序的运行目录",
	})

	//程序执行错误的时候执行的方法
	if err := app.Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

}

/*
 接受系统中断程序的信号量比如 在windows中的 ctrl+c linux下面的KILL -9
*/
func signalListen() {

	c := make(chan os.Signal)
	signal.Notify(c)

	for {
		s := <-c
		//收到信号后的处理，这里只是输出信号内容，可以做一些更有意思的事
		gologging.Flush()
		fmt.Println("get signal:", s)

		os.Exit(1)
	}

}

/**
 * 系统入口方法
 */
func action(ctx *cli.Context) error {

	//读取配置文件并存储到系统通用缓存中
	config.ReadConfig(ctx, "config", &goenv.AppConstant)
	utils.GlogShim(ctx)

	//保存当前
	appDir := ctx.GlobalString("appdir")
	goenv.AppConstant["APP_DIR"] = appDir

	gologging.Info(" ********** 程序入口方法 ****************** ")
	gologging.Info(goenv.AppConstant)

	for true {

	}

	return nil

}
