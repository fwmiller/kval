package kval

import (
	"fmt"
	"io/ioutil"
)

func (k Kval) List() ([]string, error) {
	files, err := ioutil.ReadDir(k.dir)
	if err != nil {
		return nil, fmt.Errorf("Read database dir %s failed, err: %s",
					k.dir, err)
	}
	out := make([]string, len(files))
	for i, file := range files {
		out[i] = file.Name()
	}
	return out, nil
}
