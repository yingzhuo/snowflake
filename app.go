package main

import "github.com/bwmarrin/snowflake"

var nodeInstance *snowflake.Node

var flags = &Flags{}

type Flags struct {
	host         string
	port         int
	nodeId       int64
	version      bool
	responseType string
	indentJson   bool
	dryRun       bool
}

func main() {

	if flags.version {
		printVersion()
	} else {
		startHttpServer()
	}
}
