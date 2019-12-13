/*
*	 ____  _   _  _____        _______ _        _    _  _______
*	/ ___|| \ | |/ _ \ \      / /  ___| |      / \  | |/ / ____|
*	\___ \|  \| | | | \ \ /\ / /| |_  | |     / _ \ | ' /|  _|
*	 ___) | |\  | |_| |\ V  V / |  _| | |___ / ___ \| . \| |___
*	|____/|_| \_|\___/  \_/\_/  |_|   |_____/_/   \_\_|\_\_____|
*
*  https://github.com/yingzhuo/snowflake
 */
package main

import (
	"encoding/json"
	"fmt"
	"github.com/yingzhuo/snowflake/prome"
	"net/http"
	"os"
	"strconv"

	"github.com/bwmarrin/snowflake"
	"github.com/golang/protobuf/proto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/yingzhuo/go-cli/v2"
	"github.com/yingzhuo/snowflake/cnf"
	"github.com/yingzhuo/snowflake/protomsg"
)

var (
	BuildVersion   string = "latest"
	BuildGitBranch string
	BuildGitRev    string
	BuildGitCommit string
	BuildDate      string
)

func main() {

	app := cli.NewApp()
	app.Name = "snowflake"
	app.Usage = "http server of snowflake id-generator"
	app.UsageText = "[options]"
	app.Authors = "应卓 <yingzhor@gmail.com>"
	app.Version = BuildVersion
	app.BuildInfo = &cli.BuildInfo{
		GitBranch:   BuildGitBranch,
		GitCommit:   BuildGitCommit,
		GitRevCount: BuildGitRev,
		Timestamp:   BuildDate,
	}

	app.Examples = `snowflake --http-port=8080 --node-id=512 --type=protobuf 
snowflake --http-port=8080 --node-id=512 --type=json --indent`

	app.SeeAlso = `https://github.com/yingzhuo/snowflake
https://github.com/yingzhuo/snowflake-golang-client
https://github.com/yingzhuo/snowflake-java-client`

	app.Flags = []*cli.Flag{
		{
			Name:     "n, node-id",
			Usage:    "id of snowflake node (0 ~ 1023)",
			EnvVar:   "SNOWFLAKE_NODE_ID",
			DefValue: "0",
			Value:    &cnf.NodeId,
		}, {
			Name:     "p, http-port",
			Usage:    "port of http service",
			EnvVar:   "SNOWFLAKE_HTTP_PORT",
			DefValue: "8080",
			Value:    &cnf.Port,
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
			fmt.Printf("pid        = %v\n", os.Getpid())
			fmt.Printf("http-port  = %v\n", cnf.Port)
			fmt.Printf("node-id    = %v\n", cnf.NodeId)
			fmt.Printf("type       = %v\n", cnf.ResponseType)
			fmt.Printf("indent     = %v\n", cnf.Indent)
		}

		http.HandleFunc("/ping", func(w http.ResponseWriter, _ *http.Request) {
			w.Header().Set("Content-Type", "text/plain;charset=utf-8")
			w.WriteHeader(200)
			w.Write([]byte("pong"))
		})

		http.HandleFunc("/id", func(w http.ResponseWriter, request *http.Request) {
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

			prome.IdCreatedCounter.Add(float64(n)) // prometheus

			if cnf.ResponseType == cnf.Json {
				writeJson(result, w, cnf.Indent)
			} else {
				message := protomsg.IdList{
					Ids: []int64{},
				}
				message.Ids = append(message.Ids, result...)
				writeProtobuf(&message, w)
			}
		})

		http.Handle("/metrics", promhttp.Handler())

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
