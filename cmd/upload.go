/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/0xdod/uploadtocloud/internal/s3"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// uploadCmd represents the upload command
var uploadCmd = &cobra.Command{
	Use:   "upload",
	Short: "Upload files to cloud storage",
	Long:  `upload file`,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("upload called E")
		var s3Config s3.Config
		viper.Unmarshal(&s3Config)
		uploader := s3.NewS3Uploader(s3Config)
		for _, arg := range args {
			file, err := os.Open(arg)
			if err != nil {
				return err
			}
			defer file.Close()
			var buf bytes.Buffer
			if _, err := buf.ReadFrom(file); err != nil {
				return err
			}
			url, err := uploader.Upload(context.Background(), buf.Bytes(), filepath.Base(arg))
			if err != nil {
				fmt.Println(err)
				return err
			}
			fmt.Println(url)
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(uploadCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// uploadCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// uploadCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
