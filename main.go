package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/subchen/go-cli"
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
			Value:    &cnf.Global.Port,
		}, {
			Name:     "n, node-id",
			Usage:    "id of snowflake node (0 ~ 1023)",
			DefValue: "512",
			Value:    &cnf.Global.NodeId,
		}, {
			Name:     "t, type",
			Usage:    "type of http response (protobuf | json)",
			DefValue: "protobuf",
			Value:    &cnf.Global.Type,
		}, {
			Name:     "i, indent",
			Usage:    "output indented json",
			DefValue: "false",
			Hidden:   true,
			IsBool:   true,
			Value:    &cnf.Global.Indent,
		}, {
			Name:     "q, quiet",
			Usage:    "quiet mode",
			DefValue: "false",
			IsBool:   true,
			Value:    &cnf.Global.QuietMode,
		},
	}

	app.Action = func(context *cli.Context) {
		doMain(context)
	}

	app.Run(os.Args)
}

func doMain(context *cli.Context) {

	cnf.Initialize() // 初始化所有全局变量

	if cnf.IsNotQuietMode() {
		logrus.Infof("pid            = %v", os.Getpid())
		logrus.Infof("port           = %v", cnf.GetHttpPort())
		logrus.Infof("node-id        = %v", cnf.GetNodeId())
		logrus.Infof("type           = %v", cnf.GetType())
	}

	engine := gin.Default()
	engine.GET("/healthz")
	engine.GET("/id", mappings.GenId)

	if err := engine.Run(cnf.GetHttpAddr()); err != nil {
		panic(err)
	}
}
