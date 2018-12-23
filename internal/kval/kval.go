package kval

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/user"
	"path/filepath"
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
}

type Kval struct {
	dir        string
	dbKeyCheck func(string) bool
}

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

func (k Kval) Create(dbname string) error {
	/* Assume dbname is a valid database name */
	dbpath := filepath.Join(k.dir, dbname)
	return os.Mkdir(dbpath, 0777)
}

func (k Kval) Remove(dbname string) error {
	/* Assume dbname is a valid database name */
	dbpath := filepath.Join(k.dir, dbname)
	return os.RemoveAll(dbpath)
}

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

func (k Kval) Del(dbname string, key string) error {
	if !k.dbKeyCheck(key) {
		return fmt.Errorf("Illegal characters in key %s", key)
	}
	/* Assume dbname is a valid database name */
	dbkey := filepath.Join(k.dir, dbname, key)

	return os.Remove(dbkey)
}
