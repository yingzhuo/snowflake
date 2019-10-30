package main

import (
	"errors"
	"fmt"
	"strings"
)

type HttpBasic struct {
	username string
	password string
}

func (h *HttpBasic) String() string {
	return fmt.Sprintf("%s:%s", h.username, h.password)
}

func (h *HttpBasic) Set(value string) error {
	if value == "" {
		return nil
	}

	if up := strings.Split(value, ":"); len(up) == 2 {
		username := strings.TrimSpace(up[0])
		password := strings.TrimSpace(up[1])

		if username == "" || password == "" {
			return errors.New("fatal: bad http-basic username or password")
		} else {
			h.username = username
			h.password = password
			return nil
		}
	}

	return errors.New("fatal: bad http-basic username or password")
}

func (h *HttpBasic) IsEnabled() bool {
	return h.username != "" && h.password != ""
}

func (h *HttpBasic) Matches(username, password string) bool {
	return h.username == username && h.password == password
}

func NewHttpBasic() *HttpBasic {
	return &HttpBasic{
		username: "",
		password: "",
	}
}
