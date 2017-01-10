package config

import (
	"io/ioutil"
	"gopkg.in/urfave/cli.v1"
	"errors"
	"gopkg.in/yaml.v2"
	"fmt"
)

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