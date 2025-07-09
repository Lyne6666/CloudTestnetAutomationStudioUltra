// cmd/cloudtestnetautomationstudioultra/main.go
package main

import (
	"flag"
	"log"
	"os"
	
	"cloudtestnetautomationstudioultra/internal/cloudtestnetautomationstudioultra"
)

func main() {
	verbose := flag.Bool("verbose", false, "Enable verbose logging")
	flag.Parse()
	
	app := cloudtestnetautomationstudioultra.NewApp(*verbose)
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
