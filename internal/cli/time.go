package cli

import (
	"fmt"
)

func (c *Client) Time() {
	time := c.db.Time()
	fmt.Println(time)
}
