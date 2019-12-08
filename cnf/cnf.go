package cnf

import (
	"errors"
	"strings"

	"github.com/bwmarrin/snowflake"
)

var (
	NodeId        int64
	ResponseType  Type
	Port          int
	Indent        bool
	QuietMode     bool
	SnowflakeNode *snowflake.Node
)

const (
	Json     Type = "json"
	Protobuf Type = "protobuf"
)

// --------------------------------------------------------------------------------------------------------------------

type Type string

func (t *Type) String() string {
	return string(*t)
}

func (t *Type) Set(value string) error {
	value = strings.TrimSpace(value)
	value = strings.ToLower(value)

	if string(Protobuf) == value {
		*t = Protobuf
		return nil
	}
	if string(Json) == value {
		*t = Json
		return nil
	}
	return errors.New(`fatal: flag type must be a option of "json" or "protobuf"`)
}
