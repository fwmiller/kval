package kval

import (
	"fmt"
	"os"
	"path/filepath"
)

func (k Kval) Del(dbname string, key string) error {
	if !k.dbKeyCheck(key) {
		return fmt.Errorf("Illegal characters in key %s", key)
	}
	/* Assume dbname is a valid database name */
	dbkey := filepath.Join(k.dir, dbname, key)

	return os.Remove(dbkey)
}
