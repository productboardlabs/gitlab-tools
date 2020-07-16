package main

import (
	"fmt"
	"os"

	"github.com/productboardlabs/gitlab-tools/cmd"
	"github.com/spf13/viper"
)

func main() {
	viper.AutomaticEnv()
	CMD := cmd.New(nil, nil)

	if err := CMD.RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
