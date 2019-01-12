package kval

import (
	"os"
	"os/user"
)

func New() (*Kval, error) {
	/* Get path for kval directory */
	usr, _ := user.Current()
	dir := usr.HomeDir + "/.kval"

	/* Check if kval directory exists */
	_, err := os.Stat(dir)
	if os.IsNotExist(err) {
		err = os.Mkdir(dir, 0777)
		if err != nil {
			return nil, err
		}
	}
	return &Kval {
		dir:        dir,
		dbKeyCheck: keyNameCheck,
	}, nil
}
