package kval

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func (k Kval) Set(dbname string, key string, value string) error {
	if !k.dbKeyCheck(key) {
		return fmt.Errorf("Illegal characters in key %s", key)
	}

	/* Assume dbname is a valid database name */
	dbkey := filepath.Join(k.dir, dbname, key)

	/* Check whether dbkey exists */
	if _, err := os.Stat(dbkey); !os.IsNotExist(err) {
		return fmt.Errorf("Value for key %s already set", dbkey)
	}

	/* Write value to new key file */
	return ioutil.WriteFile(dbkey, []byte(value), 0644)
}
