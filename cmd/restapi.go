/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/tomwerneruk/aws-kata-tool/pkg/syncapi"
)

var listenPort int

// restapiCmd represents the restapi command
var restapiCmd = &cobra.Command{
	Use:   "restapi",
	Short: "A rest API stub with various endpoints that behave in certain ways",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("restapi called")
		syncapi.ServeMain(listenPort)
	},
}

func init() {
	rootCmd.AddCommand(restapiCmd)
	restapiCmd.Flags().IntVarP(&listenPort, "port", "p", 8080, "Port to listen on for mimic requests")
}
