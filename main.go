package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/bwmarrin/snowflake"
	"github.com/golang/protobuf/proto"
	"github.com/sirupsen/logrus"
	"github.com/yingzhuo/go-cli/v2"
	"github.com/yingzhuo/snowflake/cnf"
	"github.com/yingzhuo/snowflake/protomsg"
)

// build info
var (
	BuildVersion   string = "latest"
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
	app.Name = "snowflake"
	app.Usage = "a http server of snowflake id-generator"
	app.UsageText = "[options]"
	app.Authors = "应卓 <yingzhor@gmail.com>"
	app.Version = BuildVersion
	app.BuildInfo = &cli.BuildInfo{
		GitBranch:   BuildGitBranch,
		GitCommit:   BuildGitCommit,
		GitRevCount: BuildGitRev,
		Timestamp:   BuildDate,
		BuiltBy:     "make",
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
			EnvVar:   "SNOWFLAKE_PORT",
			DefValue: "8080",
			Value:    &cnf.Port,
		}, {
			Name:     "n, node-id",
			Usage:    "id of snowflake node (0 ~ 1023)",
			EnvVar:   "SNOWFLAKE_NODE_ID",
			DefValue: "0",
			Value:    &cnf.NodeId,
		}, {
			Name:     "t, type",
			Usage:    "type of http response (protobuf | json)",
			EnvVar:   "SNOWFLAKE_TYPE",
			DefValue: "protobuf",
			Value:    &cnf.ResponseType,
		}, {
			Name:          "i, indent",
			Usage:         "output indented json",
			EnvVar:        "SNOWFLAKE_INDENT",
			DefValue:      "true",
			NoOptDefValue: "true",
			IsBool:        true,
			Value:         &cnf.Indent,
		}, {
			Name:     "q, quiet",
			Usage:    "quiet mode",
			EnvVar:   "SNOWFLAKE_QUIET",
			DefValue: "false",
			IsBool:   true,
			Value:    &cnf.QuietMode,
		},
	}

	app.OnAppInitialized = func(_ *cli.Context) {
		cnf.SnowflakeNode, _ = snowflake.NewNode(cnf.NodeId)
	}

	app.Action = func(c *cli.Context) {
		if !cnf.QuietMode {
			logrus.Infof("pid            = %v", os.Getpid())
			logrus.Infof("port           = %v", cnf.Port)
			logrus.Infof("node-id        = %v", cnf.NodeId)
			logrus.Infof("type           = %v", cnf.ResponseType)
			if strings.EqualFold("json", cnf.ResponseType.String()) {
				logrus.Infof("indent         = %v", cnf.Indent)
			}
		}

		http.HandleFunc("/ping", func(writer http.ResponseWriter, request *http.Request) {
			writer.WriteHeader(200)
			writer.Write([]byte("pong"))
		})

		http.HandleFunc("/id", func(writer http.ResponseWriter, request *http.Request) {
			n := 1
			vs := request.FormValue("n")
			n, _ = strconv.Atoi(vs)

			if n <= 0 {
				n = 1
			}

			var result = make([]int64, 0)
			for i := 0; i < n; i++ {
				id := cnf.SnowflakeNode.Generate()
				result = append(result, id.Int64())
			}

			if cnf.ResponseType == cnf.Json {
				writeJson(result, writer, cnf.Indent)
			} else {
				message := protomsg.IdList{
					Ids: []int64{},
				}
				message.Ids = append(message.Ids, result...)
				writeProtobuf(&message, writer)
			}
		})

		http.ListenAndServe(fmt.Sprintf("0.0.0.0:%d", cnf.Port), nil)
	}

	app.Run(os.Args)
}

func writeJson(model interface{}, w http.ResponseWriter, indent bool) {
	w.Header().Set("Content-Type", "application/json;charset=utf-8")

	var bytes []byte
	var err error

	if indent {
		bytes, err = json.MarshalIndent(model, "", "  ")
	} else {
		bytes, err = json.Marshal(model)
	}

	if err != nil {
		panic(err)
	}

	if _, err = fmt.Fprint(w, string(bytes)); err != nil {
		panic(err)
	}
}

func writeProtobuf(model proto.Message, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/x-protobuf;charset=utf-8")

	data, _ := proto.Marshal(model)
	if _, err := w.Write(data); err != nil {
		panic(err)
	}
}
