package kval

import (
	"os"
	"path/filepath"
)

func (k Kval) Create(dbname string) error {
	/* Assume dbname is a valid database name */
	dbpath := filepath.Join(k.dir, dbname)
	return os.Mkdir(dbpath, 0777)
}
