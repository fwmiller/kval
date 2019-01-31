package cli

import (
	"fmt"
	"strings"
)

func (c *Client) Get(args string) {
	if c.Currdb == "" {
		fmt.Println("Current database not set")
		return
	}
	key := strings.TrimSpace(args)

	/* Get value associated with key in current database */
	value, err := c.db.Get(c.Currdb, key)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(value)
	}
}
