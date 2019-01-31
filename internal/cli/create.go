package cli

import (
	"fmt"
	"strings"
)

func (c *Client) Create(args string) {
	/* Check for valid dbname */
	dbname := strings.TrimSpace(args)
	if !dbNameCheck(dbname) {
		fmt.Printf("Illegal characters in %s\n", dbname)
		return
	}
	/* Create new database */
	err := c.db.Create(dbname)
	if err != nil {
		fmt.Println(err)
		return
	}

	/* Set current to new database if it was clear */
	if c.Currdb == "" {
		c.Currdb = dbname
	}
}
