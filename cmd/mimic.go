/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/tomwerneruk/aws-kata-tool/pkg/mimicsrv"
)

var listenPort int

// mimicCmd represents the mimic command
var mimicCmd = &cobra.Command{
	Use:   "mimic",
	Short: "An HTTP Seriver that just mimics it's request.",
	Long: `An HTTP server that just responds with HTTP 200, and the request returned back into the body.
	
	The binding port can be changed via -p or --port. The default is 8080.
	
	Issue a GET request to /mimic from curl or your browser. The format of the response can be changed using /mimic?format=json or /mimic?format=pretty`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("mimic server starting")
		mimicsrv.ServeMimic(listenPort)
	},
}

func init() {
	rootCmd.AddCommand(mimicCmd)
	mimicCmd.Flags().IntVarP(&listenPort, "port", "p", 8080, "Port to listen on for mimic requests")
}
