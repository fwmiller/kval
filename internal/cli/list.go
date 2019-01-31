package cli

import (
	"fmt"
)

func (c *Client) List() {
	dbs, err := c.db.List()
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, db := range dbs {
		fmt.Println(db)
	}
}
