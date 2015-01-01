package main

import (
	"fmt"
	"log"
	flag "mflag"
	"os"
)

func main() {
	flag.Parse()

	var (
		cli *PlCli
	)

	if *flVersion {
		showVersion()
		return
	}

	cli = CreatePlCli(os.Stdin, os.Stdout, os.Stderr, *flVerbose)

	if err := cli.Cmd(flag.Args()...); err != nil {
		if statusError, ok := err.(*utils.StatusError); ok {
			if statusError.Status != "" {
				log.Println(statusError.Status)
			}
			os.Exit(statusError.StatusCode)
		}
		log.Fatal(err)
	}
}

func pingServer() {
	pong, err := ping.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v\n", pong)
}

func showVersion() {
	fmt.Printf("pl version 0.0.1\n")
}
