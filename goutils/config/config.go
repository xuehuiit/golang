package config

/**
  我们采用yaml格式的文件来作为通用的项目的配置文件
*/

import (
	"errors"
	"fmt"
	"gopkg.in/urfave/cli.v1"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

/**
 *  读取配置文件
 *
 *   ctx  cli上下文环境
 *
 */
func ReadConfig(ctx *cli.Context, flag string, cfg interface{}) error {

	configFile := ctx.GlobalString(flag)
	if configFile == "" {
		return errors.New(fmt.Sprintf("%s flag not set", flag))
	}
	source, err := ioutil.ReadFile(configFile)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal([]byte(source), cfg)
	if err != nil {
		return err
	}
	return nil
}
