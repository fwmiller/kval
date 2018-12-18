package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/user"
	"regexp"
)

var kvaldir string
var dbKeyCheck = regexp.MustCompile(`^[a-zA-Z0-9\-_]+$`).MatchString

func KvalInit() {
	/* Get path for kval directory */
	usr, _ := user.Current()
	kvaldir = usr.HomeDir + "/.kval"
	fmt.Printf("kval directory = %v\n", kvaldir)

	/* Check if kval directory exists */
	_, err := os.Stat(kvaldir)
	if os.IsNotExist(err) {
		fmt.Printf("Create kval directory %v", kvaldir)
		err = os.Mkdir(kvaldir, 0777)
		if err != nil {
			fmt.Printf(" failed\n")
			os.Exit(0)
		}
		fmt.Printf("\n")
	}
}

func KvalIsDb(dbname string) string {
	/* Assume dbname is a valid database name */
	dbpath := kvaldir + "/" + dbname

	/* Check whether dbpath exists */
	if _, err := os.Stat(dbpath); os.IsNotExist(err) {
		fmt.Printf("Database %v does not exist\n", dbpath)
		return ""
	}
	/* Check whether dbpath is a directory */
	fd, err := os.Open(dbpath)
	if err != nil {
		fmt.Printf("Open database %v failed\n", dbpath)
		return ""
	}
	stat, err := fd.Stat()
	if err != nil {
		fmt.Printf("Stat database %v failed\n", dbpath)
		return ""
	}
	if !stat.IsDir() {
		fmt.Printf("Database %v is not a directory\n", dbpath)
		return ""
	}
	return dbname
}

func KvalCreateDb(dbname string) {
	/* Assume dbname is a valid database name */
	dbpath := kvaldir + "/" + dbname

	fmt.Printf("Create database %v", dbpath)
	err := os.Mkdir(dbpath, 0777)
	if err != nil {
		fmt.Printf(" failed")
	}
	fmt.Printf("\n")
}

func KvalSet(dbname string, key string, value string) {
	if !dbKeyCheck(key) {
		fmt.Printf("Illegal characters in %v\n", key)
		return
	}
	/* Assume dbname is a valid database name */
	dbkey := kvaldir + "/" + dbname + "/" + key

	/* Check whether dbkey exists */
	if _, err := os.Stat(dbkey); os.IsExist(err) {
		fmt.Printf("Key %v already exists\n", dbkey)
		return
	}
	/* Write value to new key file */
	err := ioutil.WriteFile(dbkey, []byte(value), 0644)
	if err != nil {
		fmt.Printf("Write to key file %v failed\n", dbkey)
	}
}
