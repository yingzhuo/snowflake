package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/bwmarrin/snowflake"
	"github.com/golang/protobuf/proto"
	"github.com/sirupsen/logrus"
	"github.com/subchen/go-cli"
)

// build info
var (
	BuildVersion   string
	BuildGitBranch string
	BuildGitRev    string
	BuildGitCommit string
	BuildDate      string
)

var config = &Config{}

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
snowflake --port=8080 --node-id=512 --type=json --indent
snowflake --port=8080 --node-id=512 --type=json --indent --http-basic=root:root
`
	app.SeeAlso = `https://github.com/yingzhuo/snowflake
https://github.com/yingzhuo/snowflake-golang-client
https://github.com/yingzhuo/snowflake-java-client`

	app.Flags = []*cli.Flag{
		{
			Name:     "p, port",
			Usage:    "port of http service",
			DefValue: "8080",
			Value:    &config.Port,
		}, {
			Name:     "n, node-id",
			Usage:    "id of snowflake node (0 ~ 1023)",
			DefValue: "512",
			Value:    &config.NodeId,
		}, {
			Name:     "t, type",
			Usage:    "type of http response (protobuf | json)",
			DefValue: "protobuf",
			Value:    &config.Type,
		}, {
			Name:     "i, indent",
			Usage:    "output indented json",
			DefValue: "false",
			Hidden:   true,
			IsBool:   true,
			Value:    &config.Indent,
		}, {
			Name:     "q, quiet",
			Usage:    "quiet mode",
			DefValue: "false",
			IsBool:   true,
			Value:    &config.QuietMode,
		}, {
			Name:        "http-basic",
			Usage:       "enable http basic",
			Placeholder: "USERNAME:PASSWORD",
			DefValue:    "",
			Value:       &config.HttpBasic,
		},
	}

	app.Action = func(context *cli.Context) {
		doMain(config)
	}

	app.Run(os.Args)
}

func doMain(config *Config) {

	if !config.QuietMode {
		logrus.Infof("pid            = %v", os.Getpid())
		logrus.Infof("port           = %v", config.Port)
		logrus.Infof("node-id        = %v", config.NodeId)
		logrus.Infof("type           = %v", config.Type)
	}

	startHttpServer(config)
}

func startHttpServer(flags *Config) {

	node, _ := snowflake.NewNode(flags.NodeId)

	// path: "/id"
	http.HandleFunc("/id", func(w http.ResponseWriter, r *http.Request) {

		if !checkHttpBasic(config.HttpBasic, r.Header) {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		n, err := strconv.Atoi(r.FormValue("n"))
		if err != nil || n < 1 {
			n = 1
		}

		var result = make([]int64, 0)
		for i := 0; i < n; i++ {
			id := node.Generate()
			result = append(result, id.Int64())
		}

		switch {
		case strings.EqualFold(flags.Type, "json"):
			writeJson(w, result, flags.Indent)
		case strings.EqualFold(flags.Type, "protobuf"):
			writeProtobuf(w, result)
		default:
			writeProtobuf(w, result)
		}
	})

	// path: "/healthz"
	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {})

	if err := http.ListenAndServe(fmt.Sprintf("0.0.0.0:%d", flags.Port), nil); err != nil {
		log.Print(err)
		os.Exit(-1)
	}
}

func writeJson(w http.ResponseWriter, result []int64, indent bool) {

	w.Header().Set("Content-Type", "application/json;charset=utf-8")

	var data []byte
	var err error

	if indent {
		if data, err = json.MarshalIndent(result, "", "  "); err != nil {
			panic(err)
		}
	} else {
		if data, err = json.Marshal(result); err != nil {
			panic(err)
		}
	}

	if _, err := fmt.Fprint(w, string(data)); err != nil {
		_, _ = fmt.Fprint(w, "[]")
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func writeProtobuf(w http.ResponseWriter, result []int64) {
	w.Header().Set("Content-Type", "application/x-protobuf;charset=utf-8")

	message := IdList{
		Ids: []int64{},
	}

	message.Ids = append(message.Ids, result...)

	data, _ := proto.Marshal(&message)

	if _, err := w.Write(data); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func checkHttpBasic(basic HttpBasic, header http.Header) bool {

	if !basic.IsEnabled() {
		return true
	}

	headerValue := header.Get("Authorization")
	if headerValue == "" {
		return false
	}

	if !strings.HasPrefix(headerValue, "Basic ") {
		return false
	}

	headerValue = headerValue[6:]

	data, err := base64.URLEncoding.DecodeString(headerValue)
	if err != nil {
		return false
	}

	if usernameAndPassword := strings.Split(string(data), ":"); len(usernameAndPassword) != 2 {
		return false
	} else {
		return basic.Matches(usernameAndPassword[0], usernameAndPassword[1])
	}
}
