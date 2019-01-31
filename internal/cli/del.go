package cli

import (
	"fmt"
	"strings"
)

func (c *Client) Del(args string) {
	if c.Currdb == "" {
		fmt.Println("Current database not set")
		return
	}
	key := strings.TrimSpace(args)

	/* Delete key-value pair in current database */
	if err := c.db.Del(c.Currdb, key); err != nil {
		fmt.Println(err)
	}
}
