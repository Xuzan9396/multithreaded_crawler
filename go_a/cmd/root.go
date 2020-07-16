package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"pacong/zhenai/engine"
	"pacong/zhenai/parser"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command {
	Use: "git",
	Short: "Git is a distributed version control system.",
	Long: `Git is a free and open source distributed version control system
designed to handle everything from small to very large projects 
with speed and efficiency.`,
	Run: func(cmd *cobra.Command, args []string) {
		var daemon bool
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

func Execute() {

	err := rootCmd.Execute()
	if err != nil{
		fmt.Println(err);
	}
}