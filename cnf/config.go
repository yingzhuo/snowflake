package cnf

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/snowflake"
)

type Config struct {
	Port      int
	NodeId    int64
	Type      string
	Indent    bool
	QuietMode bool
}

var Global = &Config{}
var SnowflakeNode *snowflake.Node = nil

func Initialize() {
	SnowflakeNode, _ = snowflake.NewNode(Global.NodeId)
}

func GetHttpAddr() string {
	return fmt.Sprintf("0.0.0.0:%d", Global.Port)
}

func IsQuietMode() bool {
	return Global.QuietMode
}

func IsNotQuietMode() bool {
	return !IsQuietMode()
}

func GetHttpPort() int {
	return Global.Port
}

func GetNodeId() int64 {
	return Global.NodeId
}

func GetType() string {
	return Global.Type
}

func IsIndentMode() bool {
	return Global.Indent
}

func IsJsonType() bool {
	return strings.EqualFold("json", Global.Type)
}

func IsProtobufType() bool {
	return strings.EqualFold("protobuf", Global.Type)
}
