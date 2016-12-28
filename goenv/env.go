package goenv

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

//系统全局变量上下文
var AppConstant map[string]string = make(map[string]string)

//系统启动时初始化
func InitApp() {

	//1、接收启动参数
	appDir := os.Args[0]
	AppConstant["APP_DIR"] = appDir //系统目录

	//读取配置文件
	confdir := appDir + "/cfg/config.json"
	_, err := os.Stat(confdir)

	if err == nil {
		//配置文件存在读取配置文件
		cfgbytes, readfileerr := ioutil.ReadFile(confdir)
		if readfileerr == nil {
			filecontent := string(cfgbytes)
			fmt.Println(filecontent)
		}
	} else {

		fmt.Println("没有发现配置文件 " + confdir)

	}

}

/**
获取当前的系统路径
*/
func GetCurrentPath() string {
	s, _ := exec.LookPath(os.Args[0])
	i := strings.LastIndex(s, "\\")
	path := string(s[0 : i+1])
	return path
}
