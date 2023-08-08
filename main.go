/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"fmt"
	"os"

	"github.com/0xdod/uploadtocloud/cmd"
	"github.com/spf13/viper"
)

func init() {
	loadConfig()
}

func main() {
	cmd.Execute()
}


func loadConfig() {	
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")
	viper.SetConfigName(".config")
	viper.SetDefault("bucketname", "uploadtocloud")
	viper.AutomaticEnv() // read in environment variables that match
	
	if err := viper.ReadInConfig(); err != nil {
		fmt.Fprintln(os.Stderr, "Cannot read config file", viper.ConfigFileUsed())
		os.Exit(1)
	}
	// fmt.Println("Using config file:", viper.ConfigFileUsed()
	// fmt.Println(viper.AllSettings())
}
