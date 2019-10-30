package main

type Config struct {
	Port      int
	NodeId    int64
	Type      string
	Indent    bool
	QuietMode bool
	HttpBasic
}
