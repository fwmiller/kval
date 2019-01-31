package cli

import (
	"fmt"
)

func (c *Client) Keys() {
	if c.Currdb == "" {
		fmt.Println("Current database not set")
		return
	}
	keys, err := c.db.Keys(c.Currdb)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, k := range keys {
		fmt.Println(k)
	}
}
