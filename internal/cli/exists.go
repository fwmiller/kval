package cli

import (
	"fmt"
	"strings"
)

func (c *Client) Exists(args string) {
	if c.Currdb == "" {
		fmt.Println("Current database not set")
		return
	}
	key := strings.TrimSpace(args)
	fmt.Println(c.db.Exists(c.Currdb, key))
}
