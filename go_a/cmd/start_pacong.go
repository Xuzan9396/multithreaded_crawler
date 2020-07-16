package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
	"os/exec"
	"pacong/zhenai/engine"
	"pacong/zhenai/parser"
)


func init() {
	var daemon bool


	var startPacong = &cobra.Command {
		Use: "start",
		Short: "开始爬虫。。。",

		Run: func(cmd *cobra.Command, args []string) {
			if daemon {
				command := exec.Command("./test", "start")
				command.Start()

				fmt.Printf("test start, [PID] %d running...\n", command.Process.Pid)
				ioutil.WriteFile("test.lock", []byte(fmt.Sprintf("%d", command.Process.Pid)), 0666)
				daemon = false
				os.Exit(0)
			} else {
				fmt.Println("test start")
			}

			data := []engine.Request{
				{
					Url:"https://cn.investing.com/indices/shanghai-se-a-share",
					ParserFunc: func(bytes []byte) engine.ParseResult {
						return parser.Aparser(bytes,"https://cn.investing.com/indices/shanghai-se-a-share")
					},
				},
				{
					Url:"https://s.weibo.com/top/summary",
					ParserFunc: func(bytes []byte) engine.ParseResult {
						return parser.WeiboList(bytes,"https://s.weibo.com/top/summary")
					},
				},
			}


			engine.Run(data...)

		},
	}


	startPacong.Flags().BoolVarP(&daemon, "deamon", "d", false, "is daemon?")
	rootCmd.AddCommand(startPacong)
}