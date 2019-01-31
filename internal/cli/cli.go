package cli

import (
	"regexp"

	"github.com/fwmiller/kval/internal/kval"
)

var dbNameCheck = regexp.MustCompile(`^[a-zA-Z0-9\-_]+$`).MatchString

type Client struct {
	db     kval.DB
	Currdb string
}
