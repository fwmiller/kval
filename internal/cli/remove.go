package cli

import (
	"fmt"
	"strings"
)

func (c *Client) Remove(args string) {
	/* Check for valid dbname */
	dbname := strings.TrimSpace(args)
	if !dbNameCheck(dbname) {
		fmt.Printf("Illegal characters in %s\n", dbname)
		return
	}
	/* Remove existing database */
	err := c.db.Remove(dbname)
	if err != nil {
		fmt.Println(err)
		return
	}

	/* Clear current database if it was just removed */
	if dbname == c.Currdb {
		c.Currdb = ""
	}
}
