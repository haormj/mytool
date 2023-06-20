package main

import (
	"log"

	"github.com/haormj/mytool/cmd"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	cmd.Execute()
}
