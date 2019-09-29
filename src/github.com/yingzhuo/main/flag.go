package main

import (
	"flag"
	"github.com/bwmarrin/snowflake"
	"os"
)

var nodeInstance *snowflake.Node

var flags = &Flags{}

type Flags struct {
	host         string
	port         int
	nodeId       int64
	version      bool
	responseType string
}

func init() {
	cmd := flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	cmd.StringVar(&flags.host, "host", "0.0.0.0", "host of http service")
	cmd.IntVar(&flags.port, "port", 8080, "port of http service")
	cmd.Int64Var(&flags.nodeId, "node-id", 512, "node id of snowflake (0 ~ 1023)")
	cmd.BoolVar(&flags.version, "version", false, "print version")
	cmd.StringVar(&flags.responseType, "response-type", "protobuf", "content type of http response (json | protobuf)")
	_ = cmd.Parse(os.Args[1:])
}
