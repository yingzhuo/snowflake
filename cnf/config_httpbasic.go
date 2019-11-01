package cnf

import (
	"errors"
	"strings"
)

type UsernamePassword struct {
	Username string
	Password string
}

func (w *UsernamePassword) String() string {
	return w.Username + ":[****]"
}

func (w *UsernamePassword) Set(value string) error {

	if up := strings.Split(value, ":"); len(up) == 2 {
		w.Username = strings.TrimSpace(up[0])
		w.Password = strings.TrimSpace(up[1])
	}

	if w.Username == "" || w.Password == "" {
		return errors.New("invalid username or password")
	} else {
		return nil
	}
}
