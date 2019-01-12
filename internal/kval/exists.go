package kval

import (
	"os"
	"path/filepath"
)

func (k Kval) Exists(dbname string, key string) bool {
	if !k.dbKeyCheck(key) {
		return false
	}
	/* Assume dbname is a valid database name */
	dbkey := filepath.Join(k.dir, dbname, key)

	/* Check whether dbkey exists */
	if _, err := os.Stat(dbkey); os.IsNotExist(err) {
		return false
	}
	return true
}
