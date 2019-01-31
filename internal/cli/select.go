package cli

import (
	"fmt"
	"strings"
)

func (c *Client) Select(args string) {
	/* Check for valid dbname */
	n := strings.TrimSpace(args)
	if !dbNameCheck(n) {
		fmt.Printf("Illegal characters in %s\n", n)
		return
	}
	/* Check for valid database */
	name, err := c.db.IsDB(n)
	if err != nil {
		fmt.Println(err)
		return
	}
	if name != "" {
		c.Currdb = name
	}
}
