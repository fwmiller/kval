package kval

import (
	"regexp"
)

var keyNameCheck = regexp.MustCompile(`^[a-zA-Z0-9\-_]+$`).MatchString

// DB represents the things the databse can do
type DB interface {
	IsDB(dbName string) (string, error)
	Create(dbname string) error
	Remove(dbname string) error
	Keys(dbname string) ([]string, error)
	Set(dbname string, key string, value string) error
	Get(dbname string, key string) (string, error)
	Del(dbname string, key string) error
	List() ([]string, error)
	Time() (string)
	Exists(dbname string, key string) bool
}

type Kval struct {
	dir        string
	dbKeyCheck func(string) bool
}
