package kval

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func (k Kval) Get(dbname string, key string) (string, error) {
	if !k.dbKeyCheck(key) {
		return "", fmt.Errorf("Illegal characters in key %s", key)
	}
	/* Assume dbname is a valid database name */
	dbkey := filepath.Join(k.dir, dbname, key)

	/* Check whether dbkey exists */
	if _, err := os.Stat(dbkey); os.IsNotExist(err) {
		return "", err
	}

	/* Read value from key file */
	value, err := ioutil.ReadFile(dbkey)
	if err != nil {
		return "", fmt.Errorf("Read key file %s failed", dbkey)
	}
	return string(value), nil
}
