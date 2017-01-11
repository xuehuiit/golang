package goenv

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

//系统全局变量上下文
var AppConstant map[string]interface{} = make(map[string]interface{})

//系统全局变量上下文
//var AppConstant map[string]string = make(map[string]string)

//判断当前系统是否为windows操作系统
func IfWinOs() bool {
	if runtime.GOOS == "windows" {

		return true
	} else {
		return false
	}
}

//系统启动时初始化简单版本
func InitAppSample() {

	flag.Parse()
	//1、接收启动参数

	appDir := flag.String("appdir", "", "系统启动目录")

	AppConstant["APP_DIR"] = *appDir //系统目录

	fmt.Println(" 当前项目路径    " + *appDir)

	//读取配置文件
	confdir := *appDir + "/cfg/config.json"
	_, err := os.Stat(confdir)

	if err == nil {
		//配置文件存在读取配置文件
		cfgbytes, readfileerr := ioutil.ReadFile(confdir)
		if readfileerr == nil {

			filecontent := string(cfgbytes)
			fmt.Println(filecontent)

			//读取配置文件，并将配置文件转化到map中
			mapconfig := make(map[string]interface{})
			err = json.Unmarshal(cfgbytes, &mapconfig)
			if err != nil {
				fmt.Println(err)
			}

			//fmt.Println("json to map ", mapconfig)

			for key, value := range mapconfig {
				configvalue := fmt.Sprintf("%s", value)
				AppConstant[key] = configvalue
				//fmt.Printf(" the %s    and value is %s  \n", key, value)
			}

			/*for key,value := range mapconfig{
							AppConstant[key]  = value
			 			}*/

			//fmt.Println("json to map ", mapconfig)
			//fmt.Println("The value of key1 is", mapconfig["1"])

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
