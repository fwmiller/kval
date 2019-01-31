package cli

import (
	"fmt"
	"strings"
)

func (c *Client) Set(args string) {
	if c.Currdb == "" {
		fmt.Println("Current database not set")
		return
	}
	s := strings.SplitAfterN(args, " ", 2)
	if len(s) != 2 {
		fmt.Println("Missing value")
		return
	}
	key := strings.TrimSpace(s[0])
	value := strings.TrimSpace(s[1])

	/* Set key-value pair in current database */
	if err := c.db.Set(c.Currdb, key, value); err != nil {
		fmt.Println(err)
	}
}
