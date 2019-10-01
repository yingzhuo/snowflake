package main

import (
	"encoding/json"
	"fmt"
	"github.com/bwmarrin/snowflake"
	"github.com/golang/protobuf/proto"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func startHttpServer() {

	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
	log.SetPrefix("[SNOWFLAKE] ")
	log.SetOutput(os.Stdout)
	log.Printf("host           = %v", flags.host)
	log.Printf("port           = %v", flags.port)
	log.Printf("node-id        = %v", flags.nodeId)
	log.Printf("reponse-type   = %v", flags.responseType)
	if strings.EqualFold("json", flags.responseType) {
		log.Printf("indent-json    = %v", flags.indentJson)
	}
	log.Printf("status         = Running")

	if flags.dryRun {
		return
	}

	nodeInstance, _ = snowflake.NewNode(flags.nodeId)

	// path: "/id"
	http.HandleFunc("/id", func(w http.ResponseWriter, r *http.Request) {
		n, err := strconv.Atoi(r.FormValue("n"))
		if err != nil || n < 1 {
			n = 1
		}

		var result = make([]int64, 0)
		for i := 0; i < n; i++ {
			id := nodeInstance.Generate()
			result = append(result, id.Int64())
		}

		switch {
		case strings.EqualFold(flags.responseType, "json"):
			writeJson(w, result, flags.indentJson)
		case strings.EqualFold(flags.responseType, "protobuf"):
			writeProtobuf(w, result)
		default:
			writeJson(w, result, flags.indentJson) // default as json
		}
	})

	// path: "/healthz"
	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {})

	if err := http.ListenAndServe(fmt.Sprintf("%s:%d", flags.host, flags.port), nil); err != nil {
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
