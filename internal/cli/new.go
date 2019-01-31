package cli

import (
	"github.com/fwmiller/kval/internal/kval"
)

func New(db kval.DB) *Client {
	return &Client{
		db: db,
	}
}
