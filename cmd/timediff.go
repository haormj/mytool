package cmd

import (
	"fmt"
	"log"
	"time"

	"github.com/spf13/cobra"
)

var timediffCmd = &cobra.Command{
	Use:   "timediff",
	Short: "timediff",
	Run: func(cmd *cobra.Command, args []string) {
		format, err := cmd.Flags().GetString("format")
		if err != nil {
			log.Fatalln(err)
		}

		if len(args) != 2 {
			log.Fatalln("need two time string")
		}

		t1, err := time.Parse(format, args[0])
		if err != nil {
			log.Fatalln(err)
		}

		t2, err := time.Parse(format, args[1])
		if err != nil {
			log.Fatalln(err)
		}

		fmt.Println(t1.Sub(t2))
	},
}

func init() {
	rootCmd.AddCommand(timediffCmd)
	timediffCmd.Flags().StringP("format", "f", "2006-01-02 15:04:05", "time format")
}
