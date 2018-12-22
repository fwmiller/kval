package client

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/fwmiller/kval/internal/kval"
)

var dbNameCheck = regexp.MustCompile(`^[a-zA-Z0-9\-_]+$`).MatchString

type Client struct {
	db     kval.DB
	Currdb string
}

func New(db kval.DB) *Client {
	return &Client{
		db: db,
	}
}

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

func (c *Client) Help() {
	fmt.Println("Help (add something useful here)")
}
