package kval

import (
	"fmt"
	"os"
)

func (k Kval) IsDB(dbname string) (string, error) {
	/* Assume dbname is a valid database name */
	dbpath := k.dir + "/" + dbname

	/* Check whether dbpath exists */
	if _, err := os.Stat(dbpath); os.IsNotExist(err) {
		return "", fmt.Errorf("Database %s does not exist", dbpath)
	}
	/* Check whether dbpath is a directory */
	fd, err := os.Open(dbpath)
	if err != nil {
		return "", fmt.Errorf("Open database %s failed, err: %s",
			dbpath, err)
	}
	stat, err := fd.Stat()
	if err != nil {
		return "", fmt.Errorf("Stat database %s failed, err: %s",
			dbpath, err)
	}
	if !stat.IsDir() {
		return "", fmt.Errorf("Database %s is not a directory",
			dbpath)
	}
	return dbname, nil
}
