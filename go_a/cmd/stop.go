package cmd

import (
	"github.com/spf13/cobra"
	"io/ioutil"
	"os/exec"
)



func init() {
	var stopCmd = &cobra.Command{
		Use:   "stop",
		Short: "停止爬虫",
		Run: func(cmd *cobra.Command, args []string) {
			strb, _ := ioutil.ReadFile("test.lock")
			command := exec.Command("kill", string(strb))
			command.Start()
			println("gonne stop")
		},
	}
	rootCmd.AddCommand(stopCmd)
}

