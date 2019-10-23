package main

import (
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

type Flags struct {
	Host   string
	Port   int
	NodeId int64
	Type   string
	Indent bool
}

func main() {

	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetFormatter(&logrus.TextFormatter{
		DisableColors: false,
		FullTimestamp: true,
	})

	flags := &Flags{}

	app := cli.NewApp()
	app.Name = "snowflake"
	app.Usage = "a http server of id-generator"
	app.UsageText = "[options]"
	app.Authors = "应卓 <yingzhor@gmail.com>"
	app.Version = "1.0.1"
	app.BuildInfo = &cli.BuildInfo{
		Timestamp:   "",
		GitBranch:   "",
		GitCommit:   "",
		GitRevCount: "",
	}
	app.Flags = []*cli.Flag{
		{
			Name:          "h, host",
			Usage:         "host of http service",
			DefValue:      "0.0.0.0",
			NoOptDefValue: "0.0.0.0",
			Value:         &flags.Host,
		}, {
			Name:          "p, port",
			Usage:         "port of http service",
			DefValue:      "8080",
			NoOptDefValue: "8080",
			Value:         &flags.Port,
		}, {
			Name:          "n, node-id",
			Usage:         "id of snowflake node (0 ~ 1023)",
			DefValue:      "512",
			NoOptDefValue: "512",
			Value:         &flags.NodeId,
		}, {
			Name:          "t, type",
			Usage:         "type of http response (protobuf | json)",
			DefValue:      "protobuf",
			NoOptDefValue: "protobuf",
			Value:         &flags.Type,
		}, {
			Name:          "i, indent",
			Usage:         "output indented json",
			DefValue:      "false",
			NoOptDefValue: "false",
			Hidden:        true,
			IsBool:        true,
			Value:         &flags.Indent,
		},
	}

	app.Action = func(context *cli.Context) {
		doMain(flags)
	}

	app.Run(os.Args)
}

func doMain(flags *Flags) {
	logrus.Infof("pid            = %v", os.Getpid())
	logrus.Infof("host           = %v", flags.Host)
	logrus.Infof("port           = %v", flags.Port)
	logrus.Infof("node-id        = %v", flags.NodeId)
	logrus.Infof("type           = %v", flags.Type)

	startHttpServer(flags)
}

func startHttpServer(flags *Flags) {

	node, _ := snowflake.NewNode(flags.NodeId)

	// path: "/id"
	http.HandleFunc("/id", func(w http.ResponseWriter, r *http.Request) {
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

	if err := http.ListenAndServe(fmt.Sprintf("%s:%d", flags.Host, flags.Port), nil); err != nil {
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
