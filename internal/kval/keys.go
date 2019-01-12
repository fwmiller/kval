package kval

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
)

func (k Kval) Keys(dbname string) ([]string, error) {
	/* Assume dbname is a valid database name */
	dbpath := filepath.Join(k.dir, dbname)
	files, err := ioutil.ReadDir(dbpath)
	if err != nil {
		return nil, fmt.Errorf("Read database dir %s failed, err: %s",
					dbpath, err)
	}
	out := make([]string, len(files))
	for i, file := range files {
		out[i] = file.Name()
	}
	return out, nil
}
