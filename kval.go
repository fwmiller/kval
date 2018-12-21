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
	fmt.Printf("kval directory = %s\n", kvaldir)

	/* Check if kval directory exists */
	_, err := os.Stat(kvaldir)
	if os.IsNotExist(err) {
		fmt.Printf("Create kval directory %s", kvaldir)
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
		fmt.Printf("Database %s does not exist\n", dbpath)
		return ""
	}
	/* Check whether dbpath is a directory */
	fd, err := os.Open(dbpath)
	if err != nil {
		fmt.Printf("Open database %s failed\n", dbpath)
		return ""
	}
	stat, err := fd.Stat()
	if err != nil {
		fmt.Printf("Stat database %s failed\n", dbpath)
		return ""
	}
	if !stat.IsDir() {
		fmt.Printf("Database %s is not a directory\n", dbpath)
		return ""
	}
	return dbname
}

func KvalCreateDb(dbname string) bool {
	/* Assume dbname is a valid database name */
	dbpath := kvaldir + "/" + dbname

	fmt.Printf("Create database %s", dbpath)
	err := os.Mkdir(dbpath, 0777)
	if err != nil {
		fmt.Printf(" failed\n")
		return false
	}
	fmt.Printf("\n")
	return true
}

func KvalRemoveDb(dbname string) bool {
	/* Assume dbname is a valid database name */
	dbpath := kvaldir + "/" + dbname

	fmt.Printf("Remove database %s", dbpath)
	err := os.RemoveAll(dbpath)
	if err != nil {
		fmt.Printf(" failed\n")
		return false
	}
	fmt.Printf("\n")
	return true
}

func KvalSet(dbname string, key string, value string) {
	if !dbKeyCheck(key) {
		fmt.Printf("Illegal characters in key %s\n", key)
		return
	}
	/* Assume dbname is a valid database name */
	dbkey := kvaldir + "/" + dbname + "/" + key

	/* Check whether dbkey exists */
	if _, err := os.Stat(dbkey); !os.IsNotExist(err) {
		fmt.Printf("Value for key %s already set\n", dbkey)
		return
	}
	/* Write value to new key file */
	err := ioutil.WriteFile(dbkey, []byte(value), 0644)
	if err != nil {
		fmt.Printf("Write to key file %s failed\n", dbkey)
	}
}

func KvalGet(dbname string, key string) string {
	if !dbKeyCheck(key) {
		fmt.Printf("Illegal characters in key %s\n", key)
		return ""
	}
	/* Assume dbname is a valid database name */
	dbkey := kvaldir + "/" + dbname + "/" + key

	/* Check whether dbkey exists */
	if _, err := os.Stat(dbkey); os.IsNotExist(err) {
		return ""
	}
	/* Read value from key file */
	value, err := ioutil.ReadFile(dbkey)
	if err != nil {
		fmt.Printf("Read key file %s failed\n", dbkey)
		return ""
	}
	return string(value)
}

func KvalDel(dbname string, key string) {
	if !dbKeyCheck(key) {
		fmt.Printf("Illegal characters in key %s\n", key)
		return
	}
	/* Assume dbname is a valid database name */
	dbkey := kvaldir + "/" + dbname + "/" + key

	fmt.Printf("Delete %s", dbkey)
	err := os.Remove(dbkey)
	if err != nil {
		fmt.Printf(" failed\n")
		return
	}
	fmt.Printf("\n")
}
