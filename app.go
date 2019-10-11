package main

import (
	"flag"
	"github.com/bwmarrin/snowflake"
	"github.com/sirupsen/logrus"
	"os"
	"strings"
)

var nodeInstance *snowflake.Node

var flags = &Flags{}

type Flags struct {
	host         string
	port         int
	nodeId       int64
	version      bool
	responseType string
	indentJson   bool
	timezone     string
}

func main() {

	if flags.version {
		printVersion()
	} else {
		logrus.Infof("pid            = %v", os.Getpid())
		logrus.Infof("host           = %v", flags.host)
		logrus.Infof("port           = %v", flags.port)
		logrus.Infof("node-id        = %v", flags.nodeId)
		logrus.Infof("response-type  = %v", flags.responseType)
		logrus.Infof("timezone       = %v", flags.timezone)
		if strings.EqualFold("json", flags.responseType) {
			logrus.Infof("indent-json    = %v", flags.indentJson)
		}
		logrus.Infof("status         = Running")

		startHttpServer()
	}
}

func init() {
	cmd := flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	cmd.StringVar(&flags.host, "host", "0.0.0.0", "host of http service")
	cmd.IntVar(&flags.port, "port", 8080, "port of http service")
	cmd.Int64Var(&flags.nodeId, "node-id", 512, "node id of snowflake (0 ~ 1023)")
	cmd.BoolVar(&flags.version, "version", false, "print version")
	cmd.StringVar(&flags.responseType, "response-type", "protobuf", "content type of http response (json | protobuf)")
	cmd.BoolVar(&flags.indentJson, "indent-json", false, "format json response body")
	cmd.StringVar(&flags.timezone, "timezone", "Asia/Shanghai", "timezone")
	_ = cmd.Parse(os.Args[1:])

	_ = os.Setenv("TZ", flags.timezone)

	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetFormatter(&logrus.TextFormatter{
		DisableColors: false,
		FullTimestamp: true,
	})
}
