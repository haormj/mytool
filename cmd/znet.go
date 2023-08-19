package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

var znetCmd = &cobra.Command{
	Use:   "znet",
	Short: "convert znet clash config to lastest config",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			log.Fatal("znet need a file path")
		}

		b, err := os.ReadFile(args[0])
		if err != nil {
			log.Fatalln(err)
		}

		m := make(map[string]any)
		if err := yaml.Unmarshal(b, &m); err != nil {
			log.Fatalln(err)
		}

		proxyes := m["proxies"].([]any)
		for _, proxy := range proxyes {
			p, ok := proxy.(map[string]any)
			if !ok {
				continue
			}
			if p["type"] != "vmess" {
				continue
			}
			if p["network"] != "ws" {
				continue
			}
			wsPath, ok := p["ws-path"].(string)
			if !ok {
				continue
			}
			delete(p, "ws-path")

			host, ok := p["ws-headers"].(map[string]any)["Host"].(string)
			if !ok {
				continue
			}
			delete(p, "ws-headers")

			p["ws-opts"] = map[string]any{
				"path": wsPath,
				"headers": map[string]string{
					"Host": host,
				},
			}
		}

		b, err = yaml.Marshal(m)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println(string(b))
	},
}

func init() {
	rootCmd.AddCommand(znetCmd)
}
