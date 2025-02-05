package main

import (
	"log"
	"os"
	"web-scraper/internal/cmd"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("usage: go run . scrape [<args>]")
	}
	commandName := os.Args[1]
	commandArgs := os.Args[2:]

	commands := &cmd.Commands{}
	commands.Register("scrape", handlerScraper)
	cmd := cmd.Command{
		Name: commandName,
		Args: commandArgs,
	}
	err := commands.Run(cmd)
	if err != nil {
		log.Fatal(err)
	}
}
