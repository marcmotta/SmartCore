// cmd/smartcore/main.go
package main

import (
	"flag"
	"log"
	"os"

	"smartcore/internal/smartcore"
)

func main() {
	verbose := flag.Bool("verbose", false, "Enable verbose logging")
	flag.Parse()

	app := smartcore.NewApp(*verbose)
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
