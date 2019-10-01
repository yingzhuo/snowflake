package main

import (
	"flag"
	"os"
)

func init() {
	cmd := flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	cmd.StringVar(&flags.host, "host", "0.0.0.0", "host of http service")
	cmd.IntVar(&flags.port, "port", 8080, "port of http service")
	cmd.Int64Var(&flags.nodeId, "node-id", 512, "node id of snowflake (0 ~ 1023)")
	cmd.BoolVar(&flags.version, "version", false, "print version")
	cmd.StringVar(&flags.responseType, "response-type", "protobuf", "content type of http response (json | protobuf)")
	cmd.BoolVar(&flags.indentJson, "indent-json", false, "format json response body")
	cmd.BoolVar(&flags.dryRun, "dry-run", false, "dry run")
	_ = cmd.Parse(os.Args[1:])
}
