package main

import (
	"fmt"
	"os"

	"github.com/bwmarrin/snowflake"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/yingzhuo/go-cli/v2"
	"github.com/yingzhuo/snowflake/cnf"
	"github.com/yingzhuo/snowflake/mappings"
)

// build info
var (
	BuildVersion   string
	BuildGitBranch string
	BuildGitRev    string
	BuildGitCommit string
	BuildDate      string
)

func main() {

	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetFormatter(&logrus.TextFormatter{
		DisableColors: false,
		FullTimestamp: true,
	})

	app := cli.NewApp()
	app.Usage = "a http server of snowflake id-generator"
	app.UsageText = "[options]"
	app.Authors = "应卓 <yingzhor@gmail.com>"
	app.Version = BuildVersion
	app.BuildInfo = &cli.BuildInfo{
		GitBranch:   BuildGitBranch,
		GitCommit:   BuildGitCommit,
		GitRevCount: BuildGitRev,
		Timestamp:   BuildDate,
	}

	app.Examples = `snowflake --port=8080 --node-id=512 --type=protobuf 
snowflake --port=8080 --node-id=512 --type=json --indent`

	app.SeeAlso = `https://github.com/yingzhuo/snowflake
https://github.com/yingzhuo/snowflake-golang-client
https://github.com/yingzhuo/snowflake-java-client`

	app.Flags = []*cli.Flag{
		{
			Name:     "p, port",
			Usage:    "port of http service",
			DefValue: "8080",
			Value:    &cnf.Port,
		}, {
			Name:     "n, node-id",
			Usage:    "id of snowflake node (0 ~ 1023)",
			DefValue: "512",
			Value:    &cnf.NodeId,
		}, {
			Name:     "t, type",
			Usage:    "type of http response (protobuf | json)",
			DefValue: "protobuf",
			Value:    &cnf.ResponseType,
		}, {
			Name:     "i, indent",
			Usage:    "output indented json",
			DefValue: "false",
			IsBool:   true,
			Value:    &cnf.Indent,
		}, {
			Name:     "q, quiet",
			Usage:    "quiet mode",
			DefValue: "false",
			IsBool:   true,
			Value:    &cnf.QuietMode,
		},
	}

	app.OnAppInitialized = func(_ *cli.Context) {
		cnf.SnowflakeNode, _ = snowflake.NewNode(cnf.NodeId)
	}

	app.Action = func(c *cli.Context) {
		doMain(c)
	}

	app.Run(os.Args)
}

func doMain(_ *cli.Context) {

	if !cnf.QuietMode {
		logrus.Infof("pid            = %v", os.Getpid())
		logrus.Infof("port           = %v", cnf.Port)
		logrus.Infof("node-id        = %v", cnf.NodeId)
		logrus.Infof("type           = %v", cnf.ResponseType)
	}

	engine := gin.New()
	engine.Use(gin.Recovery())
	if !cnf.QuietMode {
		engine.Use(gin.Logger())
	}

	engine.GET("/id", mappings.GenId)
	engine.GET("/healthz")

	if err := engine.Run(fmt.Sprintf("0.0.0.0:%d", cnf.Port)); err != nil {
		panic(err)
	}
}
